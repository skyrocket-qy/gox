package structx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type SimpleStruct struct {
	Field1 int
	Field2 string
}

type NestedPointerStruct struct {
	FieldA string
	Nested *SimpleStruct
	Another *AnotherNestedStruct
}

type AnotherNestedStruct struct {
	Value bool
}

func TestDeepNew(t *testing.T) {
	// Test with a simple struct
	ptr := DeepNew[SimpleStruct]()
	assert.NotNil(t, ptr)
	assert.IsType(t, &SimpleStruct{}, ptr)
	assert.Equal(t, 0, ptr.Field1)
	assert.Equal(t, "", ptr.Field2)

	// Test with a struct containing nested pointers
	nestedPtr := DeepNew[NestedPointerStruct]()
	assert.NotNil(t, nestedPtr)
	assert.IsType(t, &NestedPointerStruct{}, nestedPtr)
	assert.NotNil(t, nestedPtr.Nested) // Nested pointer should be initialized
	assert.NotNil(t, nestedPtr.Another) // Another nested pointer should be initialized
	assert.IsType(t, &SimpleStruct{}, nestedPtr.Nested)
	assert.IsType(t, &AnotherNestedStruct{}, nestedPtr.Another)
	assert.Equal(t, 0, nestedPtr.Nested.Field1)
	assert.Equal(t, "", nestedPtr.Nested.Field2)
	assert.Equal(t, false, nestedPtr.Another.Value)
}

func TestInitFields(t *testing.T) {
	// Test with a simple struct (no pointers to initialize)
	simple := SimpleStruct{Field1: 1, Field2: "test"}
	InitFields(&simple)
	assert.Equal(t, 1, simple.Field1)
	assert.Equal(t, "test", simple.Field2)

	// Test with a struct containing nil nested pointers
	nested := NestedPointerStruct{}
	InitFields(&nested)
	assert.NotNil(t, nested.Nested)
	assert.NotNil(t, nested.Another)
	assert.IsType(t, &SimpleStruct{}, nested.Nested)
	assert.IsType(t, &AnotherNestedStruct{}, nested.Another)

	// Test with a struct containing already initialized nested pointers
	initializedNested := NestedPointerStruct{
		Nested:  &SimpleStruct{Field1: 100},
		Another: &AnotherNestedStruct{Value: true},
	}
	InitFields(&initializedNested)
	assert.Equal(t, 100, initializedNested.Nested.Field1)
	assert.Equal(t, true, initializedNested.Another.Value)

	// Test with a struct containing nested struct (not pointer)
	type StructWithNested struct {
		ID     int
		Nested SimpleStruct
	}
	withNested := StructWithNested{}
	InitFields(&withNested)
	assert.Equal(t, 0, withNested.Nested.Field1)

	// Test with nil input
	var nilPtr *SimpleStruct
	InitFields(nilPtr) // Should not panic

	// Test with non-pointer input (should not panic, but also not modify)
	InitFields(SimpleStruct{}) // Should not panic

	// Test with pointer to non-struct (should not panic, but also not modify)
	intPtr := new(int)
	InitFields(intPtr) // Should not panic
}
