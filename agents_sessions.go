package meibelgo

import (
	"context"
	"fmt"
	"net/url"
)

// AgentsSessionsService handles AgentsSessions operations.
type AgentsSessionsService struct {
	client *MeibelClient
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

// ListSessions List Sessions
func (s *AgentsSessionsService) ListSessions(ctx context.Context, agentId string, opts *ListSessionsOptions) *PageIterator[SessionSummary] {
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
func (s *AgentsSessionsService) CreateSession(ctx context.Context, agentId string, body *interface{}) (*CreateSessionResponse, error) {
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

// SendChatMessage Send Chat Message
func (s *AgentsSessionsService) SendChatMessage(ctx context.Context, sessionId string, body ChatMessageRequest) (*ChatMessageResponse, error) {
	path := "/sessions/" + fmt.Sprintf("%v", sessionId) + "/chat"

	var result ChatMessageResponse
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

// SendChatMessageStream Send a chat message and stream the response via SSE
func (s *AgentsSessionsService) SendChatMessageStream(ctx context.Context, sessionId string, body ChatMessageRequest) (*EventStream[interface{}], error) {
	path := "/sessions/" + fmt.Sprintf("%v", sessionId) + "/chat/stream"

	resp, err := s.client.http.DoStream(ctx, RequestOptions{
		Method: "POST",
		Path:   path,
		Body:   body,
	})
	if err != nil {
		return nil, err
	}

	return JSONEventStream[interface{}](resp), nil
}
