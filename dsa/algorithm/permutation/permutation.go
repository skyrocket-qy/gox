package permutation

/* @tags: permutation,backtracking */

/*
Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
*/

/*
Stage1: [1]
Stage2: [2,1], [1,2]
Stage3: Insert 3 based on stage 2
*/
func InsertPermutation(nums []int) [][]int {
	res := [][]int{{}}
	for _, num := range nums {
		res2 := [][]int{}
		for _, cur := range res {
			for i := 0; i <= len(cur); i++ {
				now := append([]int{}, cur[:i]...)
				now = append(now, num)
				now = append(now, cur[i:]...)
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
Stage3: [123][132]   [213][231]    [312][321]
*/
func BackTrackPermutation(nums []int) [][]int {
	var backTrack func(nums []int, curNums []int, res *[][]int)
	backTrack = func(nums []int, curNums []int, res *[][]int) {
		if len(curNums) == len(nums) {
			*res = append(*res, curNums)
			return
		}
		for _, num := range nums {
			isContained := false
			for _, v := range curNums {
				if v == num {
					isContained = true
					break
				}
			}
			if isContained {
				continue
			}
			curNums = append(curNums, num)
			backTrack(nums, curNums, res)
			curNums = curNums[:len(curNums)-1]
		}
	}

	res := [][]int{}
	backTrack(nums, []int{}, &res)
	return res
}

/*
Create a function permute() with parameters as input string, starting index of the string, ending index of the string
Call this function with values input string, 0, size of string – 1
In this function, if the value of  L and R is the same then print the same string
Else run a for loop from L to R and swap the current element in the for loop with the inputString[L]
Then again call this same function by increasing the value of L by 1
After that again swap the previously swapped values to initiate backtracking
*/
func SwapPermutation(nums []int) [][]int {
	var upset func(nums []int, begin int, res *[][]int)
	upset = func(nums []int, begin int, res *[][]int) {
		if begin == len(nums) {
			*res = append(*res, nums)
		}
		for i := begin; i < len(nums); i++ {
			nums[i], nums[begin] = nums[begin], nums[i]
			upset(nums, begin+1, res)
			nums[i], nums[begin] = nums[begin], nums[i]
		}
	}

	res := [][]int{}
	upset(nums, 0, &res)
	return res
}
