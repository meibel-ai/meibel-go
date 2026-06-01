package v2

import (
	"context"
	"fmt"
	"net/url"
)

// AgentsService handles Agents operations.
type AgentsService struct {
	client *MeibelClient
	Sessions *AgentsSessionsService
}

// AgentsListOptions contains optional parameters for List.
type AgentsListOptions struct {
	// Number of items to skip
	Offset *int64
	// Maximum number of items to return
	Limit interface{}
}

// AgentsPublishOptions contains optional parameters for Publish.
type AgentsPublishOptions struct {
	// Bypass draft head validation and publish any version directly
	OverrideDraft *bool
}

// AgentsListVersionsOptions contains optional parameters for ListVersions.
type AgentsListVersionsOptions struct {
	// If true, return only published versions. If omitted, return all versions.
	Published interface{}
	// Number of items to skip
	Offset *int64
	// Maximum number of items to return
	Limit interface{}
}

// List List Agents
func (s *AgentsService) List(ctx context.Context, opts *AgentsListOptions) *PageIterator[AgentSummary] {
	path := "/agents"
	query := url.Values{}
	if opts != nil && opts.Offset != nil {
		query.Set("offset", fmt.Sprintf("%v", *opts.Offset))
	}
	if opts != nil && opts.Limit != nil {
		query.Set("limit", fmt.Sprintf("%v", opts.Limit))
	}

	return NewPageIterator(func(ctx context.Context, cursor string) (*Page[AgentSummary], error) {
		if cursor != "" {
			query.Set("offset", cursor)
		}

		var resp struct {
			Data []AgentSummary `json:"data"`
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

		return &Page[AgentSummary]{
			Items:      resp.Data,
			NextCursor: resp.NextCursor,
		}, nil
	})
}

// Create Create Agent
func (s *AgentsService) Create(ctx context.Context, body CreateAgentDefinitionRequest) (*CreateAgentResponse, error) {
	path := "/agents"

	var result CreateAgentResponse
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

// Get Get Agent
func (s *AgentsService) Get(ctx context.Context, agentId string) (*AgentDetailResponse, error) {
	path := "/agents/" + fmt.Sprintf("%v", agentId)

	var result AgentDetailResponse
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "GET",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Update Update Agent
func (s *AgentsService) Update(ctx context.Context, agentId string, body UpdateAgentDefinitionRequest) (*UpdateAgentDefinitionResponse, error) {
	path := "/agents/" + fmt.Sprintf("%v", agentId)

	var result UpdateAgentDefinitionResponse
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

// Delete Delete Agent
func (s *AgentsService) Delete(ctx context.Context, agentId string) error {
	path := "/agents/" + fmt.Sprintf("%v", agentId)

	err := s.client.http.Do(ctx, RequestOptions{
		Method: "DELETE",
		Path:   path,
	}, nil)
	if err != nil {
		return err
	}

	return nil
}

// Publish Publish Agent
func (s *AgentsService) Publish(ctx context.Context, agentId string, body PublishAgentDefinitionRequest, opts *AgentsPublishOptions) (*PublishAgentDefinitionResponse, error) {
	path := "/agents/" + fmt.Sprintf("%v", agentId) + "/publish"
	query := url.Values{}
	if opts != nil && opts.OverrideDraft != nil {
		query.Set("override_draft", fmt.Sprintf("%v", *opts.OverrideDraft))
	}

	var result PublishAgentDefinitionResponse
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

// ListVersions List Agent Versions
func (s *AgentsService) ListVersions(ctx context.Context, agentId string, opts *AgentsListVersionsOptions) *PageIterator[AgentVersionSummary] {
	path := "/agents/" + fmt.Sprintf("%v", agentId) + "/versions"
	query := url.Values{}
	if opts != nil && opts.Published != nil {
		query.Set("published", fmt.Sprintf("%v", opts.Published))
	}
	if opts != nil && opts.Offset != nil {
		query.Set("offset", fmt.Sprintf("%v", *opts.Offset))
	}
	if opts != nil && opts.Limit != nil {
		query.Set("limit", fmt.Sprintf("%v", opts.Limit))
	}

	return NewPageIterator(func(ctx context.Context, cursor string) (*Page[AgentVersionSummary], error) {
		if cursor != "" {
			query.Set("offset", cursor)
		}

		var resp struct {
			Data []AgentVersionSummary `json:"data"`
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

		return &Page[AgentVersionSummary]{
			Items:      resp.Data,
			NextCursor: resp.NextCursor,
		}, nil
	})
}
