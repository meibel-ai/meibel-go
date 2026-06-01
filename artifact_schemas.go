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

// ArtifactSchemasListOptions contains optional parameters for List.
type ArtifactSchemasListOptions struct {
	// Number of items to skip
	Offset *int64
	// Maximum number of items to return
	Limit interface{}
	// Field to sort by: created_at, name, display_name
	SortBy interface{}
	// Sort order: asc or desc
	SortOrder interface{}
}

// List List Artifact Schemas
func (s *ArtifactSchemasService) List(ctx context.Context, opts *ArtifactSchemasListOptions) *PageIterator[ArtifactSchemaSummary] {
	path := "/artifact-schemas"
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

// ArtifactSchemasCreateOptions contains parameters for Create.
type ArtifactSchemasCreateOptions struct {
	// Human-readable name of the artifact (letters, numbers, and spaces only). Converted to kebab-case internally.
	DisplayName string
	// Artifact type (json, markdown, csv, yaml, text, html, pdf)
	Type *ArtifactType
	// Description of the artifact
	Description interface{}
	// Whether agent must produce this artifact
	Required interface{}
	// Schema definition
	Schema interface{}
	// Maximum artifact size in bytes
	MaxSizeBytes interface{}
	// Storage strategy (inline, gcs, auto)
	StorageStrategy interface{}
	AdditionalProperties map[string]interface{}
}

// Create Create Artifact Schema
func (s *ArtifactSchemasService) Create(ctx context.Context, opts ArtifactSchemasCreateOptions) (*CreateArtifactSchemaResponse, error) {
	path := "/artifact-schemas"
	var err error

	schemaResolved, err := resolveSchema(opts.Schema)
	if err != nil {
		return nil, err
	}

	schemaTyped, _ := schemaResolved.(map[string]interface{})
	body := CreateAgentArtifactRequest{
		DisplayName: opts.DisplayName,
		Description: opts.Description,
		Required: opts.Required,
		SchemaDef: schemaTyped,
		MaxSizeBytes: opts.MaxSizeBytes,
		StorageStrategy: opts.StorageStrategy,
		AdditionalProperties: opts.AdditionalProperties,
	}
	if opts.Type != nil {
		body.Type = *opts.Type
	} else {
		body.Type = ArtifactType("json")
	}

	var result CreateArtifactSchemaResponse
	err = s.client.http.Do(ctx, RequestOptions{
		Method: "POST",
		Path:   path,
		Body:   body,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Get Get Artifact Schema
func (s *ArtifactSchemasService) Get(ctx context.Context, artifactId string) (*ArtifactSchemaResponse, error) {
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

// Update Update Artifact Schema
func (s *ArtifactSchemasService) Update(ctx context.Context, artifactId string, body UpdateAgentArtifactRequest) (*UpdateArtifactSchemaResponse, error) {
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

// Delete Delete Artifact Schema
func (s *ArtifactSchemasService) Delete(ctx context.Context, artifactId string) error {
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
