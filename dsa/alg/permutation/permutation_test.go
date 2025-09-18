package permutation_test

import (
	"testing"

	"github.com/skyrocket-qy/gox/dsa/alg/permutation"
	"github.com/stretchr/testify/assert"
)

func TestInsertPermutation(t *testing.T) {
	nums := []int{1, 2, 3}
	expected := [][]int{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
	}
	result := permutation.InsertPermutation(nums)
	assert.ElementsMatch(t, expected, result)

	nums = []int{0, 1}
	expected = [][]int{{0, 1}, {1, 0}}
	result = permutation.InsertPermutation(nums)
	assert.ElementsMatch(t, expected, result)

	nums = []int{1}
	expected = [][]int{{1}}
	result = permutation.InsertPermutation(nums)
	assert.ElementsMatch(t, expected, result)
}

func TestBackTrackPermutation(t *testing.T) {
	nums := []int{1, 2, 3}
	expected := [][]int{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
	}
	result := permutation.BackTrackPermutation(nums)
	assert.ElementsMatch(t, expected, result)

	nums = []int{0, 1}
	expected = [][]int{{0, 1}, {1, 0}}
	result = permutation.BackTrackPermutation(nums)
	assert.ElementsMatch(t, expected, result)

	nums = []int{1}
	expected = [][]int{{1}}
	result = permutation.BackTrackPermutation(nums)
	assert.ElementsMatch(t, expected, result)
}

func TestSwapPermutation(t *testing.T) {
	nums := []int{1, 2, 3}
	expected := [][]int{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
	}
	result := permutation.SwapPermutation(nums)
	assert.ElementsMatch(t, expected, result)

	nums = []int{0, 1}
	expected = [][]int{{0, 1}, {1, 0}}
	result = permutation.SwapPermutation(nums)
	assert.ElementsMatch(t, expected, result)

	nums = []int{1}
	expected = [][]int{{1}}
	result = permutation.SwapPermutation(nums)
	assert.ElementsMatch(t, expected, result)
}

func TestHeapPermutation(t *testing.T) {
	nums := []int{1, 2, 3}
	expected := [][]int{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
	}
	result := permutation.HeapPermutation(nums)
	assert.ElementsMatch(t, expected, result)

	nums = []int{0, 1}
	expected = [][]int{{0, 1}, {1, 0}}
	result = permutation.HeapPermutation(nums)
	assert.ElementsMatch(t, expected, result)

	nums = []int{1}
	expected = [][]int{{1}}
	result = permutation.HeapPermutation(nums)
	assert.ElementsMatch(t, expected, result)
}
