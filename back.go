package main

import "reflect"

func setAllBoolFieldsToTrue(v interface{}) {
	val := reflect.ValueOf(v).Elem()
	setBoolFields(val)
}

// setBoolFields is a helper function to recursively set boolean fields to true
func setBoolFields(val reflect.Value) {
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		switch field.Kind() {
		case reflect.Bool:
			field.SetBool(true)
		case reflect.Struct:
			setBoolFields(field)
		}
	}
}

func copyStruct(src, dst interface{}) {
	srcVal := reflect.ValueOf(src).Elem()
	dstVal := reflect.ValueOf(dst).Elem()
	copyFields(srcVal, dstVal)
}

func copyFields(srcVal, dstVal reflect.Value) {
	for i := 0; i < srcVal.NumField(); i++ {
		srcField := srcVal.Field(i)
		dstField := dstVal.Field(i)

		if !srcField.CanInterface() || !dstField.CanSet() {
			continue
		}

		switch srcField.Kind() {
		case reflect.Struct:
			copyFields(srcField, dstField)
		case reflect.Ptr:
			if !srcField.IsNil() {
				dstField.Set(reflect.New(srcField.Elem().Type()))
				copyFields(srcField.Elem(), dstField.Elem())
			}
		default:
			dstField.Set(srcField)
		}
	}
}
