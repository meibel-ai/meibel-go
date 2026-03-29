package meibelgo

import (
	"context"
	"fmt"
	"net/url"
)

// DataElementsService handles Data Elements operations.
type DataElementsService struct {
	client *MeibelgoClient
}

// ListDataElements List Data Elements
func (s *DataElementsService) ListDataElements(ctx context.Context, datasourceId string) *PageIterator[string] {
	path := "/datasources/" + fmt.Sprintf("%v", datasourceId) + "/data-elements"
	query := url.Values{}

	return NewPageIterator(func(ctx context.Context, cursor string) (*Page[string], error) {
		if cursor != "" {
			query.Set("offset", cursor)
		}

		var resp struct {
			Items      []string `json:"items"`
			NextCursor string   `json:"next_cursor"`
		}

		err := s.client.http.Do(ctx, RequestOptions{
			Method: "GET",
			Path:   path,
			Query:  query,
		}, &resp)
		if err != nil {
			return nil, err
		}

		return &Page[string]{
			Items:      resp.Items,
			NextCursor: resp.NextCursor,
		}, nil
	})
}

// CreateDataElement Create Data Element
func (s *DataElementsService) CreateDataElement(ctx context.Context, datasourceId string, body CreateDataElementRequest) (*DataElementResponse, error) {
	path := "/datasources/" + fmt.Sprintf("%v", datasourceId) + "/data-elements"

	var result DataElementResponse
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
func (s *DataElementsService) GetDataElement(ctx context.Context, datasourceId string, dataElementId string) (*DataElementResponse, error) {
	path := "/datasources/" + fmt.Sprintf("%v", datasourceId) + "/data-elements/" + fmt.Sprintf("%v", dataElementId)

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
func (s *DataElementsService) UpdateDataElement(ctx context.Context, datasourceId string, dataElementId string, body UpdateDataElementRequest) (*DataElementResponse, error) {
	path := "/datasources/" + fmt.Sprintf("%v", datasourceId) + "/data-elements/" + fmt.Sprintf("%v", dataElementId)

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

// DeleteDataElement Delete Data Element
func (s *DataElementsService) DeleteDataElement(ctx context.Context, datasourceId string, dataElementId string) (*string, error) {
	path := "/datasources/" + fmt.Sprintf("%v", datasourceId) + "/data-elements/" + fmt.Sprintf("%v", dataElementId)

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

// SearchDataElements Search Data Elements
func (s *DataElementsService) SearchDataElements(ctx context.Context, datasourceId string, body DataElementSearchRequest) (*[]DataElementResponse, error) {
	path := "/datasources/" + fmt.Sprintf("%v", datasourceId) + "/data-elements/search"

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
