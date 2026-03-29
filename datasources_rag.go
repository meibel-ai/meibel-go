package meibelgo

import (
	"context"
	"fmt"
)

// DatasourcesRagService handles datasources.rag operations.
type DatasourcesRagService struct {
	client *MeibelgoClient
}

// GetRagConfig Get Rag Config
func (s *DatasourcesRagService) GetRagConfig(ctx context.Context, datasourceId string) (*RagConfig, error) {
	path := "/datasource/" + fmt.Sprintf("%v", datasourceId) + "/rag_config"

	var result RagConfig
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "GET",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// AddRagConfig Add Rag Config
func (s *DatasourcesRagService) AddRagConfig(ctx context.Context, datasourceId string, body AddRagConfigRequest) (*AddRagConfigResponse, error) {
	path := "/datasource/" + fmt.Sprintf("%v", datasourceId) + "/rag_config"

	var result AddRagConfigResponse
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "POST",
		Path:   path,
		Body:   body,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateRagConfig Update Rag Config
func (s *DatasourcesRagService) UpdateRagConfig(ctx context.Context, datasourceId string, body UpdateRagConfigRequest) (*UpdateRagConfigResponse, error) {
	path := "/datasource/" + fmt.Sprintf("%v", datasourceId) + "/rag_config"

	var result UpdateRagConfigResponse
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

// DeleteRagConfig Delete Rag Config
func (s *DatasourcesRagService) DeleteRagConfig(ctx context.Context, datasourceId string) (*string, error) {
	path := "/datasource/" + fmt.Sprintf("%v", datasourceId) + "/rag_config"

	var result string
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "DELETE",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// GetChunkingStrategy Get Chunking Strategy
func (s *DatasourcesRagService) GetChunkingStrategy(ctx context.Context, datasourceId string) (*RagChunkingStrategy, error) {
	path := "/datasource/" + fmt.Sprintf("%v", datasourceId) + "/chunking_strategy"

	var result RagChunkingStrategy
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "GET",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// AddChunkingStrategy Add Chunking Strategy
func (s *DatasourcesRagService) AddChunkingStrategy(ctx context.Context, datasourceId string, body AddChunkingStrategyRequest) (*AddChunkingStrategyResponse, error) {
	path := "/datasource/" + fmt.Sprintf("%v", datasourceId) + "/chunking_strategy"

	var result AddChunkingStrategyResponse
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "POST",
		Path:   path,
		Body:   body,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateChunkingStrategy Update Chunking Strategy
func (s *DatasourcesRagService) UpdateChunkingStrategy(ctx context.Context, datasourceId string, body UpdateChunkingStrategyRequest) (*UpdateChunkingStrategyResponse, error) {
	path := "/datasource/" + fmt.Sprintf("%v", datasourceId) + "/chunking_strategy"

	var result UpdateChunkingStrategyResponse
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

// DeleteChunkingStrategy Delete Chunking Strategy
func (s *DatasourcesRagService) DeleteChunkingStrategy(ctx context.Context, datasourceId string) (*DeleteChunkingStrategyResponse, error) {
	path := "/datasource/" + fmt.Sprintf("%v", datasourceId) + "/chunking_strategy"

	var result DeleteChunkingStrategyResponse
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "DELETE",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// ReprocessDatasourceMetadata Reprocess Datasource Metadata
func (s *DatasourcesRagService) ReprocessDatasourceMetadata(ctx context.Context, datasourceId string, body ReprocessDatasourceRequest) (*ReprocessDatasourceResponse, error) {
	path := "/datasource/" + fmt.Sprintf("%v", datasourceId) + "/reprocess_metadata"

	var result ReprocessDatasourceResponse
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "POST",
		Path:   path,
		Body:   body,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// GetReprocessMetadataStatus Get Reprocess Metadata Status
func (s *DatasourcesRagService) GetReprocessMetadataStatus(ctx context.Context, datasourceId string) (*GetReprocessStatusResponse, error) {
	path := "/datasource/" + fmt.Sprintf("%v", datasourceId) + "/reprocess_metadata/status"

	var result GetReprocessStatusResponse
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "GET",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// PatchRagConfigMetadata Patch Rag Config Metadata
func (s *DatasourcesRagService) PatchRagConfigMetadata(ctx context.Context, datasourceId string, body PatchRagConfigMetadataRequest) (*PatchRagConfigMetadataResponse, error) {
	path := "/datasource/" + fmt.Sprintf("%v", datasourceId) + "/rag_config/metadata"

	var result PatchRagConfigMetadataResponse
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "PATCH",
		Path:   path,
		Body:   body,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
