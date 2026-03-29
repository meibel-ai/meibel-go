package meibelgo

import (
	"context"
	"fmt"
)

// DatasourcesTagService handles datasources.tag operations.
type DatasourcesTagService struct {
	client *MeibelgoClient
}

// GetTagConfig Get Tag Config
func (s *DatasourcesTagService) GetTagConfig(ctx context.Context, datasourceId string) (*TagConfig, error) {
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
func (s *DatasourcesTagService) AddTagConfig(ctx context.Context, datasourceId string, body AddTagConfigRequest) (*AddTagConfigResponse, error) {
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
func (s *DatasourcesTagService) UpdateTagConfig(ctx context.Context, datasourceId string, body UpdateTagConfigRequest) (*UpdateTagConfigResponse, error) {
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
func (s *DatasourcesTagService) DeleteTagConfig(ctx context.Context, datasourceId string) (*string, error) {
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
func (s *DatasourcesTagService) GetTagTableInfo(ctx context.Context, datasourceId string, tableName string) (*TagTableInfo, error) {
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
func (s *DatasourcesTagService) AddTagTableInfo(ctx context.Context, datasourceId string, tableName string, body AddTagTableRequest) (*AddTagTableResponse, error) {
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
func (s *DatasourcesTagService) UpdateTagTableInfo(ctx context.Context, datasourceId string, tableName string, body UpdateTagTableRequest) (*UpdateTagTableResponse, error) {
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
func (s *DatasourcesTagService) DeleteTagTableInfo(ctx context.Context, datasourceId string, tableName string) (*DeleteTagTableResponse, error) {
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
func (s *DatasourcesTagService) GetAllTagTableInfo(ctx context.Context, datasourceId string) (*[]TagTableInfo, error) {
	path := "/datasource/" + fmt.Sprintf("%v", datasourceId) + "/tag_table_info"

	var result []TagTableInfo
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "GET",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// GetTagColumnInfo Get Tag Column Info
func (s *DatasourcesTagService) GetTagColumnInfo(ctx context.Context, datasourceId string, tableName string, columnName string) (*TagColumnInfo, error) {
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
func (s *DatasourcesTagService) AddTagColumnInfo(ctx context.Context, datasourceId string, tableName string, columnName string, body AddTagColumnRequest) (*AddTagColumnResponse, error) {
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
func (s *DatasourcesTagService) UpdateTagColumnInfo(ctx context.Context, datasourceId string, tableName string, columnName string, body UpdateTagColumnRequest) (*UpdateTagColumnResponse, error) {
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

// DeleteTagColumnInfo Delete Tag Column Info
func (s *DatasourcesTagService) DeleteTagColumnInfo(ctx context.Context, datasourceId string, tableName string, columnName string) (*DeleteTagTableResponse, error) {
	path := "/datasource/" + fmt.Sprintf("%v", datasourceId) + "/tag_table_info/" + fmt.Sprintf("%v", tableName) + "/column_info/" + fmt.Sprintf("%v", columnName)

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

// GetAllTagColumnInfo Get All Tag Column Info
func (s *DatasourcesTagService) GetAllTagColumnInfo(ctx context.Context, datasourceId string, tableName string) (*[]TagColumnInfo, error) {
	path := "/datasource/" + fmt.Sprintf("%v", datasourceId) + "/tag_table_info/" + fmt.Sprintf("%v", tableName) + "/column_info"

	var result []TagColumnInfo
	err := s.client.http.Do(ctx, RequestOptions{
		Method: "GET",
		Path:   path,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
