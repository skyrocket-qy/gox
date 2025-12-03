package bipartilegraphmatch

// BacktrackingMatching: Simple recursive backtracking.
// Complexity: Exponential O(2^E)
func BacktrackingMatching[W, J comparable](adj map[W][]J) map[W]J {
	// Convert map keys to a slice for deterministic iteration
	var workers []W
	for w := range adj {
		workers = append(workers, w)
	}

	bestMatching := make(map[W]J)
	currentMatching := make(map[W]J)
	usedJobs := make(map[J]bool)

	var backtrack func(index int)
	backtrack = func(index int) {
		if index == len(workers) {
			if len(currentMatching) > len(bestMatching) {
				// Copy current matching to best matching
				bestMatching = make(map[W]J)
				for k, v := range currentMatching {
					bestMatching[k] = v
				}
			}
			return
		}

		w := workers[index]

		// Option 1: Don't match this worker
		backtrack(index + 1)

		// Option 2: Try to match with each available neighbor
		for _, job := range adj[w] {
			if !usedJobs[job] {
				usedJobs[job] = true
				currentMatching[w] = job
				backtrack(index + 1)
				delete(currentMatching, w)
				usedJobs[job] = false
			}
		}
	}

	backtrack(0)
	return bestMatching
}
