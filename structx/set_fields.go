package structx

import (
	"errors"
	"fmt"
	"reflect"
)

// SetFields sets fields of a struct with the given values.
// The v must be a non-nil pointer of struct.
// The values is a map of field names to their values.
func SetFields(v any, values map[string]any) error {
	if v == nil {
		return errors.New("v must not be nil")
	}

	if !IsNonNilPointerOfStruct(v) {
		return errors.New("v must be a non-nil pointer to a struct")
	}

	val := reflect.ValueOf(v).Elem()
	for name, value := range values {
		field := val.FieldByName(name)
		if !field.IsValid() {
			return fmt.Errorf("field %s not found", name)
		}

		if !field.CanSet() {
			return fmt.Errorf("field %s cannot be set", name)
		}

		valueToSet := reflect.ValueOf(value)
		if field.Type() != valueToSet.Type() {
			// try to convert type
			if valueToSet.CanConvert(field.Type()) {
				valueToSet = valueToSet.Convert(field.Type())
			} else {
				return fmt.Errorf("type mismatch for field %s: expected %s, got %s", name, field.Type(), valueToSet.Type())
			}
		}

		field.Set(valueToSet)
	}

	return nil
}
