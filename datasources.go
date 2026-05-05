package meibelgo

import (
	"context"
	"fmt"
)

// DatasourcesService handles Datasources operations.
type DatasourcesService struct {
	client *MeibelClient
	Content *ContentService
	DataElements *DataElementsService
}

// GetDatasource Get Datasource
func (s *DatasourcesService) GetDatasource(ctx context.Context, datasourceId string) (*Datasource, error) {
	path := "/datasource/" + fmt.Sprintf("%v", datasourceId)

	var result Datasource
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "GET",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateDatasource Update Datasource
func (s *DatasourcesService) UpdateDatasource(ctx context.Context, datasourceId string, body DatasourceServiceClientModelsUpdateDatasourceRequestUpdateDatasourceRequest) (*UpdateDatasourceResponse, error) {
	path := "/datasource/" + fmt.Sprintf("%v", datasourceId)

	var result UpdateDatasourceResponse
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

// DeleteDatasource Delete Datasource
func (s *DatasourcesService) DeleteDatasource(ctx context.Context, datasourceId string) (*DeleteDatasourceResponse, error) {
	path := "/datasource/" + fmt.Sprintf("%v", datasourceId)

	var result DeleteDatasourceResponse
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "DELETE",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// ListDatasources List Datasources
func (s *DatasourcesService) ListDatasources(ctx context.Context) (*DatasourceListResponse, error) {
	path := "/v2/datasources"

	var result DatasourceListResponse
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "GET",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// CreateDatasource Create Datasource
func (s *DatasourcesService) CreateDatasource(ctx context.Context, body CreateDatasourceRequest) (*DatasourceResponse, error) {
	path := "/v2/datasources"

	var result DatasourceResponse
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

// GetDatasource Get Datasource
func (s *DatasourcesService) GetDatasource(ctx context.Context, datasourceId string) (*DatasourceResponse, error) {
	path := "/v2/datasources/" + fmt.Sprintf("%v", datasourceId)

	var result DatasourceResponse
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "GET",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateDatasource Update Datasource
func (s *DatasourcesService) UpdateDatasource(ctx context.Context, datasourceId string, body GatewayServiceV2ModelsDatasourcesUpdateDatasourceRequest) (*DatasourceResponse, error) {
	path := "/v2/datasources/" + fmt.Sprintf("%v", datasourceId)

	var result DatasourceResponse
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

// DeleteDatasource Delete Datasource
func (s *DatasourcesService) DeleteDatasource(ctx context.Context, datasourceId string) (*string, error) {
	path := "/v2/datasources/" + fmt.Sprintf("%v", datasourceId)

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
