package structutil

import (
	"errors"
	"fmt"
	"reflect"
)

// type Person struct {
// 	Name   string
// 	Age    int
// 	Email  string
// 	Active bool
// }

func Print(s interface{}) (string, error) {
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

// func main() {
// 	p := Person{Name: "Alice", Age: 30, Email: "alice@example.com", Active: true}
// 	out, err := Print(p)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(out)
// }
