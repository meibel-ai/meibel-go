package meibelgo

import (
	"net/http"
	"time"
)

// MeibelgoClient is the main client for the meibel API.
type MeibelgoClient struct {
	http *HTTPClient

	Blueprints *BlueprintsService
	BlueprintsExecutions *BlueprintsExecutionsService
	BlueprintsInstances *BlueprintsInstancesService
	ConfidenceScoring *ConfidenceScoringService
	ConfidenceScoring *ConfidenceScoringService
	Content *ContentService
	DataElementMetadata *DataElementMetadataService
	DataElements *DataElementsService
	Datasources *DatasourcesService
	Datasources *DatasourcesService
	DatasourcesContent *DatasourcesContentService
	DatasourcesDataelements *DatasourcesDataelementsService
	DatasourcesMetadataModelCatalog *DatasourcesMetadataModelCatalogService
	DatasourcesRag *DatasourcesRagService
	DatasourcesTag *DatasourcesTagService
	Documents *DocumentsService
	MetadataConfiguration *MetadataConfigurationService
	MetadataModelCatalog *MetadataModelCatalogService
	TagDescriptions *TagDescriptionsService
}

// ClientOption is a function that configures the client.
type ClientOption func(*clientOptions)

type clientOptions struct {
	baseURL    string
	timeout    time.Duration
	httpClient *http.Client
	headers    map[string]string
}

func defaultClientOptions() *clientOptions {
	return &clientOptions{
		baseURL: "https://api.meibel.ai/v1",
		timeout: 30 * time.Second,
		headers: make(map[string]string),
	}
}

// WithBaseURL sets the base URL for API requests.
func WithBaseURL(url string) ClientOption {
	return func(o *clientOptions) {
		o.baseURL = url
	}
}

// WithTimeout sets the request timeout.
func WithTimeout(timeout time.Duration) ClientOption {
	return func(o *clientOptions) {
		o.timeout = timeout
	}
}

// WithHTTPClient sets a custom HTTP client.
func WithHTTPClient(client *http.Client) ClientOption {
	return func(o *clientOptions) {
		o.httpClient = client
	}
}

// WithAPIKey sets the API key for authentication.
func WithAPIKey(key string) ClientOption {
	return func(o *clientOptions) {
		o.headers["Meibel-API-Key"] = key
	}
}

// WithBearerToken sets the bearer token for authentication.
func WithBearerToken(token string) ClientOption {
	return func(o *clientOptions) {
		o.headers["Authorization"] = "Bearer " + token
	}
}

// WithHeader sets a custom header for all requests.
func WithHeader(key, value string) ClientOption {
	return func(o *clientOptions) {
		o.headers[key] = value
	}
}

// NewClient creates a new MeibelgoClient with the given options.
func NewClient(opts ...ClientOption) *MeibelgoClient {
	cfg := defaultClientOptions()
	for _, opt := range opts {
		opt(cfg)
	}

	httpClient := NewHTTPClient(HTTPClientConfig{
		BaseURL:    cfg.baseURL,
		Timeout:    cfg.timeout,
		Headers:    cfg.headers,
		HTTPClient: cfg.httpClient,
	})

	c := &MeibelgoClient{
		http: httpClient,
	}

	c.Blueprints = &BlueprintsService{client: c}
	c.BlueprintsExecutions = &BlueprintsExecutionsService{client: c}
	c.BlueprintsInstances = &BlueprintsInstancesService{client: c}
	c.ConfidenceScoring = &ConfidenceScoringService{client: c}
	c.ConfidenceScoring = &ConfidenceScoringService{client: c}
	c.Content = &ContentService{client: c}
	c.DataElementMetadata = &DataElementMetadataService{client: c}
	c.DataElements = &DataElementsService{client: c}
	c.Datasources = &DatasourcesService{client: c}
	c.Datasources = &DatasourcesService{client: c}
	c.DatasourcesContent = &DatasourcesContentService{client: c}
	c.DatasourcesDataelements = &DatasourcesDataelementsService{client: c}
	c.DatasourcesMetadataModelCatalog = &DatasourcesMetadataModelCatalogService{client: c}
	c.DatasourcesRag = &DatasourcesRagService{client: c}
	c.DatasourcesTag = &DatasourcesTagService{client: c}
	c.Documents = &DocumentsService{client: c}
	c.MetadataConfiguration = &MetadataConfigurationService{client: c}
	c.MetadataModelCatalog = &MetadataModelCatalogService{client: c}
	c.TagDescriptions = &TagDescriptionsService{client: c}

	return c
}
