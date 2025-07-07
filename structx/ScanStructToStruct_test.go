package structx

import (
	"errors"
	"reflect"
	"testing"
)

// Test Struct Definitions
type NestedFrom struct {
	InnerID int
}

type FromStruct struct {
	ID     int
	Nested *NestedFrom
	EmbeddedFrom
}

type EmbeddedFrom struct {
	EmbeddedID     int
	EmbeddedString string
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
	EmbeddedTo
	EmbeddedString string
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
		},
		{
			name:        "to is nil",
			from:        &FromStruct{},
			to:          nil,
			expectedErr: errors.New("to is nil"),
		},
		{
			name:        "from is a nil pointer",
			from:        (*FromStruct)(nil),
			to:          &ToStruct{},
			expectedErr: errors.New("from is a nil pointer"),
		},
		{
			name:        "from is not a struct",
			from:        42,
			to:          &ToStruct{},
			expectedErr: errors.New("from must be a struct or pointer of struct"),
		},
		{
			name: "to is nil pointer",
			from: FromStruct{ID: 1},
			to:   (*ToStruct)(nil),
			expectedErr: errors.New(
				"to must be a non-nil pointer of struct, got type: ptr",
			),
		},
		{
			name: "to is not a struct",
			from: FromStruct{ID: 1},
			to:   42,
			expectedErr: errors.New(
				"to must be a non-nil pointer of struct, got type: int",
			),
		},
		// {
		// 	name:        "from substruct is nil pointer",
		// 	from:        FromStruct{},
		// 	to:          &ToStruct{},
		// 	expectedErr: errors.New("from substruct is non-nil pointer, key: Nested"),
		// },
		{
			name: "from and to are the same struct",
			from: &FromStruct{
				ID:     1,
				Nested: &NestedFrom{InnerID: 2},
				EmbeddedFrom: EmbeddedFrom{
					EmbeddedID: 3,
				},
			},
			to: &FromStruct{
				Nested: &NestedFrom{},
			},
			expectedErr: nil,
			expectedTo: &FromStruct{
				ID:     1,
				Nested: &NestedFrom{InnerID: 2},
				EmbeddedFrom: EmbeddedFrom{
					EmbeddedID: 3,
				},
			},
		},
		{
			name: "from and to not the same",
			from: &FromStruct{
				ID:     1,
				Nested: &NestedFrom{InnerID: 2},
				EmbeddedFrom: EmbeddedFrom{
					EmbeddedID: 3,
				},
			},
			to:          &ToStruct{},
			expectedErr: nil,
			expectedTo: &ToStruct{
				ID: 1,
				Nested: NestedTo{
					InnerID: 2,
				},
				EmbeddedTo: EmbeddedTo{
					EmbeddedID: 3,
				},
			},
		},
		{
			name: "from embedded to to",
			from: &FromStruct{
				EmbeddedFrom: EmbeddedFrom{EmbeddedString: "embstr"},
			},
			to:          &ToStruct{},
			expectedErr: nil,
			expectedTo: &ToStruct{
				EmbeddedString: "embstr",
			},
		},
		{
			name: "from embedded to to embedded",
			from: &FromStruct{
				EmbeddedFrom: EmbeddedFrom{EmbeddedID: 3},
			},
			to:          &ToStruct{EmbeddedTo: EmbeddedTo{}},
			expectedErr: nil,
			expectedTo:  &ToStruct{EmbeddedTo: EmbeddedTo{EmbeddedID: 3}},
		},
		{
			name: "from to to embedded",
			from: &ToStruct{
				EmbeddedString: "embstr",
			},
			to:          &FromStruct{},
			expectedErr: nil,
			expectedTo: &FromStruct{
				EmbeddedFrom: EmbeddedFrom{EmbeddedString: "embstr"},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := ScanStructToStruct(c.from, c.to)
			if c.expectedErr != nil {
				if err.Error() != c.expectedErr.Error() {
					t.Fatalf("expected error: %v, got: %v", c.expectedErr, err)
				}
			}

			if c.expectedErr == nil && !reflect.DeepEqual(c.to, c.expectedTo) {
				t.Errorf("expected struct: %+v, got: %+v", c.expectedTo, c.to)
			}
		})
	}
}
