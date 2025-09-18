package permutation

import (
	"testing"
)

const N = 8 // Reduced N to a reasonable size for permutation benchmarks

var BaseNums []int

func init() {
	BaseNums = make([]int, N)
	for i := 0; i < N; i++ {
		BaseNums[i] = i
	}
}

func BenchmarkInsertPermutation(b *testing.B) {
	nums := make([]int, N)
	copy(nums, BaseNums)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		InsertPermutation(nums)
	}
}

func BenchmarkBackTrackPermutation(b *testing.B) {
	nums := make([]int, N)
	copy(nums, BaseNums)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BackTrackPermutation(nums)
	}
}

func BenchmarkSwapPermutation(b *testing.B) {
	nums := make([]int, N)
	copy(nums, BaseNums)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SwapPermutation(nums)
	}
}

func BenchmarkHeapPermutation(b *testing.B) {
	nums := make([]int, N)
	copy(nums, BaseNums)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		HeapPermutation(nums)
	}
}
