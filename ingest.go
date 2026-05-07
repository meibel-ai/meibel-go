package meibelgo

import (
	"context"
	"fmt"
)

// IngestService handles Ingest operations.
type IngestService struct {
	client *MeibelgoClient
}

// TriggerIngest Trigger Ingest
func (s *IngestService) TriggerIngest(ctx context.Context, datasourceId string) (*string, error) {
	path := "/datasources/" + fmt.Sprintf("%v", datasourceId) + "/trigger-ingest"

	var result string
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "POST",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// GetIngestStatus Get Ingest Status
func (s *IngestService) GetIngestStatus(ctx context.Context, datasourceId string) (*IngestStatusResponse, error) {
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
