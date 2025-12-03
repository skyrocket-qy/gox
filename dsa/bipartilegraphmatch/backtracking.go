package bipartilegraphmatch

// BacktrackingMatching: Pure recursive trial and error.
// Complexity: Exponential O(2^E) in worst case without memoization.
func BacktrackingMatching[T comparable](adj map[T][]T, uCount, vCount int) map[T]T {
	matchR := make(map[T]T)    // Job -> Worker
	bestMatch := make(map[T]T) // Worker -> Job (Result)
	maxMatches := 0

	// Collect workers (keys of adj)
	var workers []T
	for u := range adj {
		workers = append(workers, u)
	}

	// recursive function to process worker at index 'idx'
	var solve func(idx int, currentMatches int)
	solve = func(idx int, currentMatches int) {
		if idx == len(workers) {
			if currentMatches > maxMatches {
				maxMatches = currentMatches
				// Save current state to bestMatch
				// matchR is Job->Worker. We want Worker->Job.
				clear(bestMatch)
				for v, u := range matchR {
					bestMatch[u] = v
				}
			}
			return
		}

		u := workers[idx]

		// Option 1: Don't match this worker
		solve(idx+1, currentMatches)

		// Option 2: Try to match with all neighbors
		for _, v := range adj[u] {
			// Check if job 'v' is free
			if _, occupied := matchR[v]; !occupied {
				matchR[v] = u // Assign
				solve(idx+1, currentMatches+1)
				delete(matchR, v) // Backtrack
			}
		}
	}

	solve(0, 0)
	return bestMatch
}
