package meibelgo

import (
	"context"
	"fmt"
	"net/url"
)

// TagService handles tag operations.
type TagService struct {
	client *MeibelgoClient
}

// GetAllTagTableInfoOptions contains optional parameters for GetAllTagTableInfo.
type GetAllTagTableInfoOptions struct {
	// Number of items to skip
	Offset *int64
	// Maximum number of items to return
	Limit *int64
	// Field to sort by
	SortBy interface{}
	// Sort order (asc or desc)
	SortOrder interface{}
}

// GetAllTagColumnInfoOptions contains optional parameters for GetAllTagColumnInfo.
type GetAllTagColumnInfoOptions struct {
	// Number of items to skip
	Offset *int64
	// Maximum number of items to return
	Limit *int64
	// Field to sort by
	SortBy interface{}
	// Sort order (asc or desc)
	SortOrder interface{}
}

// GetTagConfig Get Tag Config
func (s *TagService) GetTagConfig(ctx context.Context, datasourceId string) (*TagConfig, error) {
	path := "/datasource/" + fmt.Sprintf("%v", datasourceId) + "/tag_config"

	var result TagConfig
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "GET",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// AddTagConfig Add Tag Config
func (s *TagService) AddTagConfig(ctx context.Context, datasourceId string, body AddTagConfigRequest) (*AddTagConfigResponse, error) {
	path := "/datasource/" + fmt.Sprintf("%v", datasourceId) + "/tag_config"

	var result AddTagConfigResponse
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

// UpdateTagConfig Update Tag Config
func (s *TagService) UpdateTagConfig(ctx context.Context, datasourceId string, body UpdateTagConfigRequest) (*UpdateTagConfigResponse, error) {
	path := "/datasource/" + fmt.Sprintf("%v", datasourceId) + "/tag_config"

	var result UpdateTagConfigResponse
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

// DeleteTagConfig Delete Tag Config
func (s *TagService) DeleteTagConfig(ctx context.Context, datasourceId string) (*string, error) {
	path := "/datasource/" + fmt.Sprintf("%v", datasourceId) + "/tag_config"

	var result string
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "DELETE",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// GetTagTableInfo Get Tag Table Info
func (s *TagService) GetTagTableInfo(ctx context.Context, datasourceId string, tableName string) (*TagTableInfo, error) {
	path := "/datasource/" + fmt.Sprintf("%v", datasourceId) + "/tag_table_info/" + fmt.Sprintf("%v", tableName)

	var result TagTableInfo
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "GET",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// AddTagTableInfo Add Tag Table Info
func (s *TagService) AddTagTableInfo(ctx context.Context, datasourceId string, tableName string, body AddTagTableRequest) (*AddTagTableResponse, error) {
	path := "/datasource/" + fmt.Sprintf("%v", datasourceId) + "/tag_table_info/" + fmt.Sprintf("%v", tableName)

	var result AddTagTableResponse
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

// UpdateTagTableInfo Update Tag Table Info
func (s *TagService) UpdateTagTableInfo(ctx context.Context, datasourceId string, tableName string, body UpdateTagTableRequest) (*UpdateTagTableResponse, error) {
	path := "/datasource/" + fmt.Sprintf("%v", datasourceId) + "/tag_table_info/" + fmt.Sprintf("%v", tableName)

	var result UpdateTagTableResponse
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

// DeleteTagTableInfo Delete Tag Table Info
func (s *TagService) DeleteTagTableInfo(ctx context.Context, datasourceId string, tableName string) (*DeleteTagTableResponse, error) {
	path := "/datasource/" + fmt.Sprintf("%v", datasourceId) + "/tag_table_info/" + fmt.Sprintf("%v", tableName)

	var result DeleteTagTableResponse
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "DELETE",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// GetAllTagTableInfo Get All Tag Table Info
func (s *TagService) GetAllTagTableInfo(ctx context.Context, datasourceId string, opts *GetAllTagTableInfoOptions) (*[]TagTableInfo, error) {
	path := "/datasource/" + fmt.Sprintf("%v", datasourceId) + "/tag_table_info"
	query := url.Values{}
	if opts != nil && opts.Offset != nil {
		query.Set("offset", fmt.Sprintf("%v", *opts.Offset))
	}
	if opts != nil && opts.Limit != nil {
		query.Set("limit", fmt.Sprintf("%v", *opts.Limit))
	}
	if opts != nil && opts.SortBy != nil {
		query.Set("sort_by", fmt.Sprintf("%v", opts.SortBy))
	}
	if opts != nil && opts.SortOrder != nil {
		query.Set("sort_order", fmt.Sprintf("%v", opts.SortOrder))
	}

	var result []TagTableInfo
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

// GetTagColumnInfo Get Tag Column Info
func (s *TagService) GetTagColumnInfo(ctx context.Context, datasourceId string, tableName string, columnName string) (*TagColumnInfo, error) {
	path := "/datasource/" + fmt.Sprintf("%v", datasourceId) + "/tag_table_info/" + fmt.Sprintf("%v", tableName) + "/column_info/" + fmt.Sprintf("%v", columnName)

	var result TagColumnInfo
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "GET",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// AddTagColumnInfo Add Tag Column Info
func (s *TagService) AddTagColumnInfo(ctx context.Context, datasourceId string, tableName string, columnName string, body AddTagColumnRequest) (*AddTagColumnResponse, error) {
	path := "/datasource/" + fmt.Sprintf("%v", datasourceId) + "/tag_table_info/" + fmt.Sprintf("%v", tableName) + "/column_info/" + fmt.Sprintf("%v", columnName)

	var result AddTagColumnResponse
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

// UpdateTagColumnInfo Update Tag Column Info
func (s *TagService) UpdateTagColumnInfo(ctx context.Context, datasourceId string, tableName string, columnName string, body UpdateTagColumnRequest) (*UpdateTagColumnResponse, error) {
	path := "/datasource/" + fmt.Sprintf("%v", datasourceId) + "/tag_table_info/" + fmt.Sprintf("%v", tableName) + "/column_info/" + fmt.Sprintf("%v", columnName)

	var result UpdateTagColumnResponse
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

// GetAllTagColumnInfo Get All Tag Column Info
func (s *TagService) GetAllTagColumnInfo(ctx context.Context, datasourceId string, tableName string, opts *GetAllTagColumnInfoOptions) (*[]TagColumnInfo, error) {
	path := "/datasource/" + fmt.Sprintf("%v", datasourceId) + "/tag_table_info/" + fmt.Sprintf("%v", tableName) + "/column_info"
	query := url.Values{}
	if opts != nil && opts.Offset != nil {
		query.Set("offset", fmt.Sprintf("%v", *opts.Offset))
	}
	if opts != nil && opts.Limit != nil {
		query.Set("limit", fmt.Sprintf("%v", *opts.Limit))
	}
	if opts != nil && opts.SortBy != nil {
		query.Set("sort_by", fmt.Sprintf("%v", opts.SortBy))
	}
	if opts != nil && opts.SortOrder != nil {
		query.Set("sort_order", fmt.Sprintf("%v", opts.SortOrder))
	}

	var result []TagColumnInfo
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

// DeleteTagColumnInfo Delete Tag Column Info
func (s *TagService) DeleteTagColumnInfo(ctx context.Context, datasourceId string, tableName string, columnName string) (*DeleteTagTableResponse, error) {
	path := "/datasource/" + fmt.Sprintf("%v", datasourceId) + "/tag_config/" + fmt.Sprintf("%v", tableName) + "/column_info/" + fmt.Sprintf("%v", columnName)

	var result DeleteTagTableResponse
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "DELETE",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
