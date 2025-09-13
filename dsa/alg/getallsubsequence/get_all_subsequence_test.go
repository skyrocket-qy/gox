package getallsubsequence

import (
	"reflect"
	"testing"
)

func TestGetSubsequences(t *testing.T) {
	nums := []int{1, 2, 3}
	k := 2
	expected := [][]int{
		{1, 2},
		{1, 3},
		{2, 3},
	}
	result := GetSubsequences(nums, k)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GetSubsequences() = %v; want %v", result, expected)
	}
}

func TestGetSubsequencesIndex(t *testing.T) {
	n := 4
	k := 2
	expected := [][]int{
		{0, 1},
		{0, 2},
		{0, 3},
		{1, 2},
		{1, 3},
		{2, 3},
	}
	result := GetSubsequencesIndex(n, k)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GetSubsequencesIndex() = %v; want %v", result, expected)
	}
}
