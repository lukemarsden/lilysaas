package scheduler

import (
	"context"
	"fmt"
	"slices"
	"time"

	"github.com/google/uuid"
	"github.com/helixml/helix/api/pkg/config"
	"github.com/helixml/helix/api/pkg/data"
	"github.com/helixml/helix/api/pkg/types"
	"github.com/puzpuzpuz/xsync/v3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/exp/rand"
)

const (
	pendingSlotsBufferSize    = 1 // The number of slot creation requests to buffer
	runnerReconcileInterval   = 5 * time.Second
	activityReconcileInterval = 100 * time.Millisecond
	queueReconcileInterval    = 100 * time.Millisecond
)

// TimeoutFunc defines a function type that determines if a runner has timed out based on the last activity.
type TimeoutFunc func(runnerID string, lastActivityTime time.Time) bool

func NewTimeoutFunc(ttl time.Duration) TimeoutFunc {
	return func(_ string, lastActivityTime time.Time) bool {
		return lastActivityTime.Add(ttl).Before(time.Now())
	}
}

// Add these new types to track pending slot creation
type PendingSlot struct {
	Work     *Workload
	RunnerID string
	Created  chan struct{} // Channel to signal when slot is created
}

type Scheduler struct {
	ctx             context.Context
	controller      *RunnerController
	queue           *WorkQueue
	onSchedulingErr func(work *Workload, err error)
	slots           *xsync.MapOf[uuid.UUID, *Slot]
	modelStaleFunc  TimeoutFunc       // Function to check if models are stale
	slotTimeoutFunc TimeoutFunc       // Function to check if slots have timed out due to error
	pendingSlots    chan *PendingSlot // Channel for pending slot creations
}

type Params struct {
	RunnerController  *RunnerController
	QueueSize         int
	OnSchedulingErr   func(work *Workload, err error)
	OnResponseHandler func(ctx context.Context, resp *types.RunnerLLMInferenceResponse) error
}

func NewScheduler(ctx context.Context, serverConfig *config.ServerConfig, params *Params) (*Scheduler, error) {
	if params == nil {
		params = &Params{}
	}
	if params.RunnerController == nil {
		return nil, fmt.Errorf("runner controller is required")
	}
	modelTTL := serverConfig.Providers.Helix.ModelTTL
	if modelTTL == 0 {
		modelTTL = 10 * time.Second
	}
	slotTTL := serverConfig.Providers.Helix.SlotTTL
	if slotTTL == 0 {
		slotTTL = 300 * time.Second
	}
	queueSize := 50
	if params.QueueSize > 0 {
		queueSize = params.QueueSize
	}

	log.Info().Dur("model_stale_time", modelTTL).Dur("slot_timeout", slotTTL).Msg("slot timeouts")

	s := &Scheduler{
		ctx:             ctx,
		controller:      params.RunnerController,
		queue:           NewWorkQueue(queueSize),
		onSchedulingErr: params.OnSchedulingErr,
		slots:           xsync.NewMapOf[uuid.UUID, *Slot](),
		modelStaleFunc:  NewTimeoutFunc(modelTTL),
		slotTimeoutFunc: NewTimeoutFunc(slotTTL),
		pendingSlots:    make(chan *PendingSlot, pendingSlotsBufferSize),
	}

	// Start the queue processor
	go s.processQueue(ctx)

	// Start the slot reconciler
	go s.reconcileSlots(ctx)

	// Start the activity reconciler
	go s.reconcileActivity(ctx)

	// Start the runner reconciler
	go s.reconcileRunners(ctx)

	// Start the slot creator
	go s.runSlotCreator(ctx)

	return s, nil
}

func (s *Scheduler) Enqueue(work *Workload) error {
	return s.queue.Add(work)
}

func (s *Scheduler) Queue() ([]*types.WorkloadSummary, error) {
	currentQueue := s.queue.Queue()
	queue := make([]*types.WorkloadSummary, 0, len(currentQueue))
	for _, w := range currentQueue {
		summary := ""
		switch w.WorkloadType {
		case WorkloadTypeLLMInferenceRequest:
			req := w.LLMInferenceRequest()
			summary = req.Request.Messages[len(req.Request.Messages)-1].Content
		case WorkloadTypeSession:
			s := w.Session()
			interaction, _ := data.GetLastUserInteraction(s.Interactions)
			if interaction != nil {
				summary = interaction.Message
			}
		}
		queue = append(queue, &types.WorkloadSummary{
			ID:        w.ID(),
			CreatedAt: w.Created(),
			UpdatedAt: w.Updated(),
			ModelName: string(w.ModelName()),
			Mode:      string(w.Mode()),
			Runtime:   string(w.Runtime()),
			LoraDir:   w.LoraDir(),
			Summary:   summary,
		})
	}
	return queue, nil
}

func (s *Scheduler) RunnerStatus() ([]*types.RunnerStatus, error) {
	var err error
	// Get a current list of runners
	runners := s.controller.RunnerIDs()

	// Get the current state of each runner
	runnerStates := make([]*types.RunnerStatus, 0, len(runners))
	for _, runnerID := range runners {
		var runnerStatus *types.RunnerStatus
		runnerStatus, err = s.controller.GetStatus(runnerID)
		if err != nil {
			log.Warn().Err(err).Str("runner_id", runnerID).Msg("error getting runner status, this shouldn't happen, please investigate this runner")
			runnerStatus = &types.RunnerStatus{
				ID: runnerID,
			}
		}
		runnerStates = append(runnerStates, runnerStatus)
	}

	return runnerStates, nil
}

func (s *Scheduler) RunnerSlots(runnerID string) ([]*types.RunnerSlot, error) {
	return s.controller.GetSlots(runnerID)
}

// processQueue runs in a goroutine to processes the queue of requests.
func (s *Scheduler) processQueue(ctx context.Context) {
	log.Debug().Msg("starting queue processor")
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(queueReconcileInterval):
			s.processQueueOnce()
		}
	}
}

// reconcileSlots runs in a goroutine to reconcile slots.
// The reason why we do this async is because we don't want to have to check the runner on the hot
// path. When a user makes a request we want to forward it to a warm runner as quickly as possible.
func (s *Scheduler) reconcileSlots(ctx context.Context) {
	log.Debug().Msg("starting slot reconciler")
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(runnerReconcileInterval):
			s.reconcileSlotsOnce()
		}
	}
}

// reconcileActivity runs in a goroutine to reconcile activity.
func (s *Scheduler) reconcileActivity(ctx context.Context) {
	log.Debug().Msg("starting activity reconciler")
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(activityReconcileInterval):
			s.reconcileActivityOnce()
		}
	}
}

func (s *Scheduler) reconcileActivityOnce() {
	if s.slots == nil {
		return
	}
	// Check the status of all remaining slots to see if they have finished their work
	s.slots.Range(func(slotID uuid.UUID, slot *Slot) bool {
		if slot == nil {
			withSlotContext(&log.Logger, slot).Debug().Msg("slot is nil, releasing")
			slot.Release()
			return true
		}
		if slot.IsActive() {
			// Get the live slot from the runner, don't use the cached copy
			remoteSlot, err := s.controller.fetchSlot(slot.RunnerID, slotID)
			if err != nil {
				withSlotContext(&log.Logger, slot).Error().Err(err).Msg("failed to get slot, assuming it's finished")
				slot.Release()
			} else if !remoteSlot.Active {
				withSlotContext(&log.Logger, slot).Debug().Msg("slot is not active, releasing")
				slot.Release()
			}
		}
		return true
	})
}

// reconcileRunners runs in a goroutine to reconcile runners.
func (s *Scheduler) reconcileRunners(ctx context.Context) {
	log.Debug().Msg("starting runner reconciler")
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(runnerReconcileInterval):
			s.reconcileRunnersOnce()
		}
	}
}

func (s *Scheduler) reconcileRunnersOnce() {
	// Get all runners
	runnerIDs := s.controller.RunnerIDs()

	// Get the health of each runner
	for _, runnerID := range runnerIDs {
		err := s.controller.GetHealthz(runnerID)
		if err != nil {
			log.Error().Err(err).Str("runner_id", runnerID).Msg("runner is not healthy, deleting...")
			s.controller.deleteRunner(runnerID)
		}
	}
}

// reconcileSlotsOnce reconciles slots once.
func (s *Scheduler) reconcileSlotsOnce() {
	// Get all runners
	runnerIDs := s.controller.RunnerIDs()

	// Build a complete map of all actual slots across all runners
	allActualSlots := make(map[uuid.UUID]string) // maps slot ID to runner ID
	for _, runnerID := range runnerIDs {
		actualSlots, err := s.controller.GetSlots(runnerID)
		if err != nil {
			log.Error().Err(err).Str("runner_id", runnerID).Msg("failed to get slots from runner")
			continue
		}

		for _, slot := range actualSlots {
			// If we find the same slot ID on multiple runners, delete from the duplicate runner
			if existingRunnerID, exists := allActualSlots[slot.ID]; exists {
				log.Warn().
					Str("slot_id", slot.ID.String()).
					Str("existing_runner", existingRunnerID).
					Str("duplicate_runner", runnerID).
					Msg("found duplicate slot ID on multiple runners, cleaning up duplicate")
				err = s.controller.DeleteSlot(runnerID, slot.ID)
				if err != nil {
					log.Error().Err(err).Str("runner_id", runnerID).Str("slot_id", slot.ID.String()).Msg("failed to delete duplicate slot")
				}
				continue
			}
			allActualSlots[slot.ID] = runnerID
		}
	}

	// Clean up scheduler slots that don't exist on any runner
	s.slots.Range(func(slotID uuid.UUID, slot *Slot) bool {
		if runnerID, exists := allActualSlots[slotID]; !exists {
			log.Warn().
				Str("runner_id", slot.RunnerID).
				Str("slot_id", slotID.String()).
				Msg("found slot on the scheduler that doesn't exist on any runner, deleting...")
			s.slots.Delete(slotID)
		} else if runnerID != slot.RunnerID {
			// The slot exists but on a different runner than we thought
			log.Warn().
				Str("scheduler_runner", slot.RunnerID).
				Str("actual_runner", runnerID).
				Str("slot_id", slotID.String()).
				Msg("slot exists on different runner than expected, updating runner ID")
			slot.RunnerID = runnerID
		}
		return true
	})

	// Clean up runner slots that don't exist in scheduler
	for slotID, runnerID := range allActualSlots {
		if _, exists := s.slots.Load(slotID); !exists {
			log.Warn().
				Str("runner_id", runnerID).
				Str("slot_id", slotID.String()).
				Msg("found slot on runner that doesn't exist in scheduler, deleting...")
			err := s.controller.DeleteSlot(runnerID, slotID)
			if err != nil {
				log.Error().Err(err).Str("runner_id", runnerID).Str("slot_id", slotID.String()).Msg("failed to delete slot")
			}
		}
	}

	// Ensure new slots are created and ready to take work
	requiredSlots := s.queue.GetRequiredSlots()
	log.Debug().Interface("required_slots", requiredSlots).Msg("required slots")
	// For each requirement, ensure we have enough slots for that work right now
	for _, req := range requiredSlots {
		// Check if we have enough slots for this work right now
		existingCount := 0
		// TODO(Phil): be careful about this, it's detached from the concept of a warm slot. Ideally
		// refactor so that warm and this are using the same logic.
		s.slots.Range(func(_ uuid.UUID, slot *Slot) bool {
			if slot.ModelName() == req.Model &&
				slot.Runtime() == req.Runtime &&
				slot.LoraDir() == req.LoraDir &&
				!slot.IsActive() {
				existingCount++
			}
			return true
		})

		// If we need more slots, start creating them
		slotsNeeded := req.Count - existingCount
		if slotsNeeded > 0 {
			s.ensureSlots(req, slotsNeeded)
		}
	}
}

func (s *Scheduler) ensureSlots(req SlotRequirement, count int) {
	for i := 0; i < count; i++ {
		// Find best runner for this slot
		runnerID, err := s.pickBestRunner(req.ExampleWorkload)
		if err != nil {
			log.Warn().Err(err).Interface("requirement", req).Msg("failed to pick best runner for requirement")
			continue
		}

		// Delete any stale slots on this runner if required
		err = s.deleteMostStaleStrategy(runnerID, req.ExampleWorkload.Model().GetMemoryRequirements(req.ExampleWorkload.Mode()))
		if err != nil {
			log.Warn().Err(err).Str("runner_id", runnerID).Msg("failed to delete stale slots")
			continue
		}

		// Queue slot creation
		pending := &PendingSlot{
			RunnerID: runnerID,
			Created:  make(chan struct{}),
			Work:     req.ExampleWorkload,
		}

		select {
		case s.pendingSlots <- pending:
			log.Debug().Str("runner_id", runnerID).Str("model_name", req.ExampleWorkload.ModelName().String()).Msg("queued slot creation")
		default:
			log.Debug().Msg("pending slots channel full, skipping...")
			return
		}
	}
}

func (s *Scheduler) processQueueOnce() {
	// Try to take next work item that has a warm slot available
	work := s.queue.TakeNext(func(w *Workload) bool {
		warmSlots := s.warmSlots(w)
		return len(warmSlots) > 0
	})

	if work == nil {
		return // Nothing can be scheduled right now
	}

	// We know we have a warm slot, so schedule the work
	warmSlots := s.warmSlots(work)
	slot := s.pickBestWarmSlot(warmSlots)

	err := s.allocateSlot(slot.ID, work)
	if err != nil {
		// If allocation fails, put work back in queue
		err = s.queue.Add(work)
		if err != nil {
			log.Warn().Err(err).Msg("failed to add work back to queue")
			s.onSchedulingErr(work, err)
		}
	}
}

// Add new method to handle slot creation
func (s *Scheduler) runSlotCreator(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case pending := <-s.pendingSlots:
			// Create an allocated slot, lock the reconciler to prevent it from running until the slot
			// is created
			withWorkContext(&log.Logger, pending.Work).Trace().Msg("taking slot mutex")
			err := s.allocateNewSlot(ctx, pending.RunnerID, pending.Work)
			withWorkContext(&log.Logger, pending.Work).Trace().Msg("returned slot mutex")
			if err != nil {
				withWorkContext(&log.Logger, pending.Work).Error().Err(err).Msg("failed to create slot, calling error handler")
				s.onSchedulingErr(pending.Work, err)
			}
			close(pending.Created) // Signal that creation attempt is complete
		}
	}
}

// Add new helper method to find the best runner
func (s *Scheduler) pickBestRunner(work *Workload) (string, error) {
	// First get a list of all runners
	allRunners := s.controller.RunnerIDs()

	// Reach out to each runner and get their total memory
	runnerMemory := make(map[string]uint64)
	for _, runnerID := range allRunners {
		runnerMemory[runnerID] = s.controller.TotalMemory(runnerID)
	}
	withWorkContext(&log.Logger, work).Debug().Interface("runner_memory", runnerMemory).Msg("runner memory")

	// Filter out runners that don't have enough memory to allocate the new workload
	numRunnersWithNotEnoughTotalMemory := 0
	largestRunnerMemory := uint64(0)
	requiredMemory := work.Model().GetMemoryRequirements(work.Mode())
	filteredRunners := make([]string, 0)
	for runnerID, memory := range runnerMemory {
		if memory >= requiredMemory {
			filteredRunners = append(filteredRunners, runnerID)
		} else {
			numRunnersWithNotEnoughTotalMemory++
		}
		if memory > largestRunnerMemory {
			largestRunnerMemory = memory
		}
	}
	withWorkContext(&log.Logger, work).Debug().Interface("filtered_runners", filteredRunners).Msg("filtered runners")

	// Error if no runners have enough memory
	if numRunnersWithNotEnoughTotalMemory == len(allRunners) {
		return "", fmt.Errorf("no runner has enough GPU memory for this workload (desired: %d, largest: %d): %w", requiredMemory, largestRunnerMemory, ErrModelWontFit)
	}

	// Error if there are no runners left
	if len(filteredRunners) == 0 {
		return "", ErrNoRunnersAvailable
	}

	// Calculate the scheduled load on each runner according to their slots
	// Note: We discussed using real free memory by pinging the runner, but decided to use the
	// control-plane's view of free memory to avoid the overhead of pinging the runner. This also
	// has the added benefit of being able to over-commit memory slightly.
	runnerLoad := make(map[string]uint64)
	for _, runnerID := range filteredRunners {
		// Sum up all scheduled slots on the runner
		s.slots.Range(func(_ uuid.UUID, slot *Slot) bool {
			if slot.RunnerID == runnerID {
				runnerLoad[runnerID] += slot.Memory()
			}
			return true
		})
	}
	withWorkContext(&log.Logger, work).Debug().Interface("runner_load", runnerLoad).Msg("runner load")

	// Sort the runners by load, increasing, with a random shuffle for ties
	slices.SortFunc(filteredRunners, func(a, b string) int {
		if runnerLoad[a] != runnerLoad[b] {
			return int(runnerLoad[a] - runnerLoad[b])
		}
		return rand.Intn(3) - 1 // Introduces random shuffle for true ties
	})
	withWorkContext(&log.Logger, work).Debug().Interface("sorted_runners", filteredRunners).Msg("sorted runners")

	// Pick the first runner
	bestRunnerID := filteredRunners[0]
	withWorkContext(&log.Logger, work).Debug().Str("runner_id", bestRunnerID).Msg("chosen best runner")

	// Return the bestRunnerID and any error
	return bestRunnerID, nil
}

// DeleteMostStaleStrategy iteratively deletes allocated work from stale slots until there is enough
// memory to allocate the new workload.
func (s *Scheduler) deleteMostStaleStrategy(runnerID string, requiredMem uint64) error {
	totalMem := s.controller.TotalMemory(runnerID)
	for {
		var allSlots []*Slot
		allocatedMem := uint64(0)
		s.slots.Range(func(_ uuid.UUID, slot *Slot) bool {
			if slot.RunnerID == runnerID {
				allSlots = append(allSlots, slot)
				allocatedMem += slot.Memory()
			}
			return true
		})
		staleSlots := Filter(allSlots, func(slot *Slot) bool {
			return slot.IsStale()
		})
		freeMem := int64(totalMem) - int64(allocatedMem) - int64(requiredMem)
		log.Trace().Interface("slots", allSlots).Int64("freeMem", freeMem).Msg("checking if we can allocate")
		// If there is enough free space on the runner, break out of the loop.
		if freeMem > 0 {
			break
		}

		// Sort the slots by last activity time
		slices.SortFunc(staleSlots, func(i, j *Slot) int {
			return int(i.LastActivityTime.Sub(j.LastActivityTime))
		})
		log.Trace().Interface("stale_slots", staleSlots).Msg("stale slots")
		if len(staleSlots) == 0 {
			return ErrRunnersAreFull
		}
		// Then delete the most stale slot
		withSlotContext(&log.Logger, staleSlots[0]).Info().Msg("deleting stale slot")
		err := s.controller.DeleteSlot(runnerID, staleSlots[0].ID)
		if err != nil {
			return fmt.Errorf("unable to delete stale slot: %w", err)
		}
		s.slots.Delete(staleSlots[0].ID)
	}
	return nil
}

func (s *Scheduler) warmSlots(req *Workload) []*Slot {
	cosyWarm := make([]*Slot, 0, s.slots.Size())
	s.slots.Range(func(_ uuid.UUID, slot *Slot) bool {
		// If it's not the same model name, skip
		if slot.ModelName() != req.ModelName() {
			withSlotContext(&log.Logger, slot).Trace().Msg("skipping warm slot, model name mismatch")
			return true
		}

		// If it's not the same runtime, skip
		if slot.ModelName().InferenceRuntime() != req.ModelName().InferenceRuntime() {
			withSlotContext(&log.Logger, slot).Trace().Msg("skipping warm slot, inference runtime mismatch")
			return true
		}

		// If the slot is already running another job, skip
		if slot.IsActive() {
			withSlotContext(&log.Logger, slot).Trace().Msg("skipping warm slot, already active")
			return true
		}

		// If it doesn't have the right LoraDir then skip
		if slot.LoraDir() != req.LoraDir() {
			withSlotContext(&log.Logger, slot).Trace().Str("slot_lora_dir", slot.LoraDir()).Str("req_lora_dir", req.LoraDir()).Msg("skipping warm slot, LoraDir mismatch")
			return true
		}

		// Add available slots to the list.
		cosyWarm = append(cosyWarm, slot)
		return true
	})
	return cosyWarm
}

func (s *Scheduler) pickBestWarmSlot(warmSlots []*Slot) *Slot {
	// If we only have one slot, return it
	if len(warmSlots) == 1 {
		return warmSlots[0]
	}

	// Group slots by runner
	runnerSlots := make(map[string][]*Slot)
	for _, slot := range warmSlots {
		runnerSlots[slot.RunnerID] = append(runnerSlots[slot.RunnerID], slot)
	}

	// Count active slots per runner
	activeSlots := make(map[string]int)
	s.slots.Range(func(_ uuid.UUID, slot *Slot) bool {
		if slot.IsActive() {
			activeSlots[slot.RunnerID]++
		}
		return true
	})

	// Sort slots considering:
	// 1. Runner load (prefer less loaded runners)
	// 2. Last activity time (prefer more recently used slots)
	// 3. Random factor for tie-breaking
	slices.SortFunc(warmSlots, func(i, j *Slot) int {
		// First compare runner load
		if activeSlots[i.RunnerID] != activeSlots[j.RunnerID] {
			return activeSlots[i.RunnerID] - activeSlots[j.RunnerID]
		}

		// Then prefer more recently used slots (reverse of current order)
		if !i.LastActivityTime.Equal(j.LastActivityTime) {
			return int(j.LastActivityTime.Sub(i.LastActivityTime))
		}

		// Random tie-breaker
		return rand.Intn(3) - 1
	})

	return warmSlots[0]
}

// AllocateSlot assigns a workload to a specific slot, validating the model and slot before scheduling.
func (s *Scheduler) allocateSlot(slotID uuid.UUID, req *Workload) error {
	// Validate slot
	slot, ok := s.slots.Load(slotID)
	if !ok {
		return fmt.Errorf("slot not found: %s", slot.ID.String())
	}

	// Ensure the slot is not already scheduled or active.
	if slot.IsActive() {
		return fmt.Errorf("slot already active: %s", slot.ID.String())
	}

	// Marks the slot as locally active. This is reset in the reconciliation process.
	withSlotAndWorkContext(&log.Logger, slot, req).Trace().Msg("starting slot")
	slot.Start()

	// Can do the rest in a goroutine, no need to wait for it to submit
	go func() {
		// Submit the work to the slot
		switch req.WorkloadType {
		case WorkloadTypeLLMInferenceRequest:
			withSlotAndWorkContext(&log.Logger, slot, req).Trace().Msg("submitting chat completion request")
			err := s.controller.SubmitChatCompletionRequest(slot, req.LLMInferenceRequest())
			if err != nil {
				// TODO(Phil): Need to pass on the error to the session for all these cases
				log.Error().Err(err).Msg("error submitting chat completion request")
			}
		case WorkloadTypeSession:
			switch req.Session().Mode {
			case types.SessionModeInference:
				switch req.Session().Type {
				case types.SessionTypeImage:
					withSlotAndWorkContext(&log.Logger, slot, req).Trace().Msg("submitting text2image request")
					err := s.controller.SubmitImageGenerationRequest(slot, req.Session())
					if err != nil {
						log.Error().Err(err).Msg("error submitting text2image request")
					}
				case types.SessionTypeText:
					if req.Session().LoraDir != "" {
						withSlotAndWorkContext(&log.Logger, slot, req).Trace().Msg("submitting LORA text inference request")
						// Overwrite the request model name with the helix lora model details
						convertedRequest := req.ToLLMInferenceRequest()
						convertedRequest.Request.Model = req.Session().LoraDir

						// Forward the request to the chat completion handler
						err := s.controller.SubmitChatCompletionRequest(slot, convertedRequest)
						if err != nil {
							log.Error().Err(err).Msg("error submitting text request")
						}
					} else {
						panic(fmt.Sprintf("not implemented: %s and no lora dir", req.Session().Type))
					}
				default:
					panic(fmt.Sprintf("not implemented: %s", req.Session().Type))
				}
			case types.SessionModeFinetune:
				switch req.Session().Type {
				case types.SessionTypeText:
					withSlotAndWorkContext(&log.Logger, slot, req).Trace().Msg("submitting finetuning request")
					err := s.controller.SubmitFinetuningRequest(slot, req.Session())
					if err != nil {
						log.Error().Err(err).Msg("error submitting finetuning request")
					}
				default:
					panic(fmt.Sprintf("not implemented: %s", req.Session().Type))
				}
			default:
				panic(fmt.Sprintf("not implemented: %s", req.Session().Mode))
			}
		}

		withSlotAndWorkContext(&log.Logger, slot, req).Trace().Msg("finished submitting request")
	}()

	return nil
}

// AllocateNewSlot creates a new slot for a workload and allocates it to the best available runner.
func (s *Scheduler) allocateNewSlot(ctx context.Context, runnerID string, req *Workload) error {
	// Create a new slot and schedule the workload.
	slot := NewSlot(runnerID, req, s.modelStaleFunc, s.slotTimeoutFunc)
	withSlotAndWorkContext(&log.Logger, slot, req).Info().Msg("creating new slot")

	err := s.controller.CreateSlot(slot)
	if err != nil {
		return err
	}

	// Wait for the slot to be ready
	slotReady := make(chan bool)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-time.After(100 * time.Millisecond):
				slots, err := s.controller.GetSlots(slot.RunnerID)
				if err != nil {
					log.Error().Err(err).Msg("unable to get slots")
					return
				}
				for _, s := range slots {
					if s.ID == slot.ID {
						slotReady <- true
						return
					}
				}
			}
		}
	}()
	select {
	case <-slotReady:
	case <-time.After(120 * time.Second):
		return fmt.Errorf("slot not ready after 120 seconds")
	}

	withSlotAndWorkContext(&log.Logger, slot, req).Info().Msg("slot created")

	// Ensure the slot is stored.
	s.slots.Store(slot.ID, slot)

	// Schedule and store the new slot.
	return s.allocateSlot(slot.ID, req)
}

// AddWorkFields adds standard work-related fields to a log event
func withWorkContext(l *zerolog.Logger, w *Workload) *zerolog.Logger {
	nextLogger := l.With().
		Str("work_id", w.ID()).
		Str("model_name", w.ModelName().String()).
		Str("mode", string(w.Mode())).
		Str("runtime", string(w.Runtime())).
		Str("lora_dir", w.LoraDir()).
		Dur("duration_ms", time.Since(w.Created())).
		Logger()
	return &nextLogger
}

// AddSlotFields adds standard slot-related fields to a log event
func withSlotContext(l *zerolog.Logger, s *Slot) *zerolog.Logger {
	nextLogger := l.With().
		Str("runner_id", s.RunnerID).
		Str("slot_id", s.ID.String()).
		Str("model_name", s.ModelName().String()).
		Uint64("memory", s.Memory()).
		Logger()
	return &nextLogger
}

func withSlotAndWorkContext(l *zerolog.Logger, s *Slot, w *Workload) *zerolog.Logger {
	return withSlotContext(withWorkContext(l, w), s)
}
