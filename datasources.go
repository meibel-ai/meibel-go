package v2

import (
	"context"
	"fmt"
	"net/url"
)

// DatasourcesService handles Datasources operations.
type DatasourcesService struct {
	client *MeibelClient
	DataElements *DataElementsService
	Downloads *DownloadsService
	FileUploads *FileUploadsService
	Tables *TablesService
}

// GetDatasourceOptions contains optional parameters for GetDatasource.
type GetDatasourceOptions struct {
	// Include table and column details (structured datasources only)
	IncludeTables *bool
}

// ListDatasources List Datasources
func (s *DatasourcesService) ListDatasources(ctx context.Context) (*DatasourceListResponse, error) {
	path := "/datasources"

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
func (s *DatasourcesService) GetDatasource(ctx context.Context, datasourceId string, opts *GetDatasourceOptions) (*DatasourceResponse, error) {
	path := "/datasources/" + fmt.Sprintf("%v", datasourceId)
	query := url.Values{}
	if opts != nil && opts.IncludeTables != nil {
		query.Set("include_tables", fmt.Sprintf("%v", *opts.IncludeTables))
	}

	var result DatasourceResponse
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "GET",
		Path:   path,
		Query:  query,
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
