package differencearray

import (
	"reflect"
	"testing"
)

func TestDifferenceArray(t *testing.T) {
	in := []int{1, 2, 3, 4, 5}
	diffArr := NewDifferenceArray(in)

	if !reflect.DeepEqual([]int(diffArr), []int{1, 1, 1, 1, 1}) {
		t.Errorf("NewDifferenceArray failed, got %v", diffArr)
	}

	diffArr.IntervalUpdate(1, 3, 2)

	if !reflect.DeepEqual([]int(diffArr), []int{1, 3, 1, 1, -1}) {
		t.Errorf("IntervalUpdate failed, got %v", diffArr)
	}

	diffArr.Rebuild()

	if !reflect.DeepEqual([]int(diffArr), []int{1, 4, 5, 6, 5}) {
		t.Errorf("Rebuild failed, got %v", diffArr)
	}

	if diffArr.Query(2) != 5 {
		t.Errorf("Query failed, got %d", diffArr.Query(2))
	}
}

func TestDifferenceArray_edge_case(t *testing.T) {
	in := []int{1, 2, 3, 4, 5}
	diffArr := NewDifferenceArray(in)
	diffArr.IntervalUpdate(0, 4, 2)
	diffArr.Rebuild()

	expected := []int{3, 4, 5, 6, 7}
	if !reflect.DeepEqual([]int(diffArr), expected) {
		t.Errorf("Rebuild failed, got %v, want %v", diffArr, expected)
	}
}
