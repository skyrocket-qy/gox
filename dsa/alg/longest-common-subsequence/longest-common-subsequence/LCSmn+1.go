package main

func LongestCommonSubsequence2(text1, text2 string) int {
	m, n := len(text1), len(text2)
	// 2-d array
	arr := make([][]int, m+1)
	for i := range arr {
		arr[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			switch {
			case text1[i-1] == text2[j-1]:
				arr[i][j] = arr[i-1][j-1] + 1
			case arr[i-1][j] >= arr[i][j-1]:
				arr[i][j] = arr[i-1][j]
			default:
				arr[i][j] = arr[i][j-1]
			}
		}
	}

	return arr[m][n]
}
