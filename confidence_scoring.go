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

// ListScoringJobsOptions contains optional parameters for ListScoringJobs.
type ListScoringJobsOptions struct {
	AgentName interface{}
	AgentVersion interface{}
	AgentExecutionId interface{}
	AgentWorkflowName interface{}
	AgentWorkflowVersion interface{}
	AgentWorkflowExecutionId interface{}
	ToolId interface{}
	ToolInstanceId interface{}
	ToolExecutionId interface{}
}

// GetScoringJobsSummaryOptions contains optional parameters for GetScoringJobsSummary.
type GetScoringJobsSummaryOptions struct {
	Secondary interface{}
}

// GetScoringJob Get Scoring Job
func (s *ConfidenceScoringService) GetScoringJob(ctx context.Context, jobId string) (*string, error) {
	path := "/confidence-scoring/job/" + fmt.Sprintf("%v", jobId)

	var result string
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "GET",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// ListScoringJobs List Scoring Jobs
func (s *ConfidenceScoringService) ListScoringJobs(ctx context.Context, opts *ListScoringJobsOptions) (*string, error) {
	path := "/confidence-scoring/jobs"
	query := url.Values{}
	if opts != nil && opts.AgentName != nil {
		query.Set("agent_name", fmt.Sprintf("%v", opts.AgentName))
	}
	if opts != nil && opts.AgentVersion != nil {
		query.Set("agent_version", fmt.Sprintf("%v", opts.AgentVersion))
	}
	if opts != nil && opts.AgentExecutionId != nil {
		query.Set("agent_execution_id", fmt.Sprintf("%v", opts.AgentExecutionId))
	}
	if opts != nil && opts.AgentWorkflowName != nil {
		query.Set("agent_workflow_name", fmt.Sprintf("%v", opts.AgentWorkflowName))
	}
	if opts != nil && opts.AgentWorkflowVersion != nil {
		query.Set("agent_workflow_version", fmt.Sprintf("%v", opts.AgentWorkflowVersion))
	}
	if opts != nil && opts.AgentWorkflowExecutionId != nil {
		query.Set("agent_workflow_execution_id", fmt.Sprintf("%v", opts.AgentWorkflowExecutionId))
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

	var result string
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

// GetScoringJobsSummary Get Scoring Jobs Summary
func (s *ConfidenceScoringService) GetScoringJobsSummary(ctx context.Context, primary string, opts *GetScoringJobsSummaryOptions) (*ScoreSummary, error) {
	path := "/confidence-scoring/summary"
	query := url.Values{}
	query.Set("primary", fmt.Sprintf("%v", primary))
	if opts != nil && opts.Secondary != nil {
		query.Set("secondary", fmt.Sprintf("%v", opts.Secondary))
	}

	var result ScoreSummary
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
