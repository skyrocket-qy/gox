package main

import (
	"fmt"
	"reflect"
)

type Address struct {
	City  string
	State string
}

type User struct {
	Name string
	Age  int
	Address
	Info struct {
		Email string
		Phone string
	}
}

func main() {
	u := User{}
	f, _ := reflect.ValueOf(u).Type().FieldByName("Address")
	fmt.Println(f.Type.Kind())
	fmt.Println(f.Anonymous)
}
