package permutation_test

import (
	"testing"

	"github.com/skyrocket-qy/gox/dsa/permutation"
)

const N = 10 // Reduced N to a reasonable size for permutation benchmarks

var BaseNums []int

func init() {
	BaseNums = make([]int, N)
	for i := range N {
		BaseNums[i] = i
	}
}

func BenchmarkInsertPermutation(b *testing.B) {
	nums := make([]int, N)
	copy(nums, BaseNums)
	b.ResetTimer()

	for range b.N {
		permutation.InsertPermutation(nums)
	}
}

func BenchmarkBackTrackPermutation(b *testing.B) {
	nums := make([]int, N)
	copy(nums, BaseNums)
	b.ResetTimer()

	for range b.N {
		permutation.BackTrackPermutation(nums)
	}
}

func BenchmarkSwapPermutation(b *testing.B) {
	nums := make([]int, N)
	copy(nums, BaseNums)
	b.ResetTimer()

	for range b.N {
		permutation.SwapPermutation(nums)
	}
}

func BenchmarkHeapPermutation(b *testing.B) {
	nums := make([]int, N)
	copy(nums, BaseNums)
	b.ResetTimer()

	for range b.N {
		permutation.HeapPermutation(nums)
	}
}
