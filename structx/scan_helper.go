package structx

import (
	"fmt"
	"reflect"
)

func commonScanHelper(from, to reflect.Value) {
	if !from.IsValid() || !to.IsValid() {
		return
	}

	for i := range to.NumField() {
		toFieldType := to.Type().Field(i)

		toField := to.Field(i)
		if IsEmbedded(toFieldType) {
			commonScanHelper(from, toField)
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
				if GetElem(fromField).Type().Kind() == reflect.Struct {
					commonScanHelper(fromField.Elem(), toField)
				}
			}

			if fromField.Type().Kind() == reflect.Struct {
				commonScanHelper(fromField, toField)
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
