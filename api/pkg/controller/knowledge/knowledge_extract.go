package knowledge

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/helixml/helix/api/pkg/extract"
	"github.com/helixml/helix/api/pkg/types"
)

func (r *Reconciler) getIndexingData(ctx context.Context, k *types.Knowledge) ([]*indexerData, error) {
	switch {
	case k.Source.Web != nil:
		return r.extractDataFromWeb(ctx, k)
	case k.Source.HelixDrive != nil:
		return r.extractDataFromHelixDrive(ctx, k)
	default:
		return nil, fmt.Errorf("unknown source")
	}
}

func (r *Reconciler) extractDataFromWeb(ctx context.Context, k *types.Knowledge) ([]*indexerData, error) {
	if k.Source.Web == nil {
		return nil, fmt.Errorf("no web source defined")
	}

	var result []*indexerData

	if len(k.Source.Web.URLs) == 0 {
		return result, nil
	}

	// Optional mode to disable text extractor and chunking,
	// useful when the indexing server will know how to handle
	// raw data directly
	extractorEnabled := true

	if k.RAGSettings.DisableChunking {
		extractorEnabled = false
	}

	// TODO: add concurrency
	for _, u := range k.Source.Web.URLs {
		if extractorEnabled {
			extracted, err := r.extractor.Extract(ctx, &extract.ExtractRequest{
				URL: u,
			})
			if err != nil {
				return nil, fmt.Errorf("failed to extract data from %s, error: %w", u, err)
			}

			result = append(result, &indexerData{
				Data:   []byte(extracted),
				Source: u,
			})

			continue
		}

		bts, err := r.downloadDirectly(ctx, k, u)
		if err != nil {
			return nil, fmt.Errorf("failed to download data from %s, error: %w", u, err)
		}

		result = append(result, &indexerData{
			Data:   bts,
			Source: u,
		})
	}

	return result, nil
}

func (r *Reconciler) downloadDirectly(ctx context.Context, k *types.Knowledge, u string) ([]byte, error) {
	// Extractor and indexer disabled, downloading directly
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("failed to get %s, error: %w", u, err)
	}

	// If username and password are specified, use them for basic auth
	if k.Source.Web.Auth.Username != "" || k.Source.Web.Auth.Password != "" {
		req.SetBasicAuth(k.Source.Web.Auth.Username, k.Source.Web.Auth.Password)
	}

	resp, err := r.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to download, error: %w", err)
	}
	defer resp.Body.Close()

	bts, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body, error: %w", err)
	}

	return bts, nil
}

func (r *Reconciler) extractDataFromHelixDrive(ctx context.Context, k *types.Knowledge) ([]*indexerData, error) {
	return nil, fmt.Errorf("TODO")
}