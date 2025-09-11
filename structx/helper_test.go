package structx

import (
	"bytes"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	Embedded

	Field1 int
	Field2 string `json:"f2"`
}

type Embedded struct {
	EmbeddedField bool
}

func TestGetElem(t *testing.T) {
	// Test with a pointer
	ptr := &TestStruct{Field1: 1}
	elem := getElem(ptr)
	assert.Equal(t, reflect.Struct, elem.Kind())
	assert.Equal(t, int64(1), elem.FieldByName("Field1").Int())

	// Test with a non-pointer
	val := TestStruct{Field1: 2}
	elem = getElem(val)
	assert.Equal(t, reflect.Struct, elem.Kind())
	assert.Equal(t, int64(2), elem.FieldByName("Field1").Int())
}

func TestIsEmbedded(t *testing.T) {
	// Test with an embedded field
	field, _ := reflect.TypeOf(TestStruct{}).FieldByName("Embedded")
	assert.True(t, isEmbedded(field))

	// Test with a non-embedded field
	field, _ = reflect.TypeOf(TestStruct{}).FieldByName("Field1")
	assert.False(t, isEmbedded(field))
}

func TestIsNonNilPointerOfStruct(t *testing.T) {
	// Test with a non-nil pointer to a struct
	ptr := &TestStruct{}
	assert.True(t, isNonNilPointerOfStruct(ptr))

	// Test with a nil pointer to a struct
	var nilPtr *TestStruct
	assert.False(t, isNonNilPointerOfStruct(nilPtr))

	// Test with a non-pointer struct
	val := TestStruct{}
	assert.False(t, isNonNilPointerOfStruct(val))

	// Test with a pointer to a non-struct type
	intPtr := new(int)
	assert.False(t, isNonNilPointerOfStruct(intPtr))

	// Test with nil interface
	var nilInterface any
	assert.False(t, isNonNilPointerOfStruct(nilInterface))
}

func TestPrintStructInfo(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	p := Person{Name: "John", Age: 30}

	// Redirect log output to a buffer
	var buf bytes.Buffer
	log.SetOutput(&buf)

	defer func() {
		log.SetOutput(os.Stderr)
	}()

	PrintStructInfo(p)

	output := buf.String()

	assert.Contains(t, output, "Field Name: Name")
	assert.Contains(t, output, "Field Type: string")
	assert.Contains(t, output, "Field Name: Age")
	assert.Contains(t, output, "Field Type: int")
}

func TestPrintFields(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	p := Person{Name: "John", Age: 30}

	// Redirect log output to a buffer
	var buf bytes.Buffer
	log.SetOutput(&buf)

	defer func() {
		log.SetOutput(os.Stderr)
	}()

	printFields(reflect.TypeOf(p), reflect.ValueOf(p))

	output := buf.String()

	assert.Contains(t, output, "Field Name: Name")
	assert.Contains(t, output, "Field Type: string")
	assert.Contains(t, output, "Field Name: Age")
	assert.Contains(t, output, "Field Type: int")
}
