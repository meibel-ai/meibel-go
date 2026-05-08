package v2

import (
	"context"
	"fmt"
	"net/url"
)

// ArtifactSchemasService handles ArtifactSchemas operations.
type ArtifactSchemasService struct {
	client *MeibelClient
}

// ListArtifactSchemasOptions contains optional parameters for ListArtifactSchemas.
type ListArtifactSchemasOptions struct {
	// Number of items to skip
	Offset *int64
	// Maximum number of items to return
	Limit interface{}
	// Field to sort by: created_at, name, display_name
	SortBy interface{}
	// Sort order: asc or desc
	SortOrder interface{}
}

// ListArtifactSchemas List Artifact Schemas
func (s *ArtifactSchemasService) ListArtifactSchemas(ctx context.Context, opts *ListArtifactSchemasOptions) *PageIterator[ArtifactSchemaSummary] {
	path := "/artifact-schemas/"
	query := url.Values{}
	if opts != nil && opts.Offset != nil {
		query.Set("offset", fmt.Sprintf("%v", *opts.Offset))
	}
	if opts != nil && opts.Limit != nil {
		query.Set("limit", fmt.Sprintf("%v", opts.Limit))
	}
	if opts != nil && opts.SortBy != nil {
		query.Set("sort_by", fmt.Sprintf("%v", opts.SortBy))
	}
	if opts != nil && opts.SortOrder != nil {
		query.Set("sort_order", fmt.Sprintf("%v", opts.SortOrder))
	}

	return NewPageIterator(func(ctx context.Context, cursor string) (*Page[ArtifactSchemaSummary], error) {
		if cursor != "" {
			query.Set("offset", cursor)
		}

		var resp struct {
			Data []ArtifactSchemaSummary `json:"data"`
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

		return &Page[ArtifactSchemaSummary]{
			Items:      resp.Data,
			NextCursor: resp.NextCursor,
		}, nil
	})
}

// CreateArtifactSchema Create Artifact Schema
func (s *ArtifactSchemasService) CreateArtifactSchema(ctx context.Context, body CreateAgentArtifactRequest) (*CreateArtifactSchemaResponse, error) {
	path := "/artifact-schemas/"

	var result CreateArtifactSchemaResponse
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

// GetArtifactSchema Get Artifact Schema
func (s *ArtifactSchemasService) GetArtifactSchema(ctx context.Context, artifactId string) (*ArtifactSchemaResponse, error) {
	path := "/artifact-schemas/" + fmt.Sprintf("%v", artifactId)

	var result ArtifactSchemaResponse
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "GET",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateArtifactSchema Update Artifact Schema
func (s *ArtifactSchemasService) UpdateArtifactSchema(ctx context.Context, artifactId string, body UpdateAgentArtifactRequest) (*UpdateArtifactSchemaResponse, error) {
	path := "/artifact-schemas/" + fmt.Sprintf("%v", artifactId)

	var result UpdateArtifactSchemaResponse
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

// DeleteArtifactSchema Delete Artifact Schema
func (s *ArtifactSchemasService) DeleteArtifactSchema(ctx context.Context, artifactId string) error {
	path := "/artifact-schemas/" + fmt.Sprintf("%v", artifactId)

	err := s.client.http.Do(ctx, RequestOptions{
		Method: "DELETE",
		Path:   path,
	}, nil)
	if err != nil {
		return err
	}

	return nil
}
