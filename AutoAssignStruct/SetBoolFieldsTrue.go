package AutoSet

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
	if !isPointerOfStruct(v) {
		return errors.New("v must be pointer of struct")
	}
	setBoolFieldsTrueHelper(reflect.ValueOf(v).Elem())
	return nil
}

// setBoolFieldsTrueHelper is a helper function to recursively set boolean fields to
// true
func setBoolFieldsTrueHelper(val reflect.Value) {
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		switch field.Kind() {
		case reflect.Bool:
			field.SetBool(true)
		case reflect.Struct:
			setBoolFieldsTrueHelper(field)
		}
	}
}
