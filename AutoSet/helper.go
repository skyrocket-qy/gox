package AutoSet

import (
	"fmt"
	"reflect"
)

func getElem(in any) reflect.Value {
	if reflect.TypeOf(in).Kind() == reflect.Ptr {
		return reflect.ValueOf(in).Elem()
	} else {
		return reflect.ValueOf(in)
	}
}

func PrintStructInfo(s interface{}) {
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)

	if t.Kind() == reflect.Ptr {
		v = v.Elem()
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		fmt.Println("Input is not a struct or a pointer of struct")
		return
	}

	printFields(t, v)
}

// printFields prints the fields of a struct
func printFields(t reflect.Type, v reflect.Value) {
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		fieldName := field.Name

		if isEmbedded(field) {
			printFields(field.Type, value)
			continue
		}

		fmt.Printf("Field Name: %s\n", fieldName)
		fmt.Printf("Field Type: %s\n", field.Type)
		fmt.Printf("Field Tag: %s\n", field.Tag)
		fmt.Println()
	}
}

func isEmbedded(field reflect.StructField) bool {
	return field.Anonymous && field.Type.Kind() == reflect.Struct
}

func isPointerOfStruct(v any) bool {
	if reflect.TypeOf(v).Kind() != reflect.Ptr {
		return false
	}
	if reflect.ValueOf(v).Elem().Kind() != reflect.Struct {
		return false
	}
	return true
}
