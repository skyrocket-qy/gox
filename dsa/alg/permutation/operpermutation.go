package permutation

/*
Stage1: [1]
Stage2: [2,1], [1,2]
Stage3: Insert 3 based on stage 2.

This is a recursive implementation of the insertion algorithm. It calls the `oper`
function for each permutation found, avoiding the need to store all permutations in memory.
*/
func OperInsertPermutation[T any](nums []T, oper func([]T)) {
	if len(nums) == 0 {
		return
	}

	var insert func(currentPerm []T, remainingNums []T)
	insert = func(currentPerm []T, remainingNums []T) {
		if len(remainingNums) == 0 {
			oper(currentPerm)
			return
		}

		num := remainingNums[0]
		rest := remainingNums[1:]
		for i := 0; i <= len(currentPerm); i++ {
			newPerm := make([]T, 0, len(currentPerm)+1)
			newPerm = append(newPerm, currentPerm[:i]...)
			newPerm = append(newPerm, num)
			newPerm = append(newPerm, currentPerm[i:]...)
			insert(newPerm, rest)
		}
	}

	insert([]T{}, nums)
}

/*
Stage1: [1]          [2]           [3]
Stage2: [12][13]     [21][23]      [31][32]
Stage3: [123][132]   [213][231]    [312][321].

This version uses a callback `oper` to process each permutation, avoiding storing them.
It uses a boolean `used` slice for O(1) tracking of elements.
*/
func OperBackTrackPermutation[T any](nums []T, oper func([]T)) {
	if len(nums) == 0 {
		return
	}

	path := make([]T, 0, len(nums))
	used := make([]bool, len(nums))

	var backTrack func()
	backTrack = func() {
		if len(path) == len(nums) {
			oper(path)
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
}

/*
Create a function permute() with parameters as input string, starting index of the string, ending index of the string
Call this function with values input string, 0, size of string – 1
In this function, if the value of  L and R is the same then print the same string
Else run a for loop from L to R and swap the current element in the for loop with the inputString[L]
Then again call this same function by increasing the value of L by 1
After that again swap the previously swapped values to initiate backtracking.

This version uses a callback `oper` to process each permutation. It works on a copy
of the input to avoid side effects.
*/
func OperSwapPermutation[T any](nums []T, oper func([]T)) {
	if len(nums) == 0 {
		return
	}

	p := append([]T(nil), nums...)

	var upset func(begin int)
	upset = func(begin int) {
		if begin == len(p) {
			oper(p)
			return
		}

		for i := begin; i < len(p); i++ {
			p[begin], p[i] = p[i], p[begin]
			upset(begin + 1)
			p[begin], p[i] = p[i], p[begin]
		}
	}

	upset(0)
}

// HeapPermutation generates all permutations of the given slice using Heap's algorithm.
// This version uses a callback `oper` to process each permutation and works on a copy
// of the input to avoid side effects.
func OperHeapPermutation[T any](nums []T, oper func([]T)) {
	if len(nums) == 0 {
		return
	}

	p := append([]T(nil), nums...)

	var generate func(k int)
	generate = func(k int) {
		if k == 1 {
			oper(p)
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
}
