package meibelgo

import (
	"context"
	"fmt"
	"net/url"
)

// AgentsService handles Agents operations.
type AgentsService struct {
	client *MeibelClient
}

// ListAgentsOptions contains optional parameters for ListAgents.
type ListAgentsOptions struct {
	// Number of items to skip
	Offset *int64
	// Maximum number of items to return
	Limit interface{}
}

// PublishAgentOptions contains optional parameters for PublishAgent.
type PublishAgentOptions struct {
	// Bypass draft head validation and publish any version directly
	OverrideDraft *bool
}

// ListAgentVersionsOptions contains optional parameters for ListAgentVersions.
type ListAgentVersionsOptions struct {
	// If true, return only published versions. If omitted, return all versions.
	Published interface{}
	// Number of items to skip
	Offset *int64
	// Maximum number of items to return
	Limit interface{}
}

// ListSessionsOptions contains optional parameters for ListSessions.
type ListSessionsOptions struct {
	// Number of items to skip
	Offset *int64
	// Maximum number of items to return
	Limit interface{}
	// Field to sort by: start_time, status
	SortBy *string
	// Sort order: asc or desc
	SortOrder *string
	// Filter by execution status: RUNNING, COMPLETED, FAILED, CANCELED, TERMINATED
	Status interface{}
}

// ListAgents List Agents
func (s *AgentsService) ListAgents(ctx context.Context, opts *ListAgentsOptions) *PageIterator[AgentSummary] {
	path := "/agents/"
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

// CreateAgent Create Agent
func (s *AgentsService) CreateAgent(ctx context.Context, body CreateAgentDefinitionRequest) (*CreateAgentResponse, error) {
	path := "/agents/"

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

// GetAgent Get Agent
func (s *AgentsService) GetAgent(ctx context.Context, agentId string) (*AgentDetailResponse, error) {
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

// UpdateAgent Update Agent
func (s *AgentsService) UpdateAgent(ctx context.Context, agentId string, body UpdateAgentDefinitionRequest) (*UpdateAgentDefinitionResponse, error) {
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

// DeleteAgent Delete Agent
func (s *AgentsService) DeleteAgent(ctx context.Context, agentId string) error {
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

// PublishAgent Publish Agent
func (s *AgentsService) PublishAgent(ctx context.Context, agentId string, body PublishAgentDefinitionRequest, opts *PublishAgentOptions) (*PublishAgentDefinitionResponse, error) {
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

// ListAgentVersions List Agent Versions
func (s *AgentsService) ListAgentVersions(ctx context.Context, agentId string, opts *ListAgentVersionsOptions) *PageIterator[AgentVersionSummary] {
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

// ListSessions List Sessions
func (s *AgentsService) ListSessions(ctx context.Context, agentId string, opts *ListSessionsOptions) *PageIterator[SessionSummary] {
	path := "/agents/" + fmt.Sprintf("%v", agentId) + "/sessions"
	query := url.Values{}
	if opts != nil && opts.Offset != nil {
		query.Set("offset", fmt.Sprintf("%v", *opts.Offset))
	}
	if opts != nil && opts.Limit != nil {
		query.Set("limit", fmt.Sprintf("%v", opts.Limit))
	}
	if opts != nil && opts.SortBy != nil {
		query.Set("sort_by", fmt.Sprintf("%v", *opts.SortBy))
	}
	if opts != nil && opts.SortOrder != nil {
		query.Set("sort_order", fmt.Sprintf("%v", *opts.SortOrder))
	}
	if opts != nil && opts.Status != nil {
		query.Set("status", fmt.Sprintf("%v", opts.Status))
	}

	return NewPageIterator(func(ctx context.Context, cursor string) (*Page[SessionSummary], error) {
		if cursor != "" {
			query.Set("offset", cursor)
		}

		var resp struct {
			Data []SessionSummary `json:"data"`
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

		return &Page[SessionSummary]{
			Items:      resp.Data,
			NextCursor: resp.NextCursor,
		}, nil
	})
}

// CreateSession Create Session
func (s *AgentsService) CreateSession(ctx context.Context, agentId string, body *interface{}) (*CreateSessionResponse, error) {
	path := "/agents/" + fmt.Sprintf("%v", agentId) + "/sessions"

	var result CreateSessionResponse
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
