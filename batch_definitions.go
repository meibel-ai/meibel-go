package v2

import (
	"context"
	"fmt"
	"net/url"
)

// BatchDefinitionsService handles BatchDefinitions operations.
type BatchDefinitionsService struct {
	client *MeibelClient
}

// BatchDefinitionsListOptions contains optional parameters for List.
type BatchDefinitionsListOptions struct {
	Offset *int64
	Limit *int64
}

// BatchDefinitionsListVersionsOptions contains optional parameters for ListVersions.
type BatchDefinitionsListVersionsOptions struct {
	Offset *int64
	Limit interface{}
}

// List List Batch Definitions
func (s *BatchDefinitionsService) List(ctx context.Context, opts *BatchDefinitionsListOptions) *PageIterator[BatchDefinitionResponse] {
	path := "/batch-definitions/"
	query := url.Values{}
	if opts != nil && opts.Offset != nil {
		query.Set("offset", fmt.Sprintf("%v", *opts.Offset))
	}
	if opts != nil && opts.Limit != nil {
		query.Set("limit", fmt.Sprintf("%v", *opts.Limit))
	}

	return NewPageIterator(func(ctx context.Context, cursor string) (*Page[BatchDefinitionResponse], error) {
		if cursor != "" {
			query.Set("offset", cursor)
		}

		var resp struct {
			Data []BatchDefinitionResponse `json:"data"`
			NextCursor string `json:"next_cursor"`
		}

		err := s.client.http.Do(ctx, RequestOptions{
			Method: "GET",
			Path:   path,
			Query:  query,
		}, &resp)
		if err != nil {
			return nil, err
		}

		return &Page[BatchDefinitionResponse]{
			Items:      resp.Data,
			NextCursor: resp.NextCursor,
		}, nil
	})
}

// Create Create Batch Definition
func (s *BatchDefinitionsService) Create(ctx context.Context, body CreateBatchDefinitionRequest) (*CreateBatchDefinitionResponse, error) {
	path := "/batch-definitions/"

	var result CreateBatchDefinitionResponse
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

// GetByCatalogUrn Get Batch Definition By Catalog Urn
func (s *BatchDefinitionsService) GetByCatalogUrn(ctx context.Context, catalogUrn string) (*BatchDefinitionResponse, error) {
	path := "/batch-definitions/catalog-urn"
	query := url.Values{}
	query.Set("catalog_urn", fmt.Sprintf("%v", catalogUrn))

	var result BatchDefinitionResponse
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

// GetById Get Batch Definition By Id
func (s *BatchDefinitionsService) GetById(ctx context.Context, definitionId string) (*BatchDefinitionResponse, error) {
	path := "/batch-definitions/id/" + fmt.Sprintf("%v", definitionId)

	var result BatchDefinitionResponse
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "GET",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateById Update Batch Definition By Id
func (s *BatchDefinitionsService) UpdateById(ctx context.Context, definitionId string, body UpdateBatchDefinitionRequest) (*UpdateBatchDefinitionResponse, error) {
	path := "/batch-definitions/id/" + fmt.Sprintf("%v", definitionId)

	var result UpdateBatchDefinitionResponse
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

// DeleteById Delete Batch Definition By Id
func (s *BatchDefinitionsService) DeleteById(ctx context.Context, definitionId string) error {
	path := "/batch-definitions/id/" + fmt.Sprintf("%v", definitionId)

	err := s.client.http.Do(ctx, RequestOptions{
		Method: "DELETE",
		Path:   path,
	}, nil)
	if err != nil {
		return err
	}

	return nil
}

// ListVersions List Batch Definition Versions
func (s *BatchDefinitionsService) ListVersions(ctx context.Context, definitionId string, opts *BatchDefinitionsListVersionsOptions) *PageIterator[BatchDefinitionResponse] {
	path := "/batch-definitions/id/" + fmt.Sprintf("%v", definitionId) + "/versions"
	query := url.Values{}
	if opts != nil && opts.Offset != nil {
		query.Set("offset", fmt.Sprintf("%v", *opts.Offset))
	}
	if opts != nil && opts.Limit != nil {
		query.Set("limit", fmt.Sprintf("%v", opts.Limit))
	}

	return NewPageIterator(func(ctx context.Context, cursor string) (*Page[BatchDefinitionResponse], error) {
		if cursor != "" {
			query.Set("offset", cursor)
		}

		var resp struct {
			Data []BatchDefinitionResponse `json:"data"`
			NextCursor string `json:"next_cursor"`
		}

		err := s.client.http.Do(ctx, RequestOptions{
			Method: "GET",
			Path:   path,
			Query:  query,
		}, &resp)
		if err != nil {
			return nil, err
		}

		return &Page[BatchDefinitionResponse]{
			Items:      resp.Data,
			NextCursor: resp.NextCursor,
		}, nil
	})
}

// Execute Execute Batch Definition
func (s *BatchDefinitionsService) Execute(ctx context.Context, definitionId string) (*ExecuteBatchDefinitionResponse, error) {
	path := "/batch-definitions/id/" + fmt.Sprintf("%v", definitionId) + "/execute"

	var result ExecuteBatchDefinitionResponse
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "POST",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
