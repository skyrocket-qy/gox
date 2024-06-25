package StructHelper

import (
	"errors"
	"reflect"
	"testing"
)

// Test Struct Definitions
type NestedFrom struct {
	InnerID   int
	InnerName string
}

type FromStruct struct {
	ID      int
	Name    string
	Age     int
	Nested  NestedFrom
	Address string
}

type NestedTo struct {
	InnerID   int
	InnerName string
}

type ToStruct struct {
	ID      int
	Name    string
	Age     int
	Nested  NestedTo
	Address string
}

func TestScanStructToStruct(t *testing.T) {
	// Define the source struct
	cases := []struct {
		name        string
		from        any
		to          any
		expectedErr error
		expectedTo  any
	}{
		{
			name:        "from is nil",
			from:        nil,
			to:          &ToStruct{},
			expectedErr: errors.New("from is nil"),
			expectedTo:  &ToStruct{},
		},
		{
			name:        "to is nil",
			from:        &FromStruct{},
			to:          nil,
			expectedErr: errors.New("to is nil"),
			expectedTo:  &ToStruct{},
		},
		{
			name:        "from is a nil pointer",
			from:        (*FromStruct)(nil),
			to:          &ToStruct{},
			expectedErr: errors.New("from is a nil pointer"),
			expectedTo:  &ToStruct{},
		},
		{
			name:        "from is not a struct",
			from:        42,
			to:          &ToStruct{},
			expectedErr: errors.New("from must be a struct or pointer of struct"),
			expectedTo:  &ToStruct{},
		},
		{
			name:        "to is nil pointer",
			from:        FromStruct{ID: 1},
			to:          (*ToStruct)(nil),
			expectedErr: errors.New("to must be a non-nil pointer of struct, got type: ptr"),
			expectedTo:  (*ToStruct)(nil),
		},
		{
			name:        "to is not a struct",
			from:        FromStruct{ID: 1},
			to:          42,
			expectedErr: errors.New("to must be a non-nil pointer of struct, got type: int"),
			expectedTo:  42,
		},
		{name: "from substruct is nil pointer"},
		{name: "from and to are the same"},
		{name: "from and to not the same"},
		{name: "from embedded to to"},
		{name: "from embedded to to embedded"},
		{name: "from to to embedded"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := ScanStructToStruct(c.from, c.to)
			if err.Error() != c.expectedErr.Error() {
				t.Fatalf("expected error: %v, got: %v", c.expectedErr, err)
			}

			if c.expectedErr == nil && !reflect.DeepEqual(c.to, c.expectedTo) {
				t.Errorf("expected struct: %+v, got: %+v", c.expectedTo, c.to)
			}
		})
	}
}
