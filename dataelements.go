package meibelgo

import (
	"context"
	"fmt"
	"net/url"
)

// DataelementsService handles dataelements operations.
type DataelementsService struct {
	client *MeibelgoClient
}

// GetDataElementsOptions contains optional parameters for GetDataElements.
type GetDataElementsOptions struct {
	// Number of items to skip
	Offset *int64
	// Maximum number of items to return
	Limit *int64
	// Field to sort by
	SortBy interface{}
	// Sort order (asc or desc)
	SortOrder interface{}
}

// GetDataElementsByFiltersOptions contains optional parameters for GetDataElementsByFilters.
type GetDataElementsByFiltersOptions struct {
	RegexFilter *string
	MediaTypeFilters []string
	// Number of items to skip
	Offset *int64
	// Maximum number of items to return
	Limit *int64
	// Field to sort by
	SortBy interface{}
	// Sort order (asc or desc)
	SortOrder interface{}
}

// GetDataElements Get Data Elements
func (s *DataelementsService) GetDataElements(ctx context.Context, datasourceId string, opts *GetDataElementsOptions) (*[]DataElement, error) {
	path := "/datasource/" + fmt.Sprintf("%v", datasourceId) + "/data_element"
	query := url.Values{}
	if opts != nil && opts.Offset != nil {
		query.Set("offset", fmt.Sprintf("%v", *opts.Offset))
	}
	if opts != nil && opts.Limit != nil {
		query.Set("limit", fmt.Sprintf("%v", *opts.Limit))
	}
	if opts != nil && opts.SortBy != nil {
		query.Set("sort_by", fmt.Sprintf("%v", opts.SortBy))
	}
	if opts != nil && opts.SortOrder != nil {
		query.Set("sort_order", fmt.Sprintf("%v", opts.SortOrder))
	}

	var result []DataElement
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

// AddDataElement Add Data Element
func (s *DataelementsService) AddDataElement(ctx context.Context, datasourceId string, body AddDataElementRequest) (*AddDataElementResponse, error) {
	path := "/datasource/" + fmt.Sprintf("%v", datasourceId) + "/data_element"

	var result AddDataElementResponse
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

// GetDataElement Get Data Element
func (s *DataelementsService) GetDataElement(ctx context.Context, datasourceId string, dataElementId string) (*DataElement, error) {
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
func (s *DataelementsService) UpdateDataElement(ctx context.Context, datasourceId string, dataElementId string, body UpdateDataElementRequest) (*UpdateDataElementResponse, error) {
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

// DeleteDataElement Delete Data Element
func (s *DataelementsService) DeleteDataElement(ctx context.Context, datasourceId string, dataElementId string) (*DeleteDataElementResponse, error) {
	path := "/datasource/" + fmt.Sprintf("%v", datasourceId) + "/data_element/" + fmt.Sprintf("%v", dataElementId)

	var result DeleteDataElementResponse
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "DELETE",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// GetDataElementsByFilters Get Data Elements By Filters
func (s *DataelementsService) GetDataElementsByFilters(ctx context.Context, datasourceId string, body *DataElementFilterRequest, opts *GetDataElementsByFiltersOptions) (*[]DataElement, error) {
	path := "/datasource/" + fmt.Sprintf("%v", datasourceId) + "/data_elements_by_filters"
	query := url.Values{}
	if opts != nil && opts.RegexFilter != nil {
		query.Set("regex_filter", fmt.Sprintf("%v", *opts.RegexFilter))
	}
	if opts != nil && opts.MediaTypeFilters != nil {
		query.Set("media_type_filters", fmt.Sprintf("%v", opts.MediaTypeFilters))
	}
	if opts != nil && opts.Offset != nil {
		query.Set("offset", fmt.Sprintf("%v", *opts.Offset))
	}
	if opts != nil && opts.Limit != nil {
		query.Set("limit", fmt.Sprintf("%v", *opts.Limit))
	}
	if opts != nil && opts.SortBy != nil {
		query.Set("sort_by", fmt.Sprintf("%v", opts.SortBy))
	}
	if opts != nil && opts.SortOrder != nil {
		query.Set("sort_order", fmt.Sprintf("%v", opts.SortOrder))
	}

	var result []DataElement
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "POST",
		Path:   path,
		Query:  query,
		Body:   body,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
