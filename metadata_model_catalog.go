package v2

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

// ListMetadataModelCatalog List Metadata Model Catalog
func (s *MetadataModelCatalogService) ListMetadataModelCatalog(ctx context.Context, opts *ListMetadataModelCatalogOptions) (*ListMetadataModelCatalogResponse, error) {
	path := "/metadata-model-catalog"
	query := url.Values{}
	if opts != nil && opts.Scope != nil {
		query.Set("scope", fmt.Sprintf("%v", opts.Scope))
	}

	var result ListMetadataModelCatalogResponse
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

// GetMetadataModelCatalogEntry Get Metadata Model Catalog Entry
func (s *MetadataModelCatalogService) GetMetadataModelCatalogEntry(ctx context.Context, modelId string) (*MetadataModelCatalogEntry, error) {
	path := "/metadata-model-catalog/" + fmt.Sprintf("%v", modelId)

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
