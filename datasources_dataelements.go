package meibelgo

import (
	"context"
	"fmt"
	"net/url"
)

// DatasourcesDataelementsService handles datasources.dataelements operations.
type DatasourcesDataelementsService struct {
	client *MeibelgoClient
}

// GetDataElementsByFiltersOptions contains optional parameters for GetDataElementsByFilters.
type GetDataElementsByFiltersOptions struct {
	RegexFilter *string
	MediaTypeFilters []string
}

// GetDataElements Get Data Elements
func (s *DatasourcesDataelementsService) GetDataElements(ctx context.Context, datasourceId string) (*[]DataElement, error) {
	path := "/datasource/" + fmt.Sprintf("%v", datasourceId) + "/data_element"

	var result []DataElement
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "GET",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// AddDataElement Add Data Element
func (s *DatasourcesDataelementsService) AddDataElement(ctx context.Context, datasourceId string, body AddDataElementRequest) (*AddDataElementResponse, error) {
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
func (s *DatasourcesDataelementsService) GetDataElement(ctx context.Context, datasourceId string, dataElementId string) (*DataElement, error) {
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
func (s *DatasourcesDataelementsService) UpdateDataElement(ctx context.Context, datasourceId string, dataElementId string, body DatasourceServiceClientModelsUpdateDataElementRequestUpdateDataElementRequest) (*UpdateDataElementResponse, error) {
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
func (s *DatasourcesDataelementsService) DeleteDataElement(ctx context.Context, datasourceId string, dataElementId string) (*DeleteDataElementResponse, error) {
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
func (s *DatasourcesDataelementsService) GetDataElementsByFilters(ctx context.Context, datasourceId string, body *DataElementFilterRequest, opts *GetDataElementsByFiltersOptions) (*[]DataElement, error) {
	path := "/datasource/" + fmt.Sprintf("%v", datasourceId) + "/data_elements_by_filters"
	query := url.Values{}
	if opts != nil && opts.RegexFilter != nil {
		query.Set("regex_filter", fmt.Sprintf("%v", *opts.RegexFilter))
	}
	if opts != nil && opts.MediaTypeFilters != nil {
		query.Set("media_type_filters", fmt.Sprintf("%v", opts.MediaTypeFilters))
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

// GetDataElementMetadata Get Data Element Metadata
func (s *DatasourcesDataelementsService) GetDataElementMetadata(ctx context.Context, datasourceId string, dataElementId string) (*GetDataElementMetadataResponse, error) {
	path := "/datasource/" + fmt.Sprintf("%v", datasourceId) + "/data_element/" + fmt.Sprintf("%v", dataElementId) + "/metadata"

	var result GetDataElementMetadataResponse
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
func (s *DatasourcesDataelementsService) UpdateDataElementMetadata(ctx context.Context, datasourceId string, dataElementId string, body PutDataElementMetadataRequest) (*PutDataElementMetadataResponse, error) {
	path := "/datasource/" + fmt.Sprintf("%v", datasourceId) + "/data_element/" + fmt.Sprintf("%v", dataElementId) + "/metadata"

	var result PutDataElementMetadataResponse
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
func (s *DatasourcesDataelementsService) GetDataElementMetadataResult(ctx context.Context, datasourceId string, dataElementId string, requestId string) (*GetDataElementMetadataResultResponse, error) {
	path := "/datasource/" + fmt.Sprintf("%v", datasourceId) + "/data_element/" + fmt.Sprintf("%v", dataElementId) + "/metadata/result/" + fmt.Sprintf("%v", requestId)

	var result GetDataElementMetadataResultResponse
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "GET",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
