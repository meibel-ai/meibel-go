package v2

import (
	"context"
	"fmt"
)

// DownloadsService handles Downloads operations.
type DownloadsService struct {
	client *MeibelClient
}

// CreateJob Create Download Job (async)
func (s *DownloadsService) CreateJob(ctx context.Context, datasourceId string, body *interface{}) (*DownloadJobResponse, error) {
	path := "/datasources/" + fmt.Sprintf("%v", datasourceId) + "/downloads"

	var result DownloadJobResponse
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

// StreamProgress Stream Download Progress
func (s *DownloadsService) StreamProgress(ctx context.Context, datasourceId string, jobId string) (*EventStream[interface{}], error) {
	path := "/datasources/" + fmt.Sprintf("%v", datasourceId) + "/downloads/" + fmt.Sprintf("%v", jobId) + "/progress"

	resp, err := s.client.http.DoStream(ctx, RequestOptions{
		Method: "GET",
		Path:   path,
	})
	if err != nil {
		return nil, err
	}

	return JSONEventStream[interface{}](resp), nil
}

// DownloadFile Download File
func (s *DownloadsService) DownloadFile(ctx context.Context, datasourceId string, jobId string) (*string, error) {
	path := "/datasources/" + fmt.Sprintf("%v", datasourceId) + "/downloads/" + fmt.Sprintf("%v", jobId) + "/file"

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

// Process Process Download (sync)
func (s *DownloadsService) Process(ctx context.Context, datasourceId string, body *interface{}) (*string, error) {
	path := "/datasources/" + fmt.Sprintf("%v", datasourceId) + "/downloads/process"

	var result string
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
