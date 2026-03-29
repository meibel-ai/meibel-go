package meibelgo

import (
	"context"
	"fmt"
	"net/url"
)

// DatasourcesService handles Datasources operations.
type DatasourcesService struct {
	client *MeibelgoClient
}

// ListDatasources List Datasources
func (s *DatasourcesService) ListDatasources(ctx context.Context) *PageIterator[DatasourceResponse] {
	path := "/datasources"
	query := url.Values{}

	return NewPageIterator(func(ctx context.Context, cursor string) (*Page[DatasourceResponse], error) {
		if cursor != "" {
			query.Set("offset", cursor)
		}

		var resp struct {
			Datasources []DatasourceResponse `json:"datasources"`
			NextCursor  string               `json:"next_cursor"`
		}

		err := s.client.http.Do(ctx, RequestOptions{
			Method: "GET",
			Path:   path,
			Query:  query,
		}, &resp)
		if err != nil {
			return nil, err
		}

		return &Page[DatasourceResponse]{
			Items:      resp.Datasources,
			NextCursor: resp.NextCursor,
		}, nil
	})
}

// CreateDatasource Create Datasource
func (s *DatasourcesService) CreateDatasource(ctx context.Context, body CreateDatasourceRequest) (*DatasourceResponse, error) {
	path := "/datasources"

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
	path := "/datasources/" + fmt.Sprintf("%v", datasourceId)

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
func (s *DatasourcesService) UpdateDatasource(ctx context.Context, datasourceId string, body UpdateDatasourceRequest) (*DatasourceResponse, error) {
	path := "/datasources/" + fmt.Sprintf("%v", datasourceId)

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
	path := "/datasources/" + fmt.Sprintf("%v", datasourceId)

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
