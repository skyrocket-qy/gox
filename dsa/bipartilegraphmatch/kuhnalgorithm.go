package bipartilegraphmatch

// KuhnsAlgorithm: DFS based augmenting paths.
// Complexity: O(V * E)
// KuhnsAlgorithm: DFS based augmenting paths.
// Complexity: O(V * E)
func KuhnsAlgorithm[T comparable](adj map[T][]T, uCount, vCount int) int {
	matchR := make(map[T]T) // Stores which worker is assigned to job 'v'

	var visited map[T]bool

	// DFS to find augmenting path
	var dfs func(u T) bool
	dfs = func(u T) bool {
		for _, v := range adj[u] {
			if !visited[v] {
				visited[v] = true
				// If job 'v' is free OR the worker currently holding 'v' can find another job
				worker, occupied := matchR[v]
				if !occupied || dfs(worker) {
					matchR[v] = u
					return true
				}
			}
		}
		return false
	}

	result := 0
	for u := range adj {
		visited = make(map[T]bool) // Reset visited for every worker
		if dfs(u) {
			result++
		}
	}
	return result
}
