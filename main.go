package main

import (
	"fmt"

	"github.com/skyrocketOoO/GoUtils/StructHelper"
)

func main() {
	exampleFunction()
}

func exampleFunction() {
	nestedFunction()
}

type A struct {
	Name string
	*Aa
}

type B struct {
	Name string
}

type Aa struct {
	Number int
}

func nestedFunction() {
	a := new(A)
	b := new(B)
	if err := StructHelper.ScanStructToStruct(b, a); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(a)
}
