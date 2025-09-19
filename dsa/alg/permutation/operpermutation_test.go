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

func TestOperInsertPermutation(t *testing.T) {
	tests := []struct {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualPerms := [][]int{}
			oper := func(p []int) {
				temp := make([]int, len(p))
				copy(temp, p)
				actualPerms = append(actualPerms, temp)
			}

			permutation.OperInsertPermutation[int](tt.input, oper)

			sortPermutations(actualPerms)
			sortPermutations(tt.expected)

			if !reflect.DeepEqual(actualPerms, tt.expected) {
				t.Errorf(
					"OperInsertPermutation(%v) got %v, want %v",
					tt.input,
					actualPerms,
					tt.expected,
				)
			}
		})
	}

	// Test with string type
	stringTests := []struct {
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

	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			actualPerms := [][]string{}
			oper := func(p []string) {
				temp := make([]string, len(p))
				copy(temp, p)
				actualPerms = append(actualPerms, temp)
			}

			permutation.OperInsertPermutation[string](tt.input, oper)

			sortPermutations(actualPerms)
			sortPermutations(tt.expected)

			if !reflect.DeepEqual(actualPerms, tt.expected) {
				t.Errorf(
					"OperInsertPermutation(%v) got %v, want %v",
					tt.input,
					actualPerms,
					tt.expected,
				)
			}
		})
	}
}

func TestOperBackTrackPermutation(t *testing.T) {
	tests := []struct {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualPerms := [][]int{}
			oper := func(p []int) {
				temp := make([]int, len(p))
				copy(temp, p)
				actualPerms = append(actualPerms, temp)
			}

			permutation.OperBackTrackPermutation[int](tt.input, oper)

			sortPermutations(actualPerms)
			sortPermutations(tt.expected)

			if !reflect.DeepEqual(actualPerms, tt.expected) {
				t.Errorf(
					"OperBackTrackPermutation(%v) got %v, want %v",
					tt.input,
					actualPerms,
					tt.expected,
				)
			}
		})
	}

	// Test with string type
	stringTests := []struct {
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

	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			actualPerms := [][]string{}
			oper := func(p []string) {
				temp := make([]string, len(p))
				copy(temp, p)
				actualPerms = append(actualPerms, temp)
			}

			permutation.OperBackTrackPermutation[string](tt.input, oper)

			sortPermutations(actualPerms)
			sortPermutations(tt.expected)

			if !reflect.DeepEqual(actualPerms, tt.expected) {
				t.Errorf(
					"OperBackTrackPermutation(%v) got %v, want %v",
					tt.input,
					actualPerms,
					tt.expected,
				)
			}
		})
	}
}

func TestOperSwapPermutation(t *testing.T) {
	tests := []struct {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualPerms := [][]int{}
			oper := func(p []int) {
				temp := make([]int, len(p))
				copy(temp, p)
				actualPerms = append(actualPerms, temp)
			}

			permutation.OperSwapPermutation[int](tt.input, oper)

			sortPermutations(actualPerms)
			sortPermutations(tt.expected)

			if !reflect.DeepEqual(actualPerms, tt.expected) {
				t.Errorf(
					"OperSwapPermutation(%v) got %v, want %v",
					tt.input,
					actualPerms,
					tt.expected,
				)
			}
		})
	}

	// Test with string type
	stringTests := []struct {
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

	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			actualPerms := [][]string{}
			oper := func(p []string) {
				temp := make([]string, len(p))
				copy(temp, p)
				actualPerms = append(actualPerms, temp)
			}

			permutation.OperSwapPermutation[string](tt.input, oper)

			sortPermutations(actualPerms)
			sortPermutations(tt.expected)

			if !reflect.DeepEqual(actualPerms, tt.expected) {
				t.Errorf(
					"OperSwapPermutation(%v) got %v, want %v",
					tt.input,
					actualPerms,
					tt.expected,
				)
			}
		})
	}
}

func TestOperHeapPermutation(t *testing.T) {
	tests := []struct {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualPerms := [][]int{}
			oper := func(p []int) {
				temp := make([]int, len(p))
				copy(temp, p)
				actualPerms = append(actualPerms, temp)
			}

			permutation.OperHeapPermutation[int](tt.input, oper)

			sortPermutations(actualPerms)
			sortPermutations(tt.expected)

			if !reflect.DeepEqual(actualPerms, tt.expected) {
				t.Errorf(
					"OperHeapPermutation(%v) got %v, want %v",
					tt.input,
					actualPerms,
					tt.expected,
				)
			}
		})
	}

	// Test with string type
	stringTests := []struct {
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

	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			actualPerms := [][]string{}
			oper := func(p []string) {
				temp := make([]string, len(p))
				copy(temp, p)
				actualPerms = append(actualPerms, temp)
			}

			permutation.OperHeapPermutation[string](tt.input, oper)

			sortPermutations(actualPerms)
			sortPermutations(tt.expected)

			if !reflect.DeepEqual(actualPerms, tt.expected) {
				t.Errorf(
					"OperHeapPermutation(%v) got %v, want %v",
					tt.input,
					actualPerms,
					tt.expected,
				)
			}
		})
	}
}
