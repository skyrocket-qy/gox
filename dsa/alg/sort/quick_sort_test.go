package sort_test

import (
	"reflect"
	"sort"
	"testing"

	sortx "github.com/skyrocket-qy/gox/dsa/alg/sort"
)

// The original qSortOutMem has a bug. The pivot is not excluded from the partitioning,
// which causes it to be included in the `r` slice and then added again.
// This test will fail on most inputs.
func TestQSortOutMem(t *testing.T) {
	testCases := [][]int{
		{},
		{1},
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
		{1, 5, 2, 4, 3},
		{1, 1, 1, 1, 1},
	}

	for _, tc := range testCases {
		original := make([]int, len(tc))
		copy(original, tc)
		sorted := sortx.QSortOutMem(tc)

		sort.Ints(original)

		if !reflect.DeepEqual(sorted, original) {
			t.Errorf("qSortOutMem(%v) = %v, want %v", tc, sorted, original)
		}
	}
}

// The original qSortInMem has multiple bugs. The partitioning logic is flawed
// and the final swap is incorrect. This test will fail on most inputs.
func TestQSortInMem(t *testing.T) {
	testCases := [][]int{
		{},
		{1},
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
		{1, 5, 2, 4, 3},
		{1, 1, 1, 1, 1},
	}

	for _, tc := range testCases {
		original := make([]int, len(tc))
		copy(original, tc)
		sortx.QSortInMem(tc, 0, len(tc)-1)
		sort.Ints(original)

		if !reflect.DeepEqual(tc, original) {
			t.Errorf("qSortInMem(%v) = %v, want %v", original, tc, original)
		}
	}
}

// The qSortInPartition function is implemented correctly.
func TestQSortInPartition(t *testing.T) {
	testCases := [][]int{
		{},
		{1},
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
		{1, 5, 2, 4, 3},
		{1, 1, 1, 1, 1},
	}

	for _, tc := range testCases {
		original := make([]int, len(tc))
		copy(original, tc)
		sortx.QSortInPartition(tc, 0, len(tc)-1)
		sort.Ints(original)

		if !reflect.DeepEqual(tc, original) {
			t.Errorf("qSortInPartition(%v) = %v, want %v", original, tc, original)
		}
	}
}
