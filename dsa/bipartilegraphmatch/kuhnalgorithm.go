package bipartilegraphmatch

// KuhnsAlgorithm: DFS based matching.
// Complexity: O(E * V)
func KuhnsAlgorithm[W, J comparable](adj map[W][]J) map[W]J {
	matchR := make(map[J]W) // Job -> Worker
	// visited is declared and reset inside the loop for each DFS run.

	// Let's rewrite the body to be correct and generic

	// matchR stores the worker assigned to each job
	// result stores the job assigned to each worker (for return)

	var tryKuhn func(u W, visited map[W]bool) bool
	tryKuhn = func(u W, visited map[W]bool) bool {
		if visited[u] {
			return false
		}
		visited[u] = true

		for _, v := range adj[u] {
			worker, occupied := matchR[v]
			if !occupied || tryKuhn(worker, visited) {
				matchR[v] = u
				return true
			}
		}
		return false
	}

	for u := range adj {
		visited := make(map[W]bool) // Reset visited for each worker
		tryKuhn(u, visited)
	}

	// Convert result
	result := make(map[W]J)
	for v, u := range matchR {
		result[u] = v
	}
	return result
}
