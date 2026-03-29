package meibelgo

import (
	"context"
	"fmt"
)

// MetadataConfigurationService handles Metadata Configuration operations.
type MetadataConfigurationService struct {
	client *MeibelgoClient
}

// GetMetadataConfig Get Metadata Config
func (s *MetadataConfigurationService) GetMetadataConfig(ctx context.Context, datasourceId string) (*MetadataConfigResponse, error) {
	path := "/datasources/" + fmt.Sprintf("%v", datasourceId) + "/metadata-config"

	var result MetadataConfigResponse
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "GET",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateMetadataConfig Update Metadata Config
func (s *MetadataConfigurationService) UpdateMetadataConfig(ctx context.Context, datasourceId string, body MetadataConfigRequest) (*MetadataConfigResponse, error) {
	path := "/datasources/" + fmt.Sprintf("%v", datasourceId) + "/metadata-config"

	var result MetadataConfigResponse
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "PUT",
		Path:   path,
		Body:   body,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// ReprocessMetadata Reprocess Metadata
func (s *MetadataConfigurationService) ReprocessMetadata(ctx context.Context, datasourceId string) (*string, error) {
	path := "/datasources/" + fmt.Sprintf("%v", datasourceId) + "/reprocess-metadata"

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

// GetReprocessMetadataStatus Get Reprocess Metadata Status
func (s *MetadataConfigurationService) GetReprocessMetadataStatus(ctx context.Context, datasourceId string) (*string, error) {
	path := "/datasources/" + fmt.Sprintf("%v", datasourceId) + "/reprocess-metadata/status"

	var result string
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "GET",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
