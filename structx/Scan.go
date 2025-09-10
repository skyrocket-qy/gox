package structx

import (
	"errors"
	"fmt"
	"reflect"
)

// Recursively assign value with same key including embedded fields on same layer
// It will try to convert the from type if to type not match.
func Scan(from, to any) error {
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
		if fromVal.Kind() != reflect.Struct {
			return fmt.Errorf(
				"from must be a struct or pointer of struct, got pointer of %v",
				fromVal.Kind().String(),
			)
		}
	} else if fromVal.Kind() != reflect.Struct {
		return fmt.Errorf(
			"from must be a struct or pointer of struct, got %v",
			fromVal.Kind().String(),
		)
	}

	if !isNonNilPointerOfStruct(to) {
		return fmt.Errorf(
			"to must be a non-nil pointer of struct, got type: %s",
			reflect.TypeOf(to).Kind().String(),
		)
	}

	scanHelper(fromVal, reflect.ValueOf(to).Elem())

	return nil
}

func scanHelper(from, to reflect.Value) {
	if !from.IsValid() || !to.IsValid() {
		return
	}

	for i := 0; i < to.NumField(); i++ {
		toFieldType := to.Type().Field(i)

		toField := to.Field(i)
		if isEmbedded(toFieldType) {
			scanHelper(from, toField)
		}

		if !toField.CanSet() {
			continue
		}

		fromField := from.FieldByName(toFieldType.Name)
		if !fromField.IsValid() {
			continue
		}

		toKind := toFieldType.Type.Kind()
		if toKind == reflect.Ptr {
			toField = toField.Elem()
			toKind = toField.Kind()
		}

		if !toField.CanSet() {
			continue
		}

		switch toKind {
		case reflect.Struct:
			if fromField.Type().Kind() == reflect.Ptr {
				if getElem(fromField).Type().Kind() == reflect.Struct {
					scanHelper(fromField.Elem(), toField)
				}
			}

			if fromField.Type().Kind() == reflect.Struct {
				scanHelper(fromField, toField)
			}
		case reflect.String:
			switch fromField.Type().Kind() {
			case reflect.Struct:
			default:
				toField.SetString(fmt.Sprintf("%v", fromField.Interface()))
			}
		default:
			if fromField.Type() == toFieldType.Type {
				toField.Set(fromField)
			} else if fromField.Type().AssignableTo(toFieldType.Type) {
				toField.Set(fromField.Convert(toFieldType.Type))
			}
		}
	}
}
