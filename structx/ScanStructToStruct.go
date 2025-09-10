package structx

import (
	"errors"
	"fmt"
	"reflect"
)

// Recursively assign value with same key including embedded fields on same layer
// It will try to convert the from type if to type not match.
func ScanStructToStruct(from, to any) error {
	if from == nil {
		return errors.New("from is nil")
	}

	if to == nil {
		return errors.New("to is nil")
	}

	fromVal := reflect.ValueOf(from)
	if fromVal.Kind() == reflect.Ptr {
		if fromVal.IsNil() {
			return errors.New("from is a nil pointer")
		}

		fromVal = fromVal.Elem()
	}

	if fromVal.Kind() != reflect.Struct {
		return errors.New("from must be a struct or pointer of struct")
	}

	if !isNonNilPointerOfStruct(to) {
		return fmt.Errorf(
			"to must be a non-nil pointer of struct, got type: %s",
			reflect.TypeOf(to).Kind().String(),
		)
	}

	commonScanHelper(fromVal, reflect.ValueOf(to).Elem())

	return nil
}
