package permutation

import (
	"math/rand"
	"testing"
)

const N = 100

var BaseNums []int

func init() {
	BaseNums = make([]int, N)
	for i := 0; i < N; i++ {
		BaseNums[i] = i
	}
	rand.Shuffle(len(BaseNums), func(i, j int) {
		BaseNums[i], BaseNums[j] = BaseNums[j], BaseNums[i]
	})
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
