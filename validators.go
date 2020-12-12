package gskma

import (
	"fmt"
	"reflect"
)

var invalid = reflect.Value{}

func validate(v reflect.Value, s *schema) (reflect.Value, error) {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	kind := v.Kind()
	if kind == reflect.Invalid && s.Default != nil {
		return reflect.ValueOf(s.Default), nil
	}

	if kind != s.rkind && kind != reflect.Invalid {
		return invalid, fmt.Errorf("invalid type, expected value of type %s", s.Type)
	}

	switch kind {
	case reflect.String:
		return validateString(v, s)
	case reflect.Int32, reflect.Int64, reflect.Float32, reflect.Float64:
		return validateNumber(v, s)
	case reflect.Map:
		return validateMap(v, s)
	case reflect.Struct:
		return validateStruct(v, s)
	case reflect.Array, reflect.Slice:
		return validateArray(v, s)
	default:
		return v, nil
	}
}

func validateString(v reflect.Value, s *schema) (reflect.Value, error) {
	if s.MaxLength != nil && v.Len() > *s.MaxLength {
		return invalid, fmt.Errorf("value length must not exceed %d character(s)", *s.MaxLength)
	}

	if s.MinLength != nil && v.Len() < *s.MinLength {
		return invalid, fmt.Errorf("value length must be at least %d character(s)", *s.MinLength)
	}
	return v, nil
}

func validateNumber(v reflect.Value, s *schema) (reflect.Value, error) {
	var val float64

	switch s.rkind {
	case reflect.Int32, reflect.Int64:
		val = float64(v.Int())
	case reflect.Float32, reflect.Float64:
		val = float64(v.Float())
	}

	if s.Maximum != nil && val > *s.Maximum {
		return invalid, fmt.Errorf("value must be less than or equal %v", *s.Maximum)
	}

	if s.Minimum != nil && val < *s.Minimum {
		return invalid, fmt.Errorf("value must be greater than or equal %v", *s.Minimum)
	}

	if s.ExclusiveMaximum != nil && val >= *s.ExclusiveMaximum {
		return invalid, fmt.Errorf("value must be less than %v", *s.ExclusiveMaximum)
	}

	if s.ExclusiveMinimum != nil && val <= *s.ExclusiveMinimum {
		return invalid, fmt.Errorf("value must be greater than %v", *s.ExclusiveMinimum)
	}

	if s.MultipleOf != nil && v.Int()%(*s.MultipleOf) != 0 {
		return invalid, fmt.Errorf("value must be divisible by %d", *s.MultipleOf)
	}

	return v, nil
}

func validateStruct(v reflect.Value, s *schema) (reflect.Value, error) {
	if s.MaxProperties != nil && v.Len() > *s.MaxProperties {
		return invalid, fmt.Errorf("value must not have more than %d item(s)", *s.MaxProperties)
	}

	if s.MinProperties != nil && v.Len() < *s.MinProperties {
		return invalid, fmt.Errorf("value must have at least %d item(s)", *s.MinProperties)
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		value := v.Field(i)
		_, err := validate(value, s.Properties[nameOfField(field)])
		if err != nil {
			return invalid, err
		}
	}
	return v, nil
}

func validateMap(v reflect.Value, s *schema) (reflect.Value, error) {
	if s.MaxProperties != nil && v.Len() > *s.MaxProperties {
		return invalid, fmt.Errorf("value must not have more than %d item(s)", *s.MaxProperties)
	}

	if s.MinProperties != nil && v.Len() < *s.MinProperties {
		return invalid, fmt.Errorf("value must have at least %d item(s)", *s.MinProperties)
	}

	for _, k := range v.MapKeys() {
		_, err := validate(v.MapIndex(k), s.AdditionalProperties)
		if err != nil {
			return invalid, err
		}
	}

	return v, nil
}

func validateArray(v reflect.Value, s *schema) (reflect.Value, error) {
	if s.MaxItems != nil && v.Len() > *s.MaxItems {
		return invalid, fmt.Errorf("value must not have more than %d item(s)", *s.MaxItems)
	}

	if s.MinItems != nil && v.Len() < *s.MinItems {
		return invalid, fmt.Errorf("value must have at least %d item(s)", *s.MinItems)
	}

	for i := 0; i < v.Len(); i++ {
		_, err := validate(v.Index(i), s.Items)
		if err != nil {
			return invalid, err
		}
	}
	return v, nil
}
