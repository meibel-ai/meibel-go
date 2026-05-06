package meibelgo

import (
	"context"
	"fmt"
)

// DatasourcesService handles datasources operations.
type DatasourcesService struct {
	client *MeibelgoClient
	Content *ContentService
	Dataelements *DataelementsService
	Rag *RagService
	Tag *TagService
}

// AddDatasource Add Datasource
func (s *DatasourcesService) AddDatasource(ctx context.Context, body AddDatasourceRequest) (*AddDatasourceResponse, error) {
	path := "/datasource"

	var result AddDatasourceResponse
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
func (s *DatasourcesService) UpdateDatasource(ctx context.Context, datasourceId string, body UpdateDatasourceRequest) (*UpdateDatasourceResponse, error) {
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

// GetAllDatasourceIds Get All Datasource Ids
func (s *DatasourcesService) GetAllDatasourceIds(ctx context.Context) (*GetAllDatasourceIdsResponse, error) {
	path := "/project_datasource_ids"

	var result GetAllDatasourceIdsResponse
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "POST",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
