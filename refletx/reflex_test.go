package refletx

import (
	"reflect"
	"strings"
	"testing"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsNil(t *testing.T) {
	// Test with nil interface
	var nilInterface any
	assert.True(t, IsNil(reflect.ValueOf(nilInterface)))

	// Test with nil pointer
	var nilPointer *int
	assert.True(t, IsNil(reflect.ValueOf(nilPointer)))

	// Test with non-nil value
	nonNilValue := 10
	assert.False(t, IsNil(reflect.ValueOf(nonNilValue)))

	// Test with nil slice
	var nilSlice []int
	assert.True(t, IsNil(reflect.ValueOf(nilSlice)))

	// Test with empty slice
	emptySlice := []int{}
	assert.False(t, IsNil(reflect.ValueOf(emptySlice)))

	// Test with nil map
	var nilMap map[string]int
	assert.True(t, IsNil(reflect.ValueOf(nilMap)))

	// Test with empty map
	emptyMap := map[string]int{}
	assert.False(t, IsNil(reflect.ValueOf(emptyMap)))

	// Test with nil channel
	var nilChan chan int
	assert.True(t, IsNil(reflect.ValueOf(nilChan)))

	// Test with nil func
	var nilFunc func()
	assert.True(t, IsNil(reflect.ValueOf(nilFunc)))

	// Test with nil interface of a specific type
	var err error
	assert.True(t, IsNil(reflect.ValueOf(err)))

	err = gerror.New("test error")
	assert.False(t, IsNil(reflect.ValueOf(err)))
}

func TestValueOf(t *testing.T) {
	type testStruct struct {
		Field int
	}

	// Test with non-pointer, deep=false
	val := testStruct{Field: 1}
	v := ValueOf(val, false)
	assert.Equal(t, reflect.Struct, v.Kind())
	assert.Equal(t, val.Field, v.FieldByName("Field").Interface())

	// Test with non-pointer, deep=true
	v = ValueOf(val, true)
	assert.Equal(t, reflect.Struct, v.Kind())
	assert.Equal(t, val.Field, v.FieldByName("Field").Interface())

	// Test with single pointer, deep=false
	ptrVal := &val
	v = ValueOf(ptrVal, false)
	assert.Equal(t, reflect.Struct, v.Kind())
	assert.Equal(t, val.Field, v.FieldByName("Field").Interface())

	// Test with single pointer, deep=true
	v = ValueOf(ptrVal, true)
	assert.Equal(t, reflect.Struct, v.Kind())
	assert.Equal(t, val.Field, v.FieldByName("Field").Interface())

	// Test with double pointer, deep=false
	doublePtrVal := &ptrVal
	v = ValueOf(doublePtrVal, false)
	assert.Equal(t, reflect.Ptr, v.Kind()) // Should still be pointer
	assert.Equal(t, reflect.Struct, v.Elem().Kind())

	// Test with double pointer, deep=true
	v = ValueOf(doublePtrVal, true)
	assert.Equal(t, reflect.Struct, v.Kind())
	assert.Equal(t, val.Field, v.FieldByName("Field").Interface())

	// Test with nil pointer
	var nilPtr *testStruct

	v = ValueOf(nilPtr, false)
	assert.Equal(t, reflect.Invalid, v.Kind()) // Should be invalid for nil pointer

	v = ValueOf(nilPtr, true)
	assert.Equal(t, reflect.Invalid, v.Kind()) // Should be invalid for nil pointer
}

type TestStruct struct {
	Name  string
	Value int
}

func (ts *TestStruct) GetName() string {
	return ts.Name
}

func (ts *TestStruct) SetName(name string) {
	ts.Name = name
}

func (ts *TestStruct) Sum(a, b int) int {
	return a + b
}

func TestGetField(t *testing.T) {
	s := TestStruct{Name: "test", Value: 10}

	// Test existing field
	name, err := GetField[string](&s, "Name")
	assert.NoError(t, err)
	assert.Equal(t, "test", *name)

	value, err := GetField[int](&s, "Value")
	assert.NoError(t, err)
	assert.Equal(t, 10, *value)

	// Test non-existing field
	_, err = GetField[string](&s, "NonExistent")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "no such field")

	// Test type mismatch
	_, err = GetField[int](&s, "Name")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "assert field type fail")
}

func TestCallMethod(t *testing.T) {
	s := TestStruct{Name: "initial"}

	// Test calling a method with no arguments and a return value
	out, err := CallMethod(&s, "GetName", []reflect.Value{})
	assert.NoError(t, err)
	assert.Len(t, out, 1)
	assert.Equal(t, "initial", out[0].String())

	// Test calling a method with arguments and no return value
	out, err = CallMethod(&s, "SetName", []reflect.Value{reflect.ValueOf("new name")})
	assert.NoError(t, err)
	assert.Empty(t, out) // SetName has no return values
	assert.Equal(t, "new name", s.Name)

	// Test calling a method with arguments and a return value
	out, err = CallMethod(&s, "Sum", []reflect.Value{reflect.ValueOf(5), reflect.ValueOf(7)})
	assert.NoError(t, err)
	assert.Len(t, out, 1)
	assert.Equal(t, 12, int(out[0].Int()))

	// Test calling a non-existent method
	_, err = CallMethod(&s, "NonExistentMethod", []reflect.Value{})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "method is not exist")

	// Test calling a method with wrong number of arguments
	_, err = CallMethod(
		&s,
		"SetName",
		[]reflect.Value{reflect.ValueOf("one"), reflect.ValueOf("two")},
	)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "the method arguments is out of index")

	// Test calling a method with wrong argument type (reflect.Call will panic, so we expect an
	// error from our wrapper) This case is handled by reflect.Call panicking if types don't match,
	// which is then recovered by our IsNil.
	// For CallField, the error comes from NumIn() check.
	_, err = CallMethod(&s, "SetName", []reflect.Value{reflect.ValueOf(123)})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to call method")
}

func TestGetMap(t *testing.T) {
	type InnerStruct struct {
		ID   int
		Name string
	}

	type OuterStruct struct {
		Field1 string
		Field2 int
		Field3 InnerStruct
		Field4 *InnerStruct
	}

	// Test with a simple struct
	s1 := OuterStruct{Field1: "value1", Field2: 10}
	m1, err := GetMap[any](&s1)
	assert.NoError(t, err)
	assert.Equal(t, "value1", m1["Field1"])
	assert.Equal(t, 10, m1["Field2"])

	// Test with nested struct
	s2 := OuterStruct{
		Field1: "value1",
		Field3: InnerStruct{ID: 1, Name: "inner"},
	}
	m2, err := GetMap[any](&s2)
	assert.NoError(t, err)
	assert.Equal(t, "value1", m2["Field1"])
	assert.Equal(t, InnerStruct{ID: 1, Name: "inner"}, m2["Field3"])

	// Test with nested pointer struct
	s3 := OuterStruct{
		Field1: "value1",
		Field4: &InnerStruct{ID: 2, Name: "inner_ptr"},
	}
	m3, err := GetMap[any](&s3)
	assert.NoError(t, err)
	assert.Equal(t, "value1", m3["Field1"])
	assert.Equal(t, &InnerStruct{ID: 2, Name: "inner_ptr"}, m3["Field4"])

	// Test with type assertion failure
	type AnotherType struct {
		X int
	}

	s4 := struct {
		FieldA string
		FieldB AnotherType
	}{
		FieldA: "hello",
		FieldB: AnotherType{X: 100},
	}
	_, err = GetMap[int](&s4) // Expecting int, but FieldA is string
	require.Error(t, err)
	assert.Contains(t, err.Error(), "assert field type fail: FieldA")

	// Test with nil input
	var nilStruct *OuterStruct

	_, err = GetMap[any](nilStruct)
	require.Error(t, err) // reflections.ItemsDeep returns error for nil input
}

func TestGetFunctionName(t *testing.T) {
	// Test with shortName = false (full name)
	func() {
		name := GetFunctionName(TestGetFunctionName, false)
		assert.Contains(t, name, "refletx.TestGetFunctionName")
	}()

	// Test with shortName = true (short name)
	func() {
		name := GetFunctionName(TestGetFunctionName, true)
		assert.Equal(t, "TestGetFunctionName", name)
	}()

	// Test with anonymous function
	anonFunc := func() {}
	name := GetFunctionName(anonFunc, false)
	// The full name of an anonymous function can vary (e.g., .func1, .func2, etc.)
	// We'll assert that it contains "TestGetFunctionName.func"
	assert.Contains(t, name, "TestGetFunctionName.func")

	name = GetFunctionName(anonFunc, true)
	// The short name of an anonymous function can vary (e.g., func1, func2, etc.)
	// We'll assert that it starts with "func"
	assert.True(t, strings.HasPrefix(name, "func"))
}

func TestGetCallerName(t *testing.T) {
	// Helper function to test GetCallerName
	helperFunc := func(skip int, shortName bool) string {
		return GetCallerName(skip, shortName)
	}

	// Test GetCallerName from within a test function
	// skip = 0: GetCallerName itself
	// skip = 1: helperFunc
	// skip = 2: TestGetCallerName
	name := helperFunc(2, false)
	assert.Contains(t, name, "refletx.TestGetCallerName")

	name = helperFunc(2, true)
	assert.Equal(t, "TestGetCallerName", name)

	// Test with a higher skip that goes out of stack
	name = helperFunc(100, false)
	assert.Empty(t, name)
}

func TestGetCurrentCallerShortName(t *testing.T) {
	name := GetCurrentCallerShortName()
	assert.Equal(t, "TestGetCurrentCallerShortName", name)
}

func TestGetCurrentCallerFullName(t *testing.T) {
	name := GetCurrentCallerFullName()
	assert.Contains(t, name, "refletx.TestGetCurrentCallerFullName")
}
