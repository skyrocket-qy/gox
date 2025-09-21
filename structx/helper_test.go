package structx_test

import (
	"bytes"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/skyrocket-qy/gox/structx"
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
	elem := structx.GetElem(ptr)
	assert.Equal(t, reflect.Struct, elem.Kind())
	assert.Equal(t, int64(1), elem.FieldByName("Field1").Int())

	// Test with a non-pointer
	val := TestStruct{Field1: 2}
	elem = structx.GetElem(val)
	assert.Equal(t, reflect.Struct, elem.Kind())
	assert.Equal(t, int64(2), elem.FieldByName("Field1").Int())
}

func TestIsEmbedded(t *testing.T) {
	t.Parallel()
	// Test with an embedded field
	field, _ := reflect.TypeOf(TestStruct{}).FieldByName("Embedded")
	assert.True(t, structx.IsEmbedded(field))

	// Test with a non-embedded field
	field, _ = reflect.TypeOf(TestStruct{}).FieldByName("Field1")
	assert.False(t, structx.IsEmbedded(field))
}

func TestIsNonNilPointerOfStruct(t *testing.T) {
	t.Parallel()
	// Test with a non-nil pointer to a struct
	ptr := &TestStruct{}
	assert.True(t, structx.IsNonNilPointerOfStruct(ptr))

	// Test with a nil pointer to a struct
	var nilPtr *TestStruct
	assert.False(t, structx.IsNonNilPointerOfStruct(nilPtr))

	// Test with a non-pointer struct
	val := TestStruct{}
	assert.False(t, structx.IsNonNilPointerOfStruct(val))

	// Test with a pointer to a non-struct type
	intPtr := new(int)
	assert.False(t, structx.IsNonNilPointerOfStruct(intPtr))

	// Test with nil interface
	var nilInterface any
	assert.False(t, structx.IsNonNilPointerOfStruct(nilInterface))
}

func TestPrintStructInfo(t *testing.T) {
	t.Parallel()

	type Person struct {
		Name string
		Age  int
	}

	t.Run("Input is a struct", func(t *testing.T) {
		t.Parallel()

		p := Person{Name: "John", Age: 30}

		var buf bytes.Buffer
		log.SetOutput(&buf)

		defer func() {
			log.SetOutput(os.Stderr)
		}()

		structx.PrintStructInfo(p)

		output := buf.String()
		assert.Contains(t, output, "Field Name: Name")
		assert.Contains(t, output, "Field Type: string")
		assert.Contains(t, output, "Field Name: Age")
		assert.Contains(t, output, "Field Type: int")
	})

	t.Run("Input is a pointer to a struct", func(t *testing.T) {
		t.Parallel()

		p := &Person{Name: "John", Age: 30}

		var buf bytes.Buffer
		log.SetOutput(&buf)

		defer func() {
			log.SetOutput(os.Stderr)
		}()

		structx.PrintStructInfo(p)

		output := buf.String()
		assert.Contains(t, output, "Field Name: Name")
		assert.Contains(t, output, "Field Type: string")
		assert.Contains(t, output, "Field Name: Age")
		assert.Contains(t, output, "Field Type: int")
	})

	t.Run("Input is not a struct", func(t *testing.T) {
		t.Parallel()

		var buf bytes.Buffer
		log.SetOutput(&buf)

		defer func() {
			log.SetOutput(os.Stderr)
		}()

		structx.PrintStructInfo(123)

		output := buf.String()
		assert.Contains(t, output, "Input is not a struct or a pointer of struct")
	})
}

func TestPrintFields(t *testing.T) {
	t.Parallel()

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

	structx.PrintFields(reflect.TypeOf(p), reflect.ValueOf(p))

	output := buf.String()

	assert.Contains(t, output, "Field Name: Name")
	assert.Contains(t, output, "Field Type: string")
	assert.Contains(t, output, "Field Name: Age")
	assert.Contains(t, output, "Field Type: int")
}
