package main

func LongestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	//2-d array
	arr := make([][]int, m)
	for i, _ := range arr {
		arr[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if text1[i] == text2[j] {
				if i > 0 && j > 0 {
					arr[i][j] = arr[i-1][j-1] + 1
				} else {
					arr[i][j] = 1
				}
			} else if i > 0 {
				if j > 0 {
					if arr[i-1][j] >= arr[i][j-1] {
						arr[i][j] = arr[i-1][j]
					} else {
						arr[i][j] = arr[i][j-1]
					}
				} else {
					arr[i][j] = arr[i-1][j]
				}
			} else if j > 0 {
				arr[i][j] = arr[i][j-1]
			}
		}
	}
	return arr[m-1][n-1]
}
