package Common

import (
	"errors"
	"fmt"
	"reflect"
)

func FPrintStruct(s any) (string, error) {
	val := reflect.ValueOf(s)
	if val.Kind() != reflect.Struct {
		return "", errors.New("not a struct")
	}

	out := ""
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		fieldValue := val.Field(i)
		out += fmt.Sprintf("%s: %v\n", field.Name, fieldValue.Interface())
	}
	return out, nil
}
