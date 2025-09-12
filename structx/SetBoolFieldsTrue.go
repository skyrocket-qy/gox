package structx

import (
	"errors"
	"reflect"
)

// SetBoolFieldsTrue recursively sets all boolean fields of a struct to true
//
//	err = SetBoolFieldsTrue(&struct)
func SetBoolFieldsTrue(v any) error {
	if v == nil {
		return errors.New("v must not be nil")
	}

	if !isNonNilPointerOfStruct(v) {
		return errors.New("v must be pointer of struct")
	}

	return setBoolFieldsTrueHelper(reflect.ValueOf(v))
}

// setBoolFieldsTrueHelper is a helper function to recursively set boolean fields to
// true. It uses the SetFields function.
func setBoolFieldsTrueHelper(v reflect.Value) error {
	// if it's a pointer, get the element
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// must be a struct
	if v.Kind() != reflect.Struct {
		return nil
	}

	// get all bool fields in the current struct
	boolFields := make(map[string]any)
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		if field.Type.Kind() == reflect.Bool {
			boolFields[field.Name] = true
		}
	}

	// set the boolean fields for the current struct
	if len(boolFields) > 0 {
		// we need a pointer to the struct to call SetFields
		if !v.CanAddr() {
			// This case should not happen if the initial call is with a pointer.
			// All fields of a struct taken from a pointer are addressable.
			return errors.New("cannot get address of struct")
		}
		if err := SetFields(v.Addr().Interface(), boolFields); err != nil {
			return err
		}
	}

	// recursively call for nested structs
	for i := 0; i < v.NumField(); i++ {
		fieldVal := v.Field(i)
		if fieldVal.Kind() == reflect.Struct {
			if err := setBoolFieldsTrueHelper(fieldVal); err != nil {
				return err
			}
		}
	}

	return nil
}
