package meibelgo

import (
	"context"
	"fmt"
)

// DataElementsService handles DataElements operations.
type DataElementsService struct {
	client *MeibelClient
}

// GetDataElement Get Data Element
func (s *DataElementsService) GetDataElement(ctx context.Context, datasourceId string, dataElementId string) (*DataElement, error) {
	path := "/datasource/" + fmt.Sprintf("%v", datasourceId) + "/data_element/" + fmt.Sprintf("%v", dataElementId)

	var result DataElement
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "GET",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateDataElement Update Data Element
func (s *DataElementsService) UpdateDataElement(ctx context.Context, datasourceId string, dataElementId string, body DatasourceServiceClientModelsUpdateDataElementRequestUpdateDataElementRequest) (*UpdateDataElementResponse, error) {
	path := "/datasource/" + fmt.Sprintf("%v", datasourceId) + "/data_element/" + fmt.Sprintf("%v", dataElementId)

	var result UpdateDataElementResponse
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

// ListDataElements List Data Elements
func (s *DataElementsService) ListDataElements(ctx context.Context, datasourceId string) (*[]DataElementResponse, error) {
	path := "/v2/datasources/" + fmt.Sprintf("%v", datasourceId) + "/data-elements"

	var result []DataElementResponse
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "GET",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// GetDataElement Get Data Element
func (s *DataElementsService) GetDataElement(ctx context.Context, datasourceId string, dataElementId string) (*DataElementResponse, error) {
	path := "/v2/datasources/" + fmt.Sprintf("%v", datasourceId) + "/data-elements/" + fmt.Sprintf("%v", dataElementId)

	var result DataElementResponse
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "GET",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateDataElement Update Data Element
func (s *DataElementsService) UpdateDataElement(ctx context.Context, datasourceId string, dataElementId string, body GatewayServiceV2ModelsDataElementsUpdateDataElementRequest) (*DataElementResponse, error) {
	path := "/v2/datasources/" + fmt.Sprintf("%v", datasourceId) + "/data-elements/" + fmt.Sprintf("%v", dataElementId)

	var result DataElementResponse
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

// SearchDataElements Search Data Elements
func (s *DataElementsService) SearchDataElements(ctx context.Context, datasourceId string, body DataElementSearchRequest) (*[]DataElementResponse, error) {
	path := "/v2/datasources/" + fmt.Sprintf("%v", datasourceId) + "/data-elements/search"

	var result []DataElementResponse
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
