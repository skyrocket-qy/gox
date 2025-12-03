package bipartilegraphmatch

// BacktrackingMatching: Pure recursive trial and error.
// Complexity: Exponential O(2^E) in worst case without memoization.
func BacktrackingMatching(adj [][]int, uCount, vCount int) int {
	matchR := make([]int, vCount)
	for i := range matchR {
		matchR[i] = -1 // -1 means Job is free
	}

	var maxMatches int

	// recursive function to process worker 'u'
	var solve func(u int, currentMatches int)
	solve = func(u int, currentMatches int) {
		if u == uCount {
			if currentMatches > maxMatches {
				maxMatches = currentMatches
			}
			return
		}

		// Option 1: Don't match this worker
		solve(u+1, currentMatches)

		// Option 2: Try to match with all neighbors
		for _, v := range adj[u] {
			if matchR[v] == -1 { // If job is available
				matchR[v] = u // Assign
				solve(u+1, currentMatches+1)
				matchR[v] = -1 // Backtrack
			}
		}
	}

	solve(0, 0)
	return maxMatches
}
