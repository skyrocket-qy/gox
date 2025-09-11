package main

func LongestCommonSubsequence(text1, text2 string) int {
	m, n := len(text1), len(text2)
	// 2-d array
	arr := make([][]int, m)
	for i := range arr {
		arr[i] = make([]int, n)
	}

	for i := range m {
		for j := range n {
			switch text1[i] {
			case text2[j]:
				prevVal := 0
				if i > 0 && j > 0 {
					prevVal = arr[i-1][j-1]
				}

				arr[i][j] = prevVal + 1
			default: // text1[i] != text2[j]
				val1 := 0
				if i > 0 {
					val1 = arr[i-1][j]
				}

				val2 := 0
				if j > 0 {
					val2 = arr[i][j-1]
				}

				arr[i][j] = max(val1, val2)
			}
		}
	}

	return arr[m-1][n-1]
}
