package bipartilegraphmatch

// KuhnsAlgorithm: DFS based augmenting paths.
// Complexity: O(V * E)
func KuhnsAlgorithm(adj [][]int, uCount, vCount int) int {
	matchR := make([]int, vCount) // Stores which worker is assigned to job 'v'
	for i := range matchR {
		matchR[i] = -1
	}

	var visited []bool

	// DFS to find augmenting path
	var dfs func(u int) bool
	dfs = func(u int) bool {
		for _, v := range adj[u] {
			if !visited[v] {
				visited[v] = true
				// If job 'v' is free OR the worker currently holding 'v' can find another job
				if matchR[v] < 0 || dfs(matchR[v]) {
					matchR[v] = u
					return true
				}
			}
		}
		return false
	}

	result := 0
	for u := 0; u < uCount; u++ {
		visited = make([]bool, vCount) // Reset visited for every worker
		if dfs(u) {
			result++
		}
	}
	return result
}
