package common

import (
	"fmt"
	"reflect"
)

func PrintStruct(s any) {
	val := reflect.ValueOf(s)
	if val.Kind() != reflect.Struct {
		panic("not a struct")
	}

	out := ""
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		fieldValue := val.Field(i)
		out += fmt.Sprintf("%s: %v\n", field.Name, fieldValue.Interface())
	}

	fmt.Println(out)
}
