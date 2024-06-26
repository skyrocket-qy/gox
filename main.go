package main

import (
	"fmt"

	"github.com/skyrocketOoO/GoUtils/Common"
	"github.com/skyrocketOoO/GoUtils/StructHelper"
)

type NestedFrom struct {
	InnerID int
}

type FromStruct struct {
	ID     int
	Nested *NestedFrom
	EmbeddedFrom
}

type EmbeddedFrom struct {
	EmbeddedID int
}

type NestedTo struct {
	InnerID int
}
type EmbeddedTo struct {
	EmbeddedID int
}

type ToStruct struct {
	ID     int
	Nested NestedTo
}

func main() {
	from := FromStruct{
		ID:     1,
		Nested: &NestedFrom{InnerID: 2},
		EmbeddedFrom: EmbeddedFrom{
			EmbeddedID: 3,
		},
	}
	to := &FromStruct{Nested: &NestedFrom{}}

	if err := StructHelper.ScanStructToStruct(from, to); err != nil {
		fmt.Println(err.Error())
	}

	Common.PrintStruct(*to)
	fmt.Println(to)
}
