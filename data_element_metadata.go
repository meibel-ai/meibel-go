package meibelgo

import (
	"context"
	"fmt"
)

// DataElementMetadataService handles Data Element Metadata operations.
type DataElementMetadataService struct {
	client *MeibelgoClient
}

// GetDataElementMetadata Get Data Element Metadata
func (s *DataElementMetadataService) GetDataElementMetadata(ctx context.Context, datasourceId string, dataElementId string) (*string, error) {
	path := "/v2/datasources/" + fmt.Sprintf("%v", datasourceId) + "/data-elements/" + fmt.Sprintf("%v", dataElementId) + "/metadata"

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

// UpdateDataElementMetadata Update Data Element Metadata
func (s *DataElementMetadataService) UpdateDataElementMetadata(ctx context.Context, datasourceId string, dataElementId string, body DataElementMetadata) (*string, error) {
	path := "/v2/datasources/" + fmt.Sprintf("%v", datasourceId) + "/data-elements/" + fmt.Sprintf("%v", dataElementId) + "/metadata"

	var result string
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

// GetDataElementMetadataResult Get Data Element Metadata Result
func (s *DataElementMetadataService) GetDataElementMetadataResult(ctx context.Context, datasourceId string, dataElementId string, requestId string) (*string, error) {
	path := "/v2/datasources/" + fmt.Sprintf("%v", datasourceId) + "/data-elements/" + fmt.Sprintf("%v", dataElementId) + "/metadata/result/" + fmt.Sprintf("%v", requestId)

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
