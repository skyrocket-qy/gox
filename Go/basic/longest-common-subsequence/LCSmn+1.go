package main

func LongestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	//2-d array
	arr := make([][]int, m+1)
	for i, _ := range arr {
		arr[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				arr[i][j] = arr[i-1][j-1] + 1
			} else if arr[i-1][j] >= arr[i][j-1] {
				arr[i][j] = arr[i-1][j]
			} else {
				arr[i][j] = arr[i][j-1]
			}
		}
	}

	return arr[m][n]
}
