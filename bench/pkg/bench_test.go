package pkg_test

import (
	"reflect"
	"testing"
)

func toInterfaceSlice[T any](slice []T) []any {
	if slice == nil {
		return nil
	}
	result := make([]any, len(slice))
	for i, v := range slice {
		result[i] = v
	}
	return result
}

func TestToInterfaceSlice(t *testing.T) {
	// Test with []int
	intData := []int{1, 2, 3}
	intResult := toInterfaceSlice(intData)

	expectedInt := []any{1, 2, 3}
	if !reflect.DeepEqual(intResult, expectedInt) {
		t.Errorf("For []int, expected %v, got %v", expectedInt, intResult)
	}

	// Test with []int64
	int64Data := []int64{10, 20, 30}
	int64Result := toInterfaceSlice(int64Data)

	expectedInt64 := []any{int64(10), int64(20), int64(30)}
	if !reflect.DeepEqual(int64Result, expectedInt64) {
		t.Errorf("For []int64, expected %v, got %v", expectedInt64, int64Result)
	}

	// Test with []uint32
	uint32Data := []uint32{5, 15, 25}
	uint32Result := toInterfaceSlice(uint32Data)

	expectedUint32 := []any{uint32(5), uint32(15), uint32(25)}
	if !reflect.DeepEqual(uint32Result, expectedUint32) {
		t.Errorf("For []uint32, expected %v, got %v", expectedUint32, uint32Result)
	}

	// Test with an empty slice
	emptyData := []int{}
	emptyResult := toInterfaceSlice(emptyData)
	expectedEmpty := []any{}

	if len(emptyResult) != 0 {
		t.Errorf("For empty slice, expected empty, got %v", emptyResult)
	}
	// reflect.DeepEqual returns true for two empty non-nil slices.
	// toInterfaceSlice creates a non-nil slice, so this should be fine.
	if !reflect.DeepEqual(emptyResult, expectedEmpty) {
		t.Errorf("For empty slice, expected %v, got %v", expectedEmpty, emptyResult)
	}
}
