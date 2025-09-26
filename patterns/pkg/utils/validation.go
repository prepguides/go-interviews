package utils

import (
	"fmt"
	"reflect"
	"strings"
)

// ValidationError represents a validation error
type ValidationError struct {
	Field   string
	Message string
	Value   interface{}
}

func (ve *ValidationError) Error() string {
	return fmt.Sprintf("validation error for field '%s': %s (value: %v)", ve.Field, ve.Message, ve.Value)
}

// ValidationErrors represents multiple validation errors
type ValidationErrors []*ValidationError

func (ve ValidationErrors) Error() string {
	if len(ve) == 0 {
		return "no validation errors"
	}

	var messages []string
	for _, err := range ve {
		messages = append(messages, err.Error())
	}
	return strings.Join(messages, "; ")
}

// Validator defines the interface for validation
type Validator interface {
	Validate() error
}

// StringValidator validates string fields
type StringValidator struct {
	Field    string
	Value    string
	MinLen   int
	MaxLen   int
	Pattern  string
	Required bool
}

func (sv *StringValidator) Validate() error {
	if sv.Required && sv.Value == "" {
		return &ValidationError{
			Field:   sv.Field,
			Message: "field is required",
			Value:   sv.Value,
		}
	}

	if sv.MinLen > 0 && len(sv.Value) < sv.MinLen {
		return &ValidationError{
			Field:   sv.Field,
			Message: fmt.Sprintf("minimum length is %d", sv.MinLen),
			Value:   sv.Value,
		}
	}

	if sv.MaxLen > 0 && len(sv.Value) > sv.MaxLen {
		return &ValidationError{
			Field:   sv.Field,
			Message: fmt.Sprintf("maximum length is %d", sv.MaxLen),
			Value:   sv.Value,
		}
	}

	return nil
}

// IntValidator validates integer fields
type IntValidator struct {
	Field    string
	Value    int
	Min      int
	Max      int
	Required bool
}

func (iv *IntValidator) Validate() error {
	if iv.Required && iv.Value == 0 {
		return &ValidationError{
			Field:   iv.Field,
			Message: "field is required",
			Value:   iv.Value,
		}
	}

	if iv.Value < iv.Min {
		return &ValidationError{
			Field:   iv.Field,
			Message: fmt.Sprintf("value must be at least %d", iv.Min),
			Value:   iv.Value,
		}
	}

	if iv.Value > iv.Max {
		return &ValidationError{
			Field:   iv.Field,
			Message: fmt.Sprintf("value must be at most %d", iv.Max),
			Value:   iv.Value,
		}
	}

	return nil
}

// StructValidator validates struct fields using reflection
type StructValidator struct {
	Struct interface{}
}

func (sv *StructValidator) Validate() error {
	var errors ValidationErrors

	v := reflect.ValueOf(sv.Struct)
	t := reflect.TypeOf(sv.Struct)

	// Handle pointers
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return &ValidationError{
				Field:   "struct",
				Message: "struct cannot be nil",
				Value:   sv.Struct,
			}
		}
		v = v.Elem()
		t = t.Elem()
	}

	// Must be a struct
	if v.Kind() != reflect.Struct {
		return &ValidationError{
			Field:   "struct",
			Message: "value must be a struct",
			Value:   sv.Struct,
		}
	}

	// Validate each field
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		// Skip unexported fields
		if !field.CanInterface() {
			continue
		}

		// Check for validation tags
		if tag := fieldType.Tag.Get("validate"); tag != "" {
			if err := sv.validateField(field, fieldType, tag); err != nil {
				if validationErr, ok := err.(*ValidationError); ok {
					errors = append(errors, validationErr)
				}
			}
		}
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func (sv *StructValidator) validateField(field reflect.Value, fieldType reflect.StructField, tag string) error {
	fieldName := fieldType.Name

	// Parse validation tags (simplified)
	tags := strings.Split(tag, ",")

	for _, tag := range tags {
		switch {
		case tag == "required":
			if sv.isZeroValue(field) {
				return &ValidationError{
					Field:   fieldName,
					Message: "field is required",
					Value:   field.Interface(),
				}
			}
		case strings.HasPrefix(tag, "min="):
			// Parse min value and validate
			// This is a simplified implementation
		case strings.HasPrefix(tag, "max="):
			// Parse max value and validate
			// This is a simplified implementation
		}
	}

	return nil
}

func (sv *StructValidator) isZeroValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.String() == ""
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Ptr, reflect.Interface:
		return v.IsNil()
	case reflect.Slice, reflect.Map:
		return v.Len() == 0
	default:
		return false
	}
}

// ValidateStruct is a convenience function for validating structs
func ValidateStruct(s interface{}) error {
	validator := &StructValidator{Struct: s}
	return validator.Validate()
}
