package StructHelper

import "reflect"

func InitializeFields(a any) {
	// Get the reflect.Value of A
	val := reflect.ValueOf(a).Elem()

	// Iterate through fields of A
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		// Check if the field is a pointer and nil
		if field.Kind() == reflect.Ptr && field.IsNil() {
			// Allocate memory for the pointer field
			field.Set(reflect.New(field.Type().Elem()))

			// If the field is a struct pointer, recursively initialize its fields
			if field.Type().Elem().Kind() == reflect.Struct {
				// Get the actual value of the pointer
				fieldVal := field.Elem()

				// Initialize fields of the nested struct recursively
				initializeStructFields(fieldVal)
			}
		}
	}
}

// initializeStructFields initializes the fields of a struct using reflection recursively
func initializeStructFields(val reflect.Value) {
	// Iterate through fields of the struct
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		// Check if the field is a pointer and nil
		if field.Kind() == reflect.Ptr && field.IsNil() {
			// Allocate memory for the pointer field
			field.Set(reflect.New(field.Type().Elem()))

			// If the field is a struct pointer, recursively initialize its fields
			if field.Type().Elem().Kind() == reflect.Struct {
				// Get the actual value of the pointer
				fieldVal := field.Elem()

				// Recursively initialize fields of the nested struct
				initializeStructFields(fieldVal)
			}
		}
	}
}
