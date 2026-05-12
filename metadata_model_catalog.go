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

// MetadataModelCatalogListOptions contains optional parameters for List.
type MetadataModelCatalogListOptions struct {
	Scope interface{}
}

// List List Metadata Model Catalog
func (s *MetadataModelCatalogService) List(ctx context.Context, opts *MetadataModelCatalogListOptions) (*ListMetadataModelCatalogResponse, error) {
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

// GetEntry Get Metadata Model Catalog Entry
func (s *MetadataModelCatalogService) GetEntry(ctx context.Context, modelId string) (*MetadataModelCatalogEntry, error) {
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
