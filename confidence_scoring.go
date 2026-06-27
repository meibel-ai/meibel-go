package v2

import (
	"context"
	"fmt"
	"net/url"
)

// ConfidenceScoringService handles ConfidenceScoring operations.
type ConfidenceScoringService struct {
	client *MeibelClient
}

// ConfidenceScoringListScoringJobsOptions contains optional parameters for ListScoringJobs.
type ConfidenceScoringListScoringJobsOptions struct {
	// Filter by agent name.
	AgentName interface{}
	// Filter by agent version.
	AgentVersion interface{}
	// Filter by agent session ID.
	AgentSessionId interface{}
	// Filter by workflow name.
	AgentWorkflowName interface{}
	// Filter by workflow version.
	AgentWorkflowVersion interface{}
	// Filter by workflow session ID.
	AgentWorkflowSessionId interface{}
	// Filter by tool identifier.
	ToolId interface{}
	// Filter by tool instance identifier.
	ToolInstanceId interface{}
	// Filter by tool execution identifier.
	ToolExecutionId interface{}
}

// GetScoringJob Get a scoring job
//
// Retrieve a single confidence scoring job by its ID, including its current status and score if completed.
func (s *ConfidenceScoringService) GetScoringJob(ctx context.Context, jobId string) (*ScoringJobResponse, error) {
	path := "/confidence-scoring/job/" + fmt.Sprintf("%v", jobId)

	var result ScoringJobResponse
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "GET",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// ListScoringJobs List scoring jobs
//
// List confidence scoring jobs, optionally filtered by identity context fields. All filters are combined with AND logic.
func (s *ConfidenceScoringService) ListScoringJobs(ctx context.Context, opts *ConfidenceScoringListScoringJobsOptions) (*[]ScoringJobResponse, error) {
	path := "/confidence-scoring/jobs"
	query := url.Values{}
	if opts != nil && opts.AgentName != nil {
		query.Set("agent_name", fmt.Sprintf("%v", opts.AgentName))
	}
	if opts != nil && opts.AgentVersion != nil {
		query.Set("agent_version", fmt.Sprintf("%v", opts.AgentVersion))
	}
	if opts != nil && opts.AgentSessionId != nil {
		query.Set("agent_session_id", fmt.Sprintf("%v", opts.AgentSessionId))
	}
	if opts != nil && opts.AgentWorkflowName != nil {
		query.Set("agent_workflow_name", fmt.Sprintf("%v", opts.AgentWorkflowName))
	}
	if opts != nil && opts.AgentWorkflowVersion != nil {
		query.Set("agent_workflow_version", fmt.Sprintf("%v", opts.AgentWorkflowVersion))
	}
	if opts != nil && opts.AgentWorkflowSessionId != nil {
		query.Set("agent_workflow_session_id", fmt.Sprintf("%v", opts.AgentWorkflowSessionId))
	}
	if opts != nil && opts.ToolId != nil {
		query.Set("tool_id", fmt.Sprintf("%v", opts.ToolId))
	}
	if opts != nil && opts.ToolInstanceId != nil {
		query.Set("tool_instance_id", fmt.Sprintf("%v", opts.ToolInstanceId))
	}
	if opts != nil && opts.ToolExecutionId != nil {
		query.Set("tool_execution_id", fmt.Sprintf("%v", opts.ToolExecutionId))
	}

	var result []ScoringJobResponse
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

// GetAgentScoringSummary Get agent scoring summary
//
// Get an aggregated summary of confidence scores for a specific agent.
func (s *ConfidenceScoringService) GetAgentScoringSummary(ctx context.Context, agentName string) (*ScoreSummary, error) {
	path := "/confidence-scoring/summary/agent/" + fmt.Sprintf("%v", agentName)

	var result ScoreSummary
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "GET",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// GetAgentSessionScoringSummary Get agent session scoring summary
//
// Get an aggregated summary of confidence scores for a specific agent session.
func (s *ConfidenceScoringService) GetAgentSessionScoringSummary(ctx context.Context, agentName string, sessionId string) (*ScoreSummary, error) {
	path := "/confidence-scoring/summary/agent/" + fmt.Sprintf("%v", agentName) + "/session/" + fmt.Sprintf("%v", sessionId)

	var result ScoreSummary
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "GET",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
