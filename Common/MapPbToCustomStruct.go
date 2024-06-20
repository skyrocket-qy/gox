package Common

import (
	"fmt"
	"reflect"
)

// Auto map value with same key and type
func MapStructToStruct(from interface{}, to interface{}) error {
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

	fromVal, toVal := getElem(from), getElem(to)

	for i := 0; i < fromVal.NumField(); i++ {
		fromField := fromVal.Type().Field(i)
		toField := toVal.FieldByName(fromField.Name)

		// Check if the field exists in the custom struct and the types are assignable
		if toField.IsValid() && toField.CanSet() &&
			toField.Type() == fromField.Type {
			toField.Set(fromVal.Field(i))
		}
	}

	return nil
}

func getElem(in any) reflect.Value {
	if reflect.TypeOf(in).Kind() == reflect.Ptr {
		return reflect.ValueOf(in).Elem()
	} else {
		return reflect.ValueOf(in)
	}
}
