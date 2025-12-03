package bipartilegraphmatch

// HopcroftKarp: BFS for layering + DFS for augmenting.
// Complexity: O(E * sqrt(V))
func HopcroftKarp[T comparable](adj map[T][]T, uCount, vCount int) map[T]T {
	pairU := make(map[T]T)  // Worker -> Job
	pairV := make(map[T]T)  // Job -> Worker
	dist := make(map[T]int) // Distances for BFS layering

	// BFS: Builds level graph, returns true if an augmenting path exists
	bfs := func() bool {
		q := []T{}
		for u := range adj {
			if _, ok := pairU[u]; !ok { // If worker is free, add to queue
				dist[u] = 0
				q = append(q, u)
			} else {
				dist[u] = 999999999 // Infinite
			}
		}

		distNIL := 999999999

		for len(q) > 0 {
			u := q[0]
			q = q[1:]

			if dist[u] < distNIL {
				for _, v := range adj[u] {
					if worker, ok := pairV[v]; !ok {
						if distNIL == 999999999 {
							distNIL = dist[u] + 1
						}
					} else {
						if d, exists := dist[worker]; !exists || d == 999999999 {
							dist[worker] = dist[u] + 1
							q = append(q, worker)
						}
					}
				}
			}
		}
		return distNIL != 999999999
	}

	var dfs func(u T) bool
	dfs = func(u T) bool {
		// u is always a valid worker here
		for _, v := range adj[u] {
			worker, occupied := pairV[v]
			if !occupied {
				pairV[v] = u
				pairU[u] = v
				return true
			} else {
				if dist[worker] == dist[u]+1 {
					if dfs(worker) {
						pairV[v] = u
						pairU[u] = v
						return true
					}
				}
			}
		}
		dist[u] = 999999999 // Mark as visited/useless for this phase
		return false
	}

	matching := 0
	for bfs() {
		for u := range adj {
			if _, ok := pairU[u]; !ok {
				if dfs(u) {
					matching++
				}
			}
		}
	}
	return pairU
}
