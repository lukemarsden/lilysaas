package types

import (
	"context"
	"time"

	"github.com/bacalhau-project/lilypad/pkg/data"
)

type OwnerType string

const (
	OwnerTypeUser OwnerType = "user"
)

type PaymentType string

const (
	PaymentTypeAdmin  PaymentType = "admin"
	PaymentTypeStripe PaymentType = "stripe"
	PaymentTypeJob    PaymentType = "job"
)

type JobSpec struct {
	Module string            `json:"module"`
	Inputs map[string]string `json:"inputs"`
}

type JobData struct {
	Spec      JobSpec                `json:"spec"`
	Container data.JobOfferContainer `json:"container"`
}

type Job struct {
	ID        string    `json:"id"`
	Created   time.Time `json:"created"`
	Owner     string    `json:"owner"`
	OwnerType OwnerType `json:"owner_type"`
	State     string    `json:"state"`
	Status    string    `json:"status"`
	Data      JobData   `json:"data"`
}

type BalanceTransferData struct {
	JobID           string `json:"job_id"`
	StripePaymentID string `json:"stripe_payment_id"`
}

type BalanceTransfer struct {
	ID          string              `json:"id"`
	Created     time.Time           `json:"created"`
	Owner       string              `json:"owner"`
	OwnerType   OwnerType           `json:"owner_type"`
	PaymentType PaymentType         `json:"payment_type"`
	Amount      int                 `json:"amount"`
	Data        BalanceTransferData `json:"data"`
}

type Module struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Cost     int    `json:"cost"`
	Template string `json:"template"`
}

type UserMessage struct {
	User     string   `json:"user"`     // e.g. User
	Message  string   `json:"message"`  // e.g. Prove pythagoras
	Uploads  []string `json:"uploads"`  // list of filepath paths
	Finished bool     `json:"finished"` // if true, the message has finished being written to, and is ready for a response (e.g. from the other participant)
}

type Interactions struct {
	Messages []UserMessage `json:"messages"`
}

type Session struct {
	ID string `json:"id"`
	// name that goes in the UI - ideally autogenerated by AI but for now can be
	// named manually
	Name    string    `json:"name"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
	// e.g. create, finetune
	Mode string `json:"mode"`
	// e.g. text, images
	Type string `json:"type"`
	// huggingface model name e.g. mistralai/Mistral-7B-Instruct-v0.1 or
	// stabilityai/stable-diffusion-xl-base-1.0
	ModelName string `json:"model_name"`
	// if type == finetune, we record a filestore path to e.g. lora file here
	// currently the only place you can do inference on a finetune is within the
	// session where the finetune was generated
	FinetuneFile string `json:"finetune_file"`
	// for now we just whack the entire history of the interaction in here, json
	// style
	Interactions Interactions `json:"interactions"`
	// uuid of owner entity
	Owner string `json:"owner"`
	// e.g. user, system, org
	OwnerType OwnerType `json:"owner_type"`
}

// passed between the api server and the controller
type RequestContext struct {
	Ctx       context.Context
	Owner     string
	OwnerType OwnerType
}

type UserStatus struct {
	User    string `json:"user"`
	Credits int    `json:"credits"`
}

type WebsocketEventType string

const (
	WebsocketEventJobUpdate     WebsocketEventType = "job"
	WebsocketEventSessionUpdate WebsocketEventType = "session"
)

type WebsocketEvent struct {
	Type    WebsocketEventType `json:"type"`
	Job     *Job               `json:"job"`
	Session *Session           `json:"session"`
}
