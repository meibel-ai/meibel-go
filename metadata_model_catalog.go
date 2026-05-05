package meibelgo

import (
	"context"
	"fmt"
	"net/url"
)

// MetadataModelCatalogService handles MetadataModelCatalog operations.
type MetadataModelCatalogService struct {
	client *MeibelClient
}

// ListMetadataModelCatalogOptions contains optional parameters for ListMetadataModelCatalog.
type ListMetadataModelCatalogOptions struct {
	Scope interface{}
}

// ListMetadataModelCatalogOptions contains optional parameters for ListMetadataModelCatalog.
type ListMetadataModelCatalogOptions struct {
	Scope interface{}
}

// ListMetadataModelCatalog List Metadata Model Catalog
func (s *MetadataModelCatalogService) ListMetadataModelCatalog(ctx context.Context, opts *ListMetadataModelCatalogOptions) *PageIterator[ListMetadataModelCatalogResponse] {
	path := "/metadata_model_catalog"
	query := url.Values{}
	if opts != nil && opts.Scope != nil {
		query.Set("scope", fmt.Sprintf("%v", opts.Scope))
	}

	return NewPageIterator(func(ctx context.Context, cursor string) (*Page[ListMetadataModelCatalogResponse], error) {
		if cursor != "" {
			query.Set("page", cursor)
		}

		var resp struct {
			Items []ListMetadataModelCatalogResponse `json:"items"`
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

		return &Page[ListMetadataModelCatalogResponse]{
			Items:      resp.Items,
			NextCursor: resp.NextCursor,
		}, nil
	})
}

// GetMetadataModelCatalogEntry Get Metadata Model Catalog Entry
func (s *MetadataModelCatalogService) GetMetadataModelCatalogEntry(ctx context.Context, modelId string) (*MetadataModelCatalogEntry, error) {
	path := "/metadata_model_catalog/" + fmt.Sprintf("%v", modelId)

	var result MetadataModelCatalogEntry
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "GET",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// ListMetadataModelCatalog List Metadata Model Catalog
func (s *MetadataModelCatalogService) ListMetadataModelCatalog(ctx context.Context, opts *ListMetadataModelCatalogOptions) *PageIterator[ListMetadataModelCatalogResponse] {
	path := "/v2/metadata-model-catalog"
	query := url.Values{}
	if opts != nil && opts.Scope != nil {
		query.Set("scope", fmt.Sprintf("%v", opts.Scope))
	}

	return NewPageIterator(func(ctx context.Context, cursor string) (*Page[ListMetadataModelCatalogResponse], error) {
		if cursor != "" {
			query.Set("page", cursor)
		}

		var resp struct {
			Items []ListMetadataModelCatalogResponse `json:"items"`
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

		return &Page[ListMetadataModelCatalogResponse]{
			Items:      resp.Items,
			NextCursor: resp.NextCursor,
		}, nil
	})
}

// GetMetadataModelCatalogEntry Get Metadata Model Catalog Entry
func (s *MetadataModelCatalogService) GetMetadataModelCatalogEntry(ctx context.Context, modelId string) (*MetadataModelCatalogEntry, error) {
	path := "/v2/metadata-model-catalog/" + fmt.Sprintf("%v", modelId)

	var result MetadataModelCatalogEntry
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "GET",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
