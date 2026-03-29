package meibelgo

import (
	"context"
	"fmt"
	"net/url"
)

// DatasourcesMetadataModelCatalogService handles datasources.metadata_model_catalog operations.
type DatasourcesMetadataModelCatalogService struct {
	client *MeibelgoClient
}

// ListMetadataModelCatalogOptions contains optional parameters for ListMetadataModelCatalog.
type ListMetadataModelCatalogOptions struct {
	Scope interface{}
}

// ListMetadataModelCatalog List Metadata Model Catalog
func (s *DatasourcesMetadataModelCatalogService) ListMetadataModelCatalog(ctx context.Context, opts *ListMetadataModelCatalogOptions) *PageIterator[MetadataModelCatalogEntry] {
	path := "/metadata_model_catalog"
	query := url.Values{}
	if opts != nil && opts.Scope != nil {
		query.Set("scope", fmt.Sprintf("%v", opts.Scope))
	}

	return NewPageIterator(func(ctx context.Context, cursor string) (*Page[MetadataModelCatalogEntry], error) {
		if cursor != "" {
			query.Set("offset", cursor)
		}

		var resp struct {
			Models []MetadataModelCatalogEntry `json:"models"`
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

		return &Page[MetadataModelCatalogEntry]{
			Items:      resp.Models,
			NextCursor: resp.NextCursor,
		}, nil
	})
}

// GetMetadataModelCatalogEntry Get Metadata Model Catalog Entry
func (s *DatasourcesMetadataModelCatalogService) GetMetadataModelCatalogEntry(ctx context.Context, modelId string) (*MetadataModelCatalogEntry, error) {
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
