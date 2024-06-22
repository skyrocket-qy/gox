package StructHelper

import (
	"fmt"
	"reflect"
)

// Assign value with same key including embedded fields on same layer
func AssignStructToStruct(from any, to any) error {
	if from == nil || to == nil {
		return fmt.Errorf("from or to must not be nil")
	}
	if !(reflect.TypeOf(from).Kind() == reflect.Struct) &&
		!isPointerOfStruct(from) {
		return fmt.Errorf("from must be a struct or pointer of struct")
	}
	if !isPointerOfStruct(to) {
		return fmt.Errorf("to must be a pointer of struct")
	}

	mapStructToStructHelper(getElem(from), reflect.ValueOf(to).Elem())

	return nil
}

func mapStructToStructHelper(from, to reflect.Value) {
	for i := 0; i < to.NumField(); i++ {
		toFieldType := to.Type().Field(i)
		toField := to.Field(i)
		if isEmbedded(toFieldType) {
			mapStructToStructHelper(from, toField)
		}
		if !toField.CanSet() {
			continue
		}
		fromField := from.FieldByName(toFieldType.Name)
		if !fromField.IsValid() {
			continue
		}

		switch toFieldType.Type.Kind() {
		case reflect.Struct:
			if fromField.Type().Kind() == reflect.Struct {
				mapStructToStructHelper(fromField, toField)
			}
		case fromField.Type().Kind():
			toField.Set(fromField)
		case reflect.String:
			switch fromField.Type().Kind() {
			case reflect.Struct:
			default:
				fmt.Println(fromField.Type().Kind(), toFieldType.Type.Kind())
				toField.SetString(fmt.Sprintf("%v", fromField.Interface()))
			}
		}
	}
}
