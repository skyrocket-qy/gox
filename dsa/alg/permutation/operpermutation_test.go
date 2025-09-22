package permutation_test

import (
	"cmp"
	"reflect"
	"sort"
	"testing"

	"github.com/skyrocket-qy/gox/dsa/alg/permutation"
)

// Helper function to sort a slice of integer slices for consistent comparison.
func sortPermutations[T cmp.Ordered](perms [][]T) {
	sort.Slice(perms, func(i, j int) bool {
		for k := 0; k < len(perms[i]) && k < len(perms[j]); k++ {
			if perms[i][k] != perms[j][k] {
				return perms[i][k] < perms[j][k]
			}
		}

		return len(perms[i]) < len(perms[j])
	})
}

// runPermutationTest is a generic helper function to test permutation functions.
func runPermutationTest[T cmp.Ordered](
	t *testing.T,
	permutationFunc func([]T, func([]T)),
	input []T,
	expected [][]T,
	funcName string,
) {
	t.Helper()

	actualPerms := [][]T{}
	oper := func(p []T) {
		temp := make([]T, len(p))
		copy(temp, p)
		actualPerms = append(actualPerms, temp)
	}

	permutationFunc(input, oper)

	sortPermutations(actualPerms)
	sortPermutations(expected)

	if !reflect.DeepEqual(actualPerms, expected) {
		t.Errorf(
			"%s(%v) got %v, want %v",
			funcName,
			input,
			actualPerms,
			expected,
		)
	}
}

var intTests = []struct {
	name     string
	input    []int
	expected [][]int
}{
	{
		name:     "Empty slice",
		input:    []int{},
		expected: [][]int{},
	},
	{
		name:     "Single element",
		input:    []int{1},
		expected: [][]int{{1}},
	},
	{
		name:     "Two elements",
		input:    []int{1, 2},
		expected: [][]int{{1, 2}, {2, 1}},
	},
	{
		name:     "Three elements",
		input:    []int{1, 2, 3},
		expected: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
	},
}

var stringTests = []struct {
	name     string
	input    []string
	expected [][]string
}{
	{
		name:     "String - Two elements",
		input:    []string{"a", "b"},
		expected: [][]string{{"a", "b"}, {"b", "a"}},
	},
	{
		name:  "String - Three elements",
		input: []string{"a", "b", "c"},
		expected: [][]string{
			{"a", "b", "c"},
			{"a", "c", "b"},
			{"b", "a", "c"},
			{"b", "c", "a"},
			{"c", "a", "b"},
			{"c", "b", "a"},
		},
	},
}

func TestOperInsertPermutation(t *testing.T) {
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			runPermutationTest(
				t,
				permutation.OperInsertPermutation[int],
				tt.input,
				tt.expected,
				"OperInsertPermutation",
			)
		})
	}

	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			runPermutationTest(
				t,
				permutation.OperInsertPermutation[string],
				tt.input,
				tt.expected,
				"OperInsertPermutation",
			)
		})
	}
}

func TestOperBackTrackPermutation(t *testing.T) {
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			runPermutationTest(
				t,
				permutation.OperBackTrackPermutation[int],
				tt.input,
				tt.expected,
				"OperBackTrackPermutation",
			)
		})
	}

	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			runPermutationTest(
				t,
				permutation.OperBackTrackPermutation[string],
				tt.input,
				tt.expected,
				"OperBackTrackPermutation",
			)
		})
	}
}

func TestOperSwapPermutation(t *testing.T) {
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			runPermutationTest(
				t,
				permutation.OperSwapPermutation[int],
				tt.input,
				tt.expected,
				"OperSwapPermutation",
			)
		})
	}

	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			runPermutationTest(
				t,
				permutation.OperSwapPermutation[string],
				tt.input,
				tt.expected,
				"OperSwapPermutation",
			)
		})
	}
}

func TestOperHeapPermutation(t *testing.T) {
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			runPermutationTest(
				t,
				permutation.OperHeapPermutation[int],
				tt.input,
				tt.expected,
				"OperHeapPermutation",
			)
		})
	}

	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			runPermutationTest(
				t,
				permutation.OperHeapPermutation[string],
				tt.input,
				tt.expected,
				"OperHeapPermutation",
			)
		})
	}
}
