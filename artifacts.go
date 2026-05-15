package v2

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

// ArtifactOption configures an artifact schema request.
type ArtifactOption func(*artifactConfig)

type artifactConfig struct {
	artifactType    ArtifactType
	description     interface{}
	required        interface{}
	maxSizeBytes    interface{}
	storageStrategy interface{}
}

func defaultArtifactConfig() *artifactConfig {
	return &artifactConfig{
		artifactType: ArtifactTypeJson,
	}
}

// WithArtifactType sets the artifact type (json, csv, markdown, yaml, text, html, pdf).
func WithArtifactType(t ArtifactType) ArtifactOption {
	return func(c *artifactConfig) { c.artifactType = t }
}

// WithArtifactDescription sets the artifact description.
func WithArtifactDescription(d string) ArtifactOption {
	return func(c *artifactConfig) { c.description = d }
}

// WithArtifactRequired sets whether the agent must produce this artifact.
func WithArtifactRequired(r bool) ArtifactOption {
	return func(c *artifactConfig) { c.required = r }
}

// WithArtifactMaxSizeBytes sets the maximum artifact size in bytes.
func WithArtifactMaxSizeBytes(n int64) ArtifactOption {
	return func(c *artifactConfig) { c.maxSizeBytes = n }
}

// WithArtifactStorageStrategy sets the storage strategy (inline, gcs, auto).
func WithArtifactStorageStrategy(s ArtifactStorageStrategy) ArtifactOption {
	return func(c *artifactConfig) { c.storageStrategy = s }
}

// ArtifactSchemaFromStruct converts a Go struct type into a CreateAgentArtifactRequest.
//
// The struct's exported fields are introspected to produce a format-appropriate
// schema definition. Use the `json` tag for field names and `desc` for descriptions.
//
// Example:
//
//	type InvoiceOutput struct {
//	    VendorName  string   `json:"vendor_name"  desc:"Name of the vendor"`
//	    TotalAmount float64  `json:"total_amount" desc:"Invoice total"`
//	    LineItems   []string `json:"line_items"   desc:"Extracted line items"`
//	}
//
//	req, err := ArtifactSchemaFromStruct[InvoiceOutput]("Invoice Output")
func ArtifactSchemaFromStruct[T any](displayName string, opts ...ArtifactOption) (*CreateAgentArtifactRequest, error) {
	cfg := defaultArtifactConfig()
	for _, opt := range opts {
		opt(cfg)
	}

	var zero T
	t := reflect.TypeOf(zero)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return nil, fmt.Errorf("ArtifactSchemaFromStruct: expected struct type parameter, got %s", t.Kind())
	}

	schemaDef, err := buildArtifactSchemaDef(t, string(cfg.artifactType))
	if err != nil {
		return nil, err
	}

	return &CreateAgentArtifactRequest{
		DisplayName:     displayName,
		Type:            cfg.artifactType,
		SchemaDef:       schemaDef,
		Description:     cfg.description,
		Required:        cfg.required,
		MaxSizeBytes:    cfg.maxSizeBytes,
		StorageStrategy: cfg.storageStrategy,
	}, nil
}

// FreeformArtifactSchema creates an artifact request without a struct schema.
// Useful for markdown, text, html, or pdf types where the agent has full discretion.
//
// Example:
//
//	req, err := FreeformArtifactSchema("Report", WithArtifactType(ArtifactTypeMarkdown))
func FreeformArtifactSchema(displayName string, opts ...ArtifactOption) (*CreateAgentArtifactRequest, error) {
	cfg := defaultArtifactConfig()
	for _, opt := range opts {
		opt(cfg)
	}

	at := string(cfg.artifactType)
	if at == "json" || at == "csv" || at == "yaml" {
		return nil, fmt.Errorf(
			"FreeformArtifactSchema: artifact type '%s' requires a struct type; use ArtifactSchemaFromStruct[T]() instead",
			at,
		)
	}

	freeform := map[string]interface{}{
		"freeform": true,
		"sections": []interface{}{},
	}

	return &CreateAgentArtifactRequest{
		DisplayName:     displayName,
		Type:            cfg.artifactType,
		SchemaDef:       freeform,
		Description:     cfg.description,
		Required:        cfg.required,
		MaxSizeBytes:    cfg.maxSizeBytes,
		StorageStrategy: cfg.storageStrategy,
	}, nil
}

// ---------------------------------------------------------------------------
// Internal: schema generation
// ---------------------------------------------------------------------------

var artifactTimeType = reflect.TypeOf(time.Time{})

func buildArtifactSchemaDef(t reflect.Type, artifactType string) (map[string]interface{}, error) {
	switch artifactType {
	case "json", "yaml":
		return structToArtifactJsonSchema(t)

	case "csv":
		columns := structToArtifactCsvColumns(t, "", 0)
		return map[string]interface{}{"columns": columns}, nil

	case "markdown":
		sections := structToArtifactMarkdownSections(t, 2)
		return map[string]interface{}{"sections": sections}, nil

	default:
		// text, html, pdf — use JSON Schema as validation hints
		return structToArtifactJsonSchema(t)
	}
}

// ---------------------------------------------------------------------------
// JSON Schema from struct
// ---------------------------------------------------------------------------

func structToArtifactJsonSchema(t reflect.Type) (map[string]interface{}, error) {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected struct, got %s", t.Kind())
	}

	properties := make(map[string]interface{})
	required := []string{}

	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		if !sf.IsExported() {
			continue
		}

		name := artifactFieldName(sf)
		jsonType := goTypeToArtifactJsonSchema(sf.Type)

		if desc := sf.Tag.Get("desc"); desc != "" {
			jsonType["description"] = desc
		}

		properties[name] = jsonType

		// Non-pointer fields are required
		if sf.Type.Kind() != reflect.Ptr {
			required = append(required, name)
		}
	}

	schema := map[string]interface{}{
		"title":      t.Name(),
		"type":       "object",
		"properties": properties,
	}
	if len(required) > 0 {
		schema["required"] = required
	}
	return schema, nil
}

func goTypeToArtifactJsonSchema(t reflect.Type) map[string]interface{} {
	// Dereference pointer
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	// time.Time
	if t == artifactTimeType {
		return map[string]interface{}{"type": "string", "format": "date-time"}
	}

	switch t.Kind() {
	case reflect.String:
		return map[string]interface{}{"type": "string"}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return map[string]interface{}{"type": "integer"}
	case reflect.Float32, reflect.Float64:
		return map[string]interface{}{"type": "number"}
	case reflect.Bool:
		return map[string]interface{}{"type": "boolean"}
	case reflect.Slice:
		items := goTypeToArtifactJsonSchema(t.Elem())
		return map[string]interface{}{"type": "array", "items": items}
	case reflect.Map:
		return map[string]interface{}{"type": "object"}
	case reflect.Struct:
		nested, err := structToArtifactJsonSchema(t)
		if err != nil {
			return map[string]interface{}{"type": "string"}
		}
		return nested
	default:
		// Best-effort fallback
		return map[string]interface{}{"type": "string"}
	}
}

// ---------------------------------------------------------------------------
// CSV columns from struct
// ---------------------------------------------------------------------------

func structToArtifactCsvColumns(t reflect.Type, prefix string, depth int) []map[string]interface{} {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	columns := []map[string]interface{}{}

	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		if !sf.IsExported() {
			continue
		}

		name := artifactFieldName(sf)
		var fullName string
		if prefix != "" {
			fullName = prefix + "." + name
		} else {
			fullName = name
		}

		desc := sf.Tag.Get("desc")
		if desc == "" {
			desc = strings.ReplaceAll(name, "_", " ")
		}

		ft := sf.Type
		if ft.Kind() == reflect.Ptr {
			ft = ft.Elem()
		}

		// Nested struct (not time.Time)
		if ft.Kind() == reflect.Struct && ft != artifactTimeType && depth < 2 {
			columns = append(columns, structToArtifactCsvColumns(ft, fullName, depth+1)...)
		} else {
			colType := goTypeToCsvType(sf.Type)
			columns = append(columns, map[string]interface{}{
				"name":        fullName,
				"type":        colType,
				"description": desc,
			})
		}
	}

	return columns
}

func goTypeToCsvType(t reflect.Type) string {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	switch t.Kind() {
	case reflect.String:
		return "string"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return "integer"
	case reflect.Float32, reflect.Float64:
		return "number"
	case reflect.Bool:
		return "boolean"
	default:
		return "string"
	}
}

// ---------------------------------------------------------------------------
// Markdown sections from struct
// ---------------------------------------------------------------------------

func structToArtifactMarkdownSections(t reflect.Type, level int) []map[string]interface{} {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	sections := []map[string]interface{}{}

	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		if !sf.IsExported() {
			continue
		}

		name := artifactFieldName(sf)
		desc := sf.Tag.Get("desc")
		if desc == "" {
			desc = strings.ReplaceAll(name, "_", " ")
		}

		headingPrefix := strings.Repeat("#", level)
		headingName := strings.ReplaceAll(name, "_", " ")
		headingName = toTitleCase(headingName)

		sections = append(sections, map[string]interface{}{
			"heading":     headingPrefix + " " + headingName,
			"description": desc,
		})

		ft := sf.Type
		if ft.Kind() == reflect.Ptr {
			ft = ft.Elem()
		}
		if ft.Kind() == reflect.Struct && ft != artifactTimeType && level < 4 {
			sections = append(sections, structToArtifactMarkdownSections(ft, level+1)...)
		}
	}

	return sections
}

// ---------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------

func artifactFieldName(sf reflect.StructField) string {
	if jsonTag := sf.Tag.Get("json"); jsonTag != "" {
		parts := strings.SplitN(jsonTag, ",", 2)
		if parts[0] != "" && parts[0] != "-" {
			return parts[0]
		}
	}
	// Default: lowercase first letter
	name := sf.Name
	return strings.ToLower(name[:1]) + name[1:]
}

func toTitleCase(s string) string {
	words := strings.Fields(s)
	for i, w := range words {
		if len(w) > 0 {
			words[i] = strings.ToUpper(w[:1]) + w[1:]
		}
	}
	return strings.Join(words, " ")
}

// resolveSchema accepts a schema identifier (string), a JSON Schema map, or a Go struct
// and returns the resolved schema value. For structs, it introspects fields via reflection
// to produce a JSON Schema map.
func resolveSchema(v interface{}) (interface{}, error) {
	if v == nil {
		return nil, nil
	}
	switch v.(type) {
	case string:
		return v, nil // Schema ID or pre-serialized JSON
	case map[string]interface{}:
		return v, nil // Already a JSON Schema dict
	default:
		t := reflect.TypeOf(v)
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}
		if t.Kind() == reflect.Struct {
			schema, err := structToArtifactJsonSchema(t)
			if err != nil {
				return nil, fmt.Errorf("resolveSchema: %w", err)
			}
			return schema, nil
		}
		return v, nil // Pass through unknown types
	}
}
