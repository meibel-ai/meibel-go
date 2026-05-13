package v2

import (
	"context"
	"fmt"
)

// IngestService handles Ingest operations.
type IngestService struct {
	client *MeibelClient
}

// Trigger Trigger Ingest
func (s *IngestService) Trigger(ctx context.Context, datasourceId string) (*TriggerIngestResponse, error) {
	path := "/datasources/" + fmt.Sprintf("%v", datasourceId) + "/trigger-ingest"

	var result TriggerIngestResponse
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "POST",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// GetStatus Get Ingest Status
func (s *IngestService) GetStatus(ctx context.Context, datasourceId string) (*IngestStatusResponse, error) {
	path := "/datasources/" + fmt.Sprintf("%v", datasourceId) + "/ingest-status"

	var result IngestStatusResponse
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "GET",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
