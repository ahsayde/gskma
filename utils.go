package gskma

import (
	"reflect"
	"strings"
)

func getField(f reflect.StructField) *schema {
	s := newSchema(f.Type)
	tag := f.Tag.Get("json")

	if tag == "" {
		s.Name = f.Name
		return s
	}

	segments := strings.Split(tag, ",")
	if segments[0] == "-" {
		return nil
	}

	s.Name = segments[0]

	for _, segment := range segments[1:] {
		parts := strings.Split(segment, "=")
		switch parts[0] {
		case "max", "maximum":
			v, err := converToFloat64(parts[1])
			if err == nil {
				s.Maximum = &v
			}
		case "min", "minimum":
			v, err := converToFloat64(parts[1])
			if err == nil {
				s.Minimum = &v
			}
		case "exclmax", "exclusiveMaximum":
			v, err := converToFloat64(parts[1])
			if err == nil {
				s.ExclusiveMaximum = &v
			}
		case "exclmin", "exclusiveMinimum":
			v, err := converToFloat64(parts[1])
			if err == nil {
				s.ExclusiveMinimum = &v
			}
		case "maxlen", "maxLength":
			v, err := converToInt64(parts[1])
			a := int(v)
			if err == nil {
				s.MaxLength = &a
			}
		case "minlen", "minLength":
			v, err := converToInt64(parts[1])
			a := int(v)
			if err == nil {
				s.MinLength = &a
			}
		case "multof", "multipleOf":
			v, err := converToInt64(parts[1])
			if err == nil {
				s.MultipleOf = &v
			}
		case "maxitems", "maxItems":
			v, err := converToInt64(parts[1])
			a := int(v)
			if err == nil {
				s.MaxItems = &a
			}
		case "minitems", "minItems":
			v, err := converToInt64(parts[1])
			a := int(v)
			if err == nil {
				s.MinItems = &a
			}
		case "maxprops", "maxProperties":
			v, err := converToInt64(parts[1])
			a := int(v)
			if err == nil {
				s.MaxProperties = &a
			}
		case "minprops", "minProperties":
			v, err := converToInt64(parts[1])
			a := int(v)
			if err == nil {
				s.MinProperties = &a
			}
		}
	}
	return s
}

func newSchema(t reflect.Type) *schema {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	s := schema{rkind: t.Kind()}

	if t.Kind() == reflect.Struct {
		s.ID = t.String()
		s.Name = t.Name()
		s.Type = "object"
		s.Properties = make(map[string]*schema)
		for i := 0; i < t.NumField(); i++ {
			f := getField(t.Field(i))
			s.Properties[f.Name] = f
			if f.required {
				s.Required = append(s.Required, f.Name)
			}
		}
	} else if t.Kind() == reflect.Map {
		s.Type = "object"
		s.AdditionalProperties = newSchema(t.Elem())

	} else if t.Kind() == reflect.Slice || t.Kind() == reflect.Array {
		s.Type = "array"
		s.Items = newSchema(t.Elem())
	} else {
		s.Type = types[t.Kind()]
		s.Format = formats[t.Kind()]
	}
	return &s
}

func nameOfField(f reflect.StructField) string {
	tag := f.Tag.Get("json")

	if tag == "" {
		return f.Name
	}

	segments := strings.Split(tag, ",")
	if segments[0] == "" {
		return f.Name
	}

	return segments[0]
}

func castIfNumeric(v interface{}, s *schema) (interface{}, error) {
	if v == nil {
		return v, nil
	}

	var err error
	switch s.rkind {
	case reflect.Int32:
		v, err = converToInt32(v)
	case reflect.Int64:
		v, err = converToInt64(v)
	case reflect.Float32:
		v, err = converToFloat32(v)
	case reflect.Float64:
		v, err = converToFloat64(v)
	case reflect.Bool:
		v, err = convertToBool(v)
	}
	return v, err
}
