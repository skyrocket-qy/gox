package structx

import "reflect"

func DeepNew[V any]() *V {
	out := new(V)
	InitFields(out)

	return out
}

func InitFields(v any) {
	if !IsNonNilPointerOfStruct(v) {
		return
	}

	val := reflect.ValueOf(v).Elem()

	for i := range val.NumField() {
		field := val.Field(i)

		if field.Kind() == reflect.Pointer && field.IsNil() {
			field.Set(reflect.New(field.Type().Elem()))

			if field.Type().Elem().Kind() == reflect.Struct {
				fieldVal := field.Elem()
				initStructFields(fieldVal)
			}
		}
	}
}

func initStructFields(val reflect.Value) {
	for i := range val.NumField() {
		field := val.Field(i)

		if field.Kind() == reflect.Pointer && field.IsNil() {
			field.Set(reflect.New(field.Type().Elem()))

			if field.Type().Elem().Kind() == reflect.Struct {
				fieldVal := field.Elem()
				initStructFields(fieldVal)
			}
		}
	}
}
