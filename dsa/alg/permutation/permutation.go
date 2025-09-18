package permutation

/*
Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
*/

// factorial calculates n! for pre-allocating memory.
// Note: This will overflow for n > 20 on a 64-bit system.
// Permutation calculations are typically not feasible for large n anyway.
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	res := 1
	for i := 2; i <= n; i++ {
		res *= i
	}
	return res
}

/*
Stage1: [1]
Stage2: [2,1], [1,2]
Stage3: Insert 3 based on stage 2.

This optimized version avoids a special case for the first element and uses
a cleaner startup, which also makes the logic more straightforward.
*/
func InsertPermutation(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	res := [][]int{{nums[0]}}

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		// Pre-allocate for the next set of permutations.
		res2 := make([][]int, 0, len(res)*(i+1))
		for _, cur := range res {
			for j := 0; j <= len(cur); j++ {
				// Build the new permutation.
				now := make([]int, 0, len(cur)+1)
				now = append(now, cur[:j]...)
				now = append(now, num)
				now = append(now, cur[j:]...)
				res2 = append(res2, now)
			}
		}
		res = res2
	}
	return res
}

/*
Stage1: [1]          [2]           [3]
Stage2: [12][13]     [21][23]      [31][32]
Stage3: [123][132]   [213][231]    [312][321].

This version is optimized by:
1. Fixing a bug where result slices were mutated after being added.
2. Using a boolean `used` slice for O(1) tracking of elements, instead of O(n) `slices.Contains`.
3. Pre-allocating the result slice to avoid re-allocations.
*/
func BackTrackPermutation(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	res := make([][]int, 0, factorial(len(nums)))
	path := make([]int, 0, len(nums))
	used := make([]bool, len(nums))

	var backTrack func()
	backTrack = func() {
		if len(path) == len(nums) {
			res = append(res, append([]int(nil), path...))
			return
		}

		for i, num := range nums {
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path, num)
			backTrack()
			path = path[:len(path)-1]
			used[i] = false
		}
	}

	backTrack()
	return res
}

/*
Create a function permute() with parameters as input string, starting index of the string, ending index of the string
Call this function with values input string, 0, size of string – 1
In this function, if the value of  L and R is the same then print the same string
Else run a for loop from L to R and swap the current element in the for loop with the inputString[L]
Then again call this same function by increasing the value of L by 1
After that again swap the previously swapped values to initiate backtracking.

This version is optimized by:
1. Pre-allocating the result slice.
2. Working on a copy of the input to avoid side effects.
3. Using a cleaner closure-based implementation.
*/
func SwapPermutation(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	res := make([][]int, 0, factorial(len(nums)))
	p := append([]int(nil), nums...)

	var upset func(begin int)
	upset = func(begin int) {
		if begin == len(p) {
			res = append(res, append([]int(nil), p...))
			return
		}

		for i := begin; i < len(p); i++ {
			p[begin], p[i] = p[i], p[begin]
			upset(begin + 1)
			p[begin], p[i] = p[i], p[begin]
		}
	}

	upset(0)
	return res
}

// HeapPermutation generates all permutations of the given slice using Heap's algorithm.
// This version is optimized by pre-allocating the result slice, working on a copy of
// the input, and using a cleaner closure-based implementation.
func HeapPermutation(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	res := make([][]int, 0, factorial(len(nums)))
	p := append([]int(nil), nums...)

	var generate func(k int)
	generate = func(k int) {
		if k == 1 {
			res = append(res, append([]int(nil), p...))
			return
		}

		generate(k - 1)

		for i := 0; i < k-1; i++ {
			if k%2 == 0 {
				p[i], p[k-1] = p[k-1], p[i]
			} else {
				p[0], p[k-1] = p[k-1], p[0]
			}
			generate(k - 1)
		}
	}

	generate(len(p))
	return res
}