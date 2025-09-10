package getallsubsequence

/* @tags: subsequence,permutation */

func GetSubsequences(nums []int, k int) [][]int {
	result := [][]int{}
	GenerateSubsequences(nums, k, 0, []int{}, &result)
	return result
}

func GenerateSubsequences(nums []int, k int, start int, current []int, result *[][]int) {
	if len(current) == k {
		temp := make([]int, len(current))
		copy(temp, current)
		*result = append(*result, temp)
		return
	}

	for i := start; i < len(nums); i++ {
		current = append(current, nums[i])
		GenerateSubsequences(nums, k, i+1, current, result)
		current = current[:len(current)-1]
	}
}

func GetSubsequencesIndex(n, k int) [][]int {
	result := [][]int{}
	GenerateSubsequencesIndex(n, k, 0, []int{}, &result)
	return result
}

func GenerateSubsequencesIndex(n int, k int, start int, current []int, result *[][]int) {
	if len(current) == k {
		temp := make([]int, len(current))
		copy(temp, current)
		*result = append(*result, temp)
		return
	}

	for i := start; i < n; i++ {
		current = append(current, i)
		GenerateSubsequencesIndex(n, k, i+1, current, result)
		current = current[:len(current)-1]
	}
}
