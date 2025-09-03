package structx

import "reflect"

func DeepNew[V any]() *V {
	out := new(V)
	InitFields(out)
	return out
}

func InitFields(v any) {
	if v == nil || reflect.ValueOf(v).IsNil() {
		typ := reflect.TypeOf(v).Elem()
		newA := reflect.New(typ).Interface()
		v = newA
	}
	val := reflect.ValueOf(v).Elem()

	for i := 0; i < val.NumField(); i++ {
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
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		if field.Kind() == reflect.Ptr && field.IsNil() {
			field.Set(reflect.New(field.Type().Elem()))

			if field.Type().Elem().Kind() == reflect.Struct {
				fieldVal := field.Elem()
				initStructFields(fieldVal)
			}
		}
	}
}
