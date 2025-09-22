package structx

import (
	"log"
	"reflect"
)

func GetElem(v any) reflect.Value {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		return val.Elem()
	}

	return val
}

func PrintStructInfo(s any) {
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)

	if t.Kind() == reflect.Ptr {
		v = v.Elem()
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		log.Println("Input is not a struct or a pointer of struct")

		return
	}

	PrintFields(t, v)
}

// PrintFields prints the fields of a struct.
func PrintFields(t reflect.Type, v reflect.Value) {
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		fieldName := field.Name

		if IsEmbedded(field) {
			PrintFields(field.Type, value)

			continue
		}

		log.Printf("Field Name: %s\n", fieldName)
		log.Printf("Field Type: %s\n", field.Type)
		log.Printf("Field Tag: %s\n", field.Tag)
		log.Println()
	}
}

func IsEmbedded(field reflect.StructField) bool {
	return field.Anonymous
}

func IsNonNilPointerOfStruct(v any) bool {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Ptr {
		return false
	}

	if !val.IsValid() || val.IsNil() {
		return false
	}

	return val.Elem().Kind() == reflect.Struct
}
