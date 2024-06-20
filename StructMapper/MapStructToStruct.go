package StructMapper

import (
	"fmt"
	"reflect"
)

// Auto map value with same key and type including embedded fields
func MapStructToStruct(from any, to any) error {
	if from == nil || to == nil {
		return fmt.Errorf("from or to must not be nil")
	}
	if !((reflect.TypeOf(from).Kind() == reflect.Struct) ||
		(reflect.TypeOf(from).Kind() == reflect.Ptr)) {
		return fmt.Errorf("from must be a struct or pointer of struct")
	}
	if reflect.TypeOf(to).Kind() != reflect.Ptr &&
		reflect.ValueOf(to).Elem().Kind() != reflect.Struct {
		return fmt.Errorf("to must be a pointer of struct")
	}

	fromStruct, toStruct := getElem(from), getElem(to)

	for i := 0; i < toStruct.NumField(); i++ {
		toField := toStruct.Type().Field(i)
		toVal := toStruct.Field(i)
		fromField := fromStruct.FieldByName(toField.Name)

		if fromField.IsValid() && toVal.CanSet() {
			switch toField.Type.Kind() {
			case fromField.Type().Kind():
				toVal.Set(fromField)
			case reflect.String:
				toVal.SetString(fmt.Sprintf("%v", fromField.Interface()))
			}
		}
	}

	return nil
}
