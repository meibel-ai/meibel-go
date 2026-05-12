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
	Ingest *IngestService
	Tables *TablesService
}

// DatasourcesGetOptions contains optional parameters for Get.
type DatasourcesGetOptions struct {
	// Include table and column details (structured datasources only)
	IncludeTables *bool
}

// List List Datasources
func (s *DatasourcesService) List(ctx context.Context) (*DatasourceListResponse, error) {
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

// Create Create Datasource
func (s *DatasourcesService) Create(ctx context.Context, body CreateDatasourceRequest) (*DatasourceResponse, error) {
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

// Get Get Datasource
func (s *DatasourcesService) Get(ctx context.Context, datasourceId string, opts *DatasourcesGetOptions) (*DatasourceResponse, error) {
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

// Update Update Datasource
func (s *DatasourcesService) Update(ctx context.Context, datasourceId string, body UpdateDatasourceRequest) (*DatasourceResponse, error) {
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

// Delete Delete Datasource
func (s *DatasourcesService) Delete(ctx context.Context, datasourceId string) (*string, error) {
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
