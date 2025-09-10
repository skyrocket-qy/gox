package longestcommonsubsequence

/* @tags: LCS,dynamic programming */

func longestCommonSubsequence(word1, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	if word1[0] == word2[0] {
		dp[0][0] = 1
	}

	for i := 1; i < m; i++ {
		if word1[i] == word2[0] || dp[i-1][0] == 1 {
			dp[i][0] = 1
		}
	}

	for j := 1; j < n; j++ {
		if word1[0] == word2[j] || dp[0][j-1] == 1 {
			dp[0][j] = 1
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if word1[i] == word2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m-1][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
