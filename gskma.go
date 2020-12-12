package gskma

import (
	"encoding/json"
	"reflect"
)

var types = map[reflect.Kind]string{
	reflect.Int:       "integer",
	reflect.Int8:      "integer",
	reflect.Int16:     "integer",
	reflect.Int32:     "integer",
	reflect.Int64:     "integer",
	reflect.Uint:      "integer",
	reflect.Uint8:     "integer",
	reflect.Uint16:    "integer",
	reflect.Uint32:    "integer",
	reflect.Uint64:    "integer",
	reflect.Float32:   "number",
	reflect.Float64:   "number",
	reflect.Bool:      "boolean",
	reflect.String:    "string",
	reflect.Interface: "object",
	reflect.Map:       "object",
	reflect.Struct:    "object",
	reflect.Slice:     "array",
	reflect.Array:     "array",
}

var formats = map[reflect.Kind]string{
	reflect.Int32:   "int32",
	reflect.Int64:   "int64",
	reflect.Float32: "float",
	reflect.Float64: "double",
}

type schema struct {
	ID                   string             `json:"-"`
	Name                 string             `json:"title,omitempty"`
	Type                 string             `json:"type,omitempty"`
	Properties           map[string]*schema `json:"properties,omitempty"`
	AdditionalProperties *schema            `json:"additionalProperties,omitempty"`
	Items                *schema            `json:"items,omitempty"`
	Pattern              string             `json:"pattern,omitempty"`
	Format               string             `json:"format,omitempty"`
	Default              interface{}        `json:"default,omitempty"`
	Maximum              *float64           `json:"maximum,omitempty"`
	ExclusiveMaximum     *float64           `json:"exclusiveMaximum,omitempty"`
	Minimum              *float64           `json:"minimum,omitempty"`
	ExclusiveMinimum     *float64           `json:"exclusiveMinimum,omitempty"`
	MultipleOf           *int64             `json:"multipleOf,omitempty"`
	MaxLength            *int               `json:"maxLength,omitempty"`
	MinLength            *int               `json:"minLength,omitempty"`
	MinItems             *int               `json:"minItems,omitempty"`
	MaxItems             *int               `json:"maxItems,omitempty"`
	MinProperties        *int               `json:"minProperties,omitempty"`
	MaxProperties        *int               `json:"maxProperties,omitempty"`
	Required             []string           `json:"required,omitempty"`
	Enum                 []interface{}      `json:"enum,omitempty"`
	AllOf                []schema           `json:"allOf,omitempty"`
	AnyOf                []schema           `json:"anyOf,omitempty"`
	OneOf                []schema           `json:"oneOf,omitempty"`
	required             bool               `json:"-"`
	rkind                reflect.Kind       `json:"-"`
}

// Schema schema object
type Schema struct {
	data schema
}

// MarshalJSON marchal json
func (s *Schema) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.data)
}

// UnmarshalJSON marchal json
func (s *Schema) UnmarshalJSON(in []byte) error {
	return json.Unmarshal(in, &s.data)
}

// Default set default value
func (s *Schema) Default(value interface{}) *Schema {
	value, err := castIfNumeric(value, &s.data)
	if err != nil {
		panic("invalid default value for type")
	}
	s.data.Default = value
	return s
}

// Maximum set the maximum allowed value
// panics if the type of the schema is not integer (int32, int64) or float (float 32, float64)
func (s *Schema) Maximum(max float64) *Schema {
	if s.data.Type != "integer" && s.data.Type != "number" {
		panic("Maximum must only be used with integer and number types")
	}
	s.data.Maximum = &max
	return s
}

// Minimum set the minimum allowed value
// panics if the type of the schema is not integer (int32, int64) or float (float 32, float64)
func (s *Schema) Minimum(min float64) *Schema {
	if s.data.Type != "integer" && s.data.Type != "number" {
		panic("Minimum can be used only with integer and number types")
	}
	s.data.Minimum = &min
	return s
}

// ExclusiveMaximum set the exclusive maximum allowed value
// panics if the type of the schema is not integer (int32, int64) or float (float 32, float64)
func (s *Schema) ExclusiveMaximum(max float64) *Schema {
	if s.data.Type != "integer" && s.data.Type != "number" {
		panic("ExclusiveMaximum must only be used with integer and number types")
	}
	s.data.ExclusiveMaximum = &max
	return s
}

// ExclusiveMinimum set the exclusive minimum allowed value
// panics if the type of the schema is not integer (int32, int64) or float (float 32, float64)
func (s *Schema) ExclusiveMinimum(min float64) *Schema {
	if s.data.Type != "integer" && s.data.Type != "number" {
		panic("ExclusiveMinimum can be used only with integer and number types")
	}
	s.data.ExclusiveMinimum = &min
	return s
}

// MultipleOf set multipleOf
// panics if the type of the schema is not integer (int32, int64)
func (s *Schema) MultipleOf(value int64) *Schema {
	if s.data.Type != "integer" {
		panic("MultipleOf can be used only with integer type")
	}
	s.data.MultipleOf = &value
	return s
}

// MaxLength set the maximum length of the string
// panics if the type of the schema is not string
func (s *Schema) MaxLength(max int) *Schema {
	if s.data.Type != "string" {
		panic("MaxLength can be used only with string Schema")
	}
	s.data.MaxLength = &max
	return s
}

// MinLength set the minimum length of the string
// panics if the type of the schema is not string
func (s *Schema) MinLength(min int) *Schema {
	if s.data.Type != "string" {
		panic("MinLength can be used only with string Schema")
	}
	s.data.MinLength = &min
	return s
}

// MaxItems set the maximum number of theitems in the array
// panics if the type of the schema is not array or slice
func (s *Schema) MaxItems(max int) *Schema {
	if s.data.Type != "array" {
		panic("MaxItems can be used only with array Schema")
	}
	s.data.MaxItems = &max
	return s
}

// MinItems set the minimum number of the items in the array
// panics if the type of the schema is not array or slice
func (s *Schema) MinItems(min int) *Schema {
	if s.data.Type != "array" {
		panic("MinItems can be used only with array Schema")
	}
	s.data.MinItems = &min
	return s
}

// MaxProperties set the maximum number of the keys in the map
// panics if the type of the schema is not map
func (s *Schema) MaxProperties(max int) *Schema {
	if s.data.Type != "object" {
		panic("MaxProperties can be used only with object Schema")
	}
	s.data.MaxProperties = &max
	return s
}

// MinProperties set the minimum number of the keys in the map
// panics if the type of the schema is not map
func (s *Schema) MinProperties(min int) *Schema {
	if s.data.Type != "object" {
		panic("MinProperties can be used only with object Schema")
	}
	s.data.MinProperties = &min
	return s
}

// Validate validate schema
func (s *Schema) Validate(value interface{}) (interface{}, error) {
	var err error

	value, err = castIfNumeric(value, &s.data)
	if err != nil {
		return nil, err
	}

	v := reflect.ValueOf(value)
	val, err := validate(v, &s.data)
	if err != nil {
		return nil, err
	}

	if val.Kind() == reflect.Invalid {
		return nil, nil
	}

	return val.Interface(), nil
}

// String string Schema
func String() Schema {
	return Schema{
		data: schema{
			Type:  "string",
			rkind: reflect.String,
		},
	}
}

// Int32 int32 Schema
func Int32() Schema {
	return Schema{
		data: schema{
			Type:   "integer",
			Format: "int32",
			rkind:  reflect.Int32,
		},
	}
}

// Int64 int64 Schema
func Int64() Schema {
	return Schema{
		data: schema{
			Type:   "integer",
			Format: "int64",
			rkind:  reflect.Int64,
		},
	}
}

// Float32 float32 Schema
func Float32() Schema {
	return Schema{
		data: schema{
			Type:   "number",
			Format: "float",
			rkind:  reflect.Float32,
		},
	}
}

// Float64 float64 Schema
func Float64() Schema {
	return Schema{
		data: schema{
			Type:   "number",
			Format: "double",
			rkind:  reflect.Float64,
		},
	}
}

// Boolean boolean Schema
func Boolean() Schema {
	return Schema{
		data: schema{
			Type:  "boolean",
			rkind: reflect.Bool,
		},
	}
}

// TypeOf get schema for an interface
func TypeOf(i interface{}) Schema {
	t := reflect.TypeOf(i)
	return Schema{
		data: *newSchema(t),
	}
}
