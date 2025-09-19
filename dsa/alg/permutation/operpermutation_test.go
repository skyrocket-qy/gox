package permutation

import (
	"reflect"
	"sort"
	"testing"
)

// Helper function to sort a slice of integer slices for consistent comparison
func sortPermutations(perms [][]int) {
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
				// Make a copy to avoid issues with slice reuse
				temp := make([]int, len(p))
				copy(temp, p)
				actualPerms = append(actualPerms, temp)
			}

			OperInsertPermutation(tt.input, oper)

			sortPermutations(actualPerms)
			sortPermutations(tt.expected)

			if !reflect.DeepEqual(actualPerms, tt.expected) {
				t.Errorf("OperInsertPermutation(%v) got %v, want %v", tt.input, actualPerms, tt.expected)
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

			OperBackTrackPermutation(tt.input, oper)

			sortPermutations(actualPerms)
			sortPermutations(tt.expected)

			if !reflect.DeepEqual(actualPerms, tt.expected) {
				t.Errorf("OperBackTrackPermutation(%v) got %v, want %v", tt.input, actualPerms, tt.expected)
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

			OperSwapPermutation(tt.input, oper)

			sortPermutations(actualPerms)
			sortPermutations(tt.expected)

			if !reflect.DeepEqual(actualPerms, tt.expected) {
				t.Errorf("OperSwapPermutation(%v) got %v, want %v", tt.input, actualPerms, tt.expected)
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

			OperHeapPermutation(tt.input, oper)

			sortPermutations(actualPerms)
			sortPermutations(tt.expected)

			if !reflect.DeepEqual(actualPerms, tt.expected) {
				t.Errorf("OperHeapPermutation(%v) got %v, want %v", tt.input, actualPerms, tt.expected)
			}
		})
	}
}
