package bipartilegraphmatch

// HopcroftKarp: Finds Quickest Paths (BFS) + Adds Matches (DFS).
// Complexity: O(E * sqrt(V))
const Infinity = 1000000000

func HopcroftKarp[W, J comparable](adj map[W][]J) map[W]J {
	pairU := make(map[W]J)  // Worker -> Job
	pairV := make(map[J]W)  // Job -> Worker
	dist := make(map[W]int) // Distances for Quickest Paths

	// BFS: Finds Quickest Paths, returns true if an improvement path exists
	bfs := func() bool {
		q := []W{}
		for u := range adj {
			if _, ok := pairU[u]; !ok { // If worker is Available, add to queue
				dist[u] = 0
				q = append(q, u)
			} else {
				dist[u] = Infinity // Worker is Taken
			}
		}

		distNIL := Infinity

		for len(q) > 0 {
			u := q[0]
			q = q[1:]

			if dist[u] < distNIL { // worker u is available
				for _, v := range adj[u] {
					if worker, ok := pairV[v]; !ok { // job v is available
						if distNIL == Infinity { // first available job
							distNIL = dist[u] + 1 // set distance to 1
						}
					} else { // job v is taken
						if d, exists := dist[worker]; !exists || d == Infinity { // worker is available
							dist[worker] = dist[u] + 1
							q = append(q, worker)
						}
					}
				}
			}
		}
		return distNIL != Infinity // found an improvement path
	}

	// DFS: Finds Improvement Paths using the distances from BFS
	var dfs func(u W) bool
	dfs = func(u W) bool {
		// u is always a valid worker here
		for _, v := range adj[u] {
			worker, occupied := pairV[v]
			if !occupied {
				pairV[v] = u // job v is now taken by worker u
				pairU[u] = v // worker u is now matched to job v
				return true  // found an improvement path
			} else {
				if dist[worker] == dist[u]+1 { // can swap
					if dfs(worker) { // found an improvement path
						pairV[v] = u
						pairU[u] = v
						return true
					}
				}
			}
		}
		dist[u] = Infinity // Mark as visited/no path for this round
		return false
	}

	for bfs() {
		for u := range adj {
			if _, ok := pairU[u]; !ok {
				dfs(u)
			}
		}
	}
	return pairU
}
