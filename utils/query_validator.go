package utils

import (
	"errors"
	"reflect"
	"strings"
)

func ValidateQuery(query interface{}) error {
	if query == nil {
		return errors.New("query cannot be nil")
	}

	v := reflect.ValueOf(query)

	// Handle pointers
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return errors.New("query pointer cannot be nil")
		}
		v = v.Elem()
	}

	// Check allowed types
	switch v.Kind() {
	case reflect.Map:
		if v.Len() == 0 {
			return errors.New("query cannot be an empty map")
		}
	case reflect.Struct:
		// Validate only existing fields
		requiredFields := []string{"Name", "Email"}
		for _, field := range requiredFields {
			fv := v.FieldByName(field)
			if fv.IsValid() && isZero(fv) {
				return errors.New("missing required field: " + field)
			}
		}
	default:
		return errors.New("query must be a map or a struct")
	}

	// Additional validation for email in structs
	if v.Kind() == reflect.Struct {
		emailField := v.FieldByName("Email")
		if emailField.IsValid() && !isValidEmail(emailField.String()) {
			return errors.New("invalid email format")
		}
	}

	// Additional map-specific validation
	if v.Kind() == reflect.Map {
		nameValue := v.MapIndex(reflect.ValueOf("name"))
		if nameValue.IsValid() && nameValue.Kind() == reflect.String {
			name := nameValue.String()
			if len(name) < 3 {
				return errors.New("name must be at least 3 characters")
			}
		}
	}

	return nil
}

// Helper to check if a value is zero
func isZero(v reflect.Value) bool {
	if !v.IsValid() {
		return true
	}
	switch v.Kind() {
	case reflect.Func, reflect.Map, reflect.Slice, reflect.Ptr, reflect.Interface:
		return v.IsNil()
	case reflect.Array:
		for i := 0; i < v.Len(); i++ {
			if !isZero(v.Index(i)) {
				return false
			}
		}
		return true
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			if !isZero(v.Field(i)) {
				return false
			}
		}
		return true
	default:
		zero := reflect.Zero(v.Type())
		return reflect.DeepEqual(v.Interface(), zero.Interface())
	}
}

func isValidEmail(email string) bool {
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}
