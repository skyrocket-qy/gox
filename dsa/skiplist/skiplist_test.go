package skiplist

import (
	"reflect"
	"testing"
)

func intComparator(a, b interface{}) int {
	aInt := a.(int)
	bInt := b.(int)
	if aInt < bInt {
		return -1
	} else if aInt > bInt {
		return 1
	}
	return 0
}

func TestSkipList_InsertAndSearch(t *testing.T) {
	sl := New(intComparator)

	sl.Insert(10)
	sl.Insert(20)
	sl.Insert(5)

	if sl.Search(10) != 10 {
		t.Errorf("Expected to find 10, but got %v", sl.Search(10))
	}
	if sl.Search(20) != 20 {
		t.Errorf("Expected to find 20, but got %v", sl.Search(20))
	}
	if sl.Search(5) != 5 {
		t.Errorf("Expected to find 5, but got %v", sl.Search(5))
	}
	if sl.Search(15) != nil {
		t.Errorf("Expected not to find 15, but got %v", sl.Search(15))
	}

	if sl.Len() != 3 {
		t.Errorf("Expected length 3, but got %d", sl.Len())
	}
}

func TestSkipList_Delete(t *testing.T) {
	sl := New(intComparator)

	sl.Insert(10)
	sl.Insert(20)
	sl.Insert(5)
	sl.Insert(15)

	if sl.Len() != 4 {
		t.Errorf("Expected length 4, but got %d", sl.Len())
	}

	sl.Delete(10)

	if sl.Search(10) != nil {
		t.Errorf("Expected not to find 10 after deletion, but got %v", sl.Search(10))
	}
	if sl.Len() != 3 {
		t.Errorf("Expected length 3 after deletion, but got %d", sl.Len())
	}

	sl.Delete(5)
	sl.Delete(20)
	sl.Delete(15)

	if sl.Len() != 0 {
		t.Errorf("Expected length 0 after all deletions, but got %d", sl.Len())
	}
	if sl.Search(15) != nil {
		t.Errorf("Expected not to find 15 after deletion, but got %v", sl.Search(15))
	}
}

func TestSkipList_Update(t *testing.T) {
	sl := New(intComparator)

	sl.Insert(10)
	sl.Insert(20)
	sl.Insert(10) // Update 10

	if sl.Search(10) != 10 {
		t.Errorf("Expected to find updated 10, but got %v", sl.Search(10))
	}
	if sl.Len() != 2 {
		t.Errorf("Expected length 2 after update, but got %d", sl.Len())
	}
}

func TestSkipList_GetRangeByValue(t *testing.T) {
	sl := New(intComparator)

	sl.Insert(5)
	sl.Insert(10)
	sl.Insert(15)
	sl.Insert(20)
	sl.Insert(25)

	// Test full range
	expected := []interface{}{5, 10, 15, 20, 25}
	actual := sl.GetRangeByValue(0, 30)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	// Test partial range
	expected = []interface{}{10, 15, 20}
	actual = sl.GetRangeByValue(10, 20)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	// Test range with no elements
	expected = []interface{}(nil)
	actual = sl.GetRangeByValue(30, 40)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	// Test range with single element
	expected = []interface{}{15}
	actual = sl.GetRangeByValue(15, 15)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	// Test range with min > max
	expected = []interface{}(nil)
	actual = sl.GetRangeByValue(20, 10)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestSkipList_EdgeCases(t *testing.T) {
	sl := New(intComparator)

	// Delete from empty list
	sl.Delete(10)
	if sl.Len() != 0 {
		t.Errorf("Expected length 0 for empty list after delete, but got %d", sl.Len())
	}

	// Search in empty list
	if sl.Search(10) != nil {
		t.Errorf("Expected nil for search in empty list, but got %v", sl.Search(10))
	}

	// Insert many elements
	for i := 0; i < 1000; i++ {
		sl.Insert(i)
	}

	if sl.Len() != 1000 {
		t.Errorf("Expected length 1000 after many inserts, but got %d", sl.Len())
	}

	// Search many elements
	for i := 0; i < 1000; i++ {
		if sl.Search(i) != i {
			t.Errorf("Expected to find %d, but got %v", i, sl.Search(i))
		}
	}

	// Delete many elements
	for i := 0; i < 1000; i += 2 {
		sl.Delete(i)
	}

	if sl.Len() != 500 {
		t.Errorf("Expected length 500 after many deletes, but got %d", sl.Len())
	}

	// Search remaining elements
	for i := 1; i < 1000; i += 2 {
		if sl.Search(i) != i {
			t.Errorf("Expected to find %d, but got %v", i, sl.Search(i))
		}
	}

	// Search deleted elements
	for i := 0; i < 1000; i += 2 {
		if sl.Search(i) != nil {
			t.Errorf("Expected not to find %d after deletion, but got %v", i, sl.Search(i))
		}
	}
}
