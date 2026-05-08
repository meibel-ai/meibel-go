package v2

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

// MetadataField represents a single metadata extraction field.
type MetadataField struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Index       bool   `json:"index"`
}

// MetadataConfigRequest represents a metadata configuration for a datasource.
type MetadataConfigRequest struct {
	Type    string          `json:"type"`
	ModelID *string         `json:"model_id,omitempty"`
	Fields  []MetadataField `json:"fields,omitempty"`
}

// MetadataSchemaFromStruct creates a MetadataConfigRequest from a Go struct.
//
// Each exported field in the struct is mapped to a metadata field. The field
// name is derived from the `json` tag (or lowercased struct field name).
// Use the `description` tag to set the field description. Use `index:"false"`
// to disable indexing.
//
// Type mapping:
//
//	string        -> "string"
//	int, int32... -> "integer"
//	float32/64    -> "float"
//	bool          -> "boolean"
//	time.Time     -> "datetime"
//	[]string      -> "list[string]"
//
// Example:
//
//	type InvoiceMetadata struct {
//	    VendorName  string    `json:"vendor_name"  description:"Name of the vendor"`
//	    TotalAmount float64   `json:"total_amount" description:"Invoice total"`
//	    DueDate     time.Time `json:"due_date"     description:"Payment due date"`
//	    LineItems   []string  `json:"line_items"   description:"Line item descriptions" index:"false"`
//	}
//
//	config := MetadataSchemaFromStruct(InvoiceMetadata{})
func MetadataSchemaFromStruct(v interface{}) (*MetadataConfigRequest, error) {
	t := reflect.TypeOf(v)

	// Dereference pointer
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return nil, fmt.Errorf("MetadataSchemaFromStruct: expected struct, got %s", t.Kind())
	}

	fields := make([]MetadataField, 0, t.NumField())

	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)

		// Skip unexported fields
		if !sf.IsExported() {
			continue
		}

		// Determine JSON field name
		name := sf.Name
		if jsonTag := sf.Tag.Get("json"); jsonTag != "" {
			parts := strings.SplitN(jsonTag, ",", 2)
			if parts[0] != "" && parts[0] != "-" {
				name = parts[0]
			}
		} else {
			// Default: lowercase first letter for camelCase
			name = strings.ToLower(name[:1]) + name[1:]
		}

		// Map Go type to Meibel type
		meibelType, err := resolveGoType(sf.Type)
		if err != nil {
			return nil, fmt.Errorf("field '%s': %w", sf.Name, err)
		}

		// Description from struct tag
		description := sf.Tag.Get("description")
		if description == "" {
			description = strings.ReplaceAll(name, "_", " ")
		}

		// Index: default true, disable with `index:"false"`
		index := true
		if indexTag := sf.Tag.Get("index"); strings.EqualFold(indexTag, "false") {
			index = false
		}

		fields = append(fields, MetadataField{
			Name:        name,
			Type:        meibelType,
			Description: description,
			Index:       index,
		})
	}

	return &MetadataConfigRequest{
		Type:   "custom",
		Fields: fields,
	}, nil
}

// CatalogMetadataConfig creates a MetadataConfigRequest that references
// a pre-built model from the metadata model catalog.
func CatalogMetadataConfig(modelID string) *MetadataConfigRequest {
	return &MetadataConfigRequest{
		Type:    "catalog",
		ModelID: &modelID,
	}
}

var timeType = reflect.TypeOf(time.Time{})

func resolveGoType(t reflect.Type) (string, error) {
	// Dereference pointer
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	// Check time.Time first (before Kind switch since it's a struct)
	if t == timeType {
		return "datetime", nil
	}

	switch t.Kind() {
	case reflect.String:
		return "string", nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return "integer", nil
	case reflect.Float32, reflect.Float64:
		return "float", nil
	case reflect.Bool:
		return "boolean", nil
	case reflect.Slice:
		// []string -> "list[string]"
		if t.Elem().Kind() == reflect.String {
			return "list[string]", nil
		}
		return "", fmt.Errorf("unsupported slice element type %s; only []string is supported", t.Elem())
	default:
		return "", fmt.Errorf(
			"unsupported Go type %s; supported types: string, int*, float*, bool, time.Time, []string",
			t,
		)
	}
}