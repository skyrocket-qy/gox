package structx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ScanTestSimpleFrom struct {
	IntField    int
	StringField string
	BoolField   bool
}

type ScanTestSimpleTo struct {
	IntField    int
	StringField string
	BoolField   bool
}

type ScanTestEmbeddedFrom struct {
	EmbeddedInt int
}

type ScanTestParentFrom struct {
	ScanTestEmbeddedFrom
	ParentString string
}

type ScanTestEmbeddedTo struct {
	EmbeddedInt int
}

type ScanTestParentTo struct {
	ScanTestEmbeddedTo
	ParentString string
}

func TestScan(t *testing.T) {
	// Test case 1: Basic scan, all fields match
	from1 := ScanTestSimpleFrom{IntField: 1, StringField: "hello", BoolField: true}
	to1 := ScanTestSimpleTo{}
	err := Scan(&from1, &to1)
	assert.NoError(t, err)
	assert.Equal(t, from1.IntField, to1.IntField)
	assert.Equal(t, from1.StringField, to1.StringField)
	assert.Equal(t, from1.BoolField, to1.BoolField)

	// Test case 2: Source has extra fields
	type ScanTestExtraFrom struct {
		ScanTestSimpleFrom
		ExtraField string
	}
	from2 := ScanTestExtraFrom{ScanTestSimpleFrom: ScanTestSimpleFrom{IntField: 2}, ExtraField: "extra"}
	to2 := ScanTestSimpleTo{}
	err = Scan(&from2, &to2)
	assert.NoError(t, err)
	assert.Equal(t, from2.IntField, to2.IntField)

	// Test case 3: Destination has extra fields
	type ScanTestExtraTo struct {
		ScanTestSimpleTo
		AnotherExtra string
	}
	from3 := ScanTestSimpleFrom{IntField: 3}
	to3 := ScanTestExtraTo{}
	err = Scan(&from3, &to3)
	assert.NoError(t, err)
	assert.Equal(t, from3.IntField, to3.IntField)

	// Test case 4: Embedded structs
	from4 := ScanTestParentFrom{ScanTestEmbeddedFrom: ScanTestEmbeddedFrom{EmbeddedInt: 4}, ParentString: "parent"}
	to4 := ScanTestParentTo{}
	err = Scan(&from4, &to4)
	assert.NoError(t, err)
	assert.Equal(t, from4.EmbeddedInt, to4.EmbeddedInt) // Accessing embedded field directly
	assert.Equal(t, from4.ParentString, to4.ParentString)

	// Test case 5: Type conversion (int to string)
	type ScanTestFromInt struct { Value int }
	type ScanTestToString struct { Value string }
	from5 := ScanTestFromInt{Value: 5}
	to5 := ScanTestToString{}
	err = Scan(&from5, &to5)
	assert.NoError(t, err)
	assert.Equal(t, "5", to5.Value)

	// Test case 6: Type conversion (bool to string)
	type ScanTestFromBool struct { Value bool }
	type ScanTestToStringBool struct { Value string }
	from6 := ScanTestFromBool{Value: true}
	to6 := ScanTestToStringBool{}
	err = Scan(&from6, &to6)
	assert.NoError(t, err)
	assert.Equal(t, "true", to6.Value)

	// Test case 7: Nil 'from' input
	to7 := ScanTestSimpleTo{}
	err = Scan(nil, &to7)
	assert.Error(t, err)
	assert.EqualError(t, err, "from is nil")

	// Test case 8: Nil 'to' input
	from8 := ScanTestSimpleFrom{}
	err = Scan(&from8, nil)
	assert.Error(t, err)
	assert.EqualError(t, err, "to is nil")

	// Test case 9: 'from' is a nil pointer
	var from9 *ScanTestSimpleFrom
	to9 := ScanTestSimpleTo{}
	err = Scan(from9, &to9)
	assert.Error(t, err)
	assert.EqualError(t, err, "from is a nil pointer")

	// Test case 10: 'from' is not a struct or pointer to struct
	to10 := ScanTestSimpleTo{}
	err = Scan(123, &to10)
	assert.Error(t, err)
	assert.EqualError(t, err, "from must be a struct or pointer of struct, got int")

	// Test case 11: 'to' is not a non-nil pointer to struct
	from11 := ScanTestSimpleFrom{}
	err = Scan(&from11, ScanTestSimpleTo{})
	assert.Error(t, err)
	assert.EqualError(t, err, "to must be a non-nil pointer of struct, got type: struct")

	// Test case 12: Unexported fields (should be skipped)
	type ScanTestUnexportedFrom struct { unexported int }
	type ScanTestUnexportedTo struct { unexported int }
	from12 := ScanTestUnexportedFrom{unexported: 10}
	to12 := ScanTestUnexportedTo{}
	err = Scan(&from12, &to12)
	assert.NoError(t, err)
	assert.Equal(t, 0, to12.unexported) // Should remain zero

	// Test case 13: Nested struct with pointer in 'from' and value in 'to'
	type ScanTestFromNestedPtr struct { Nested *ScanTestSimpleFrom }
	type ScanTestToNestedVal struct { Nested ScanTestSimpleTo }
	from13 := ScanTestFromNestedPtr{Nested: &ScanTestSimpleFrom{IntField: 13}}
	to13 := ScanTestToNestedVal{}
	err = Scan(&from13, &to13)
	assert.NoError(t, err)
	assert.Equal(t, 13, to13.Nested.IntField)

	// Test case 14: Nested struct with value in 'from' and pointer in 'to'
	type ScanTestFromNestedVal struct { Nested ScanTestSimpleFrom }
	type ScanTestToNestedPtr struct { Nested *ScanTestSimpleTo }
	from14 := ScanTestFromNestedVal{Nested: ScanTestSimpleFrom{IntField: 14}}
	to14 := ScanTestToNestedPtr{Nested: &ScanTestSimpleTo{}}
	err = Scan(&from14, &to14)
	assert.NoError(t, err)
	assert.Equal(t, 14, to14.Nested.IntField)

	// Test case 15: Nested struct with nil pointer in 'from'
	type ScanTestFromNestedNilPtr struct { Nested *ScanTestSimpleFrom }
	type ScanTestToNestedValNil struct { Nested ScanTestSimpleTo }
	from15 := ScanTestFromNestedNilPtr{Nested: nil}
	to15 := ScanTestToNestedValNil{}
	err = Scan(&from15, &to15)
	assert.NoError(t, err)
	assert.Equal(t, 0, to15.Nested.IntField) // Should remain zero
}