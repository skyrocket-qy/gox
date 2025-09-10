package main

func LongestCommonSubsequence(text1 string, text2 string) int {
	//makesure text2 is smaller than text1
	if len(text2) > len(text1) {
		text1, text2 = text2, text1
	}
	m, n := len(text1), len(text2)

	arr1, arr2 := make([]int, n+1, n+1), make([]int, n+1, n+1)

	for i := 0; i < m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i] == text2[j-1] {
				arr2[j] = arr1[j-1] + 1
			} else if arr2[j-1] >= arr1[j] {
				arr2[j] = arr2[j-1]
			} else {
				arr2[j] = arr1[j]
			}
		}
		arr1, arr2 = arr2, arr1
	}
	return arr1[n]
}
