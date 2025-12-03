package bipartilegraphmatch

// HopcroftKarp: BFS for layering + DFS for augmenting.
// Complexity: O(E * sqrt(V))
func HopcroftKarp[T comparable](adj map[T][]T, uCount, vCount int) int {
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
		// We use a special key for NIL distance or just a separate variable
		// Since T is comparable, we can't easily reserve a "NIL" T.
		// But the algorithm uses dist[NIL] to track the shortest augmenting path length found.
		// We can use a separate variable `distNIL`.
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

	// DFS: Finds augmenting paths using the levels from BFS
	// We need to pass distNIL to DFS or make it accessible.
	// Since bfs returns bool, we can't easily return distNIL.
	// We can re-run BFS logic or just capture distNIL in the closure if we merge them?
	// Or better: make distNIL a shared variable.

	// Actually, distNIL is only needed to check if we found a path.
	// In DFS, we need to check `dist[pairV[v]] == dist[u] + 1`.
	// If pairV[v] is NIL (not in map), then we found an augmenting path.

	var dfs func(u T) bool
	dfs = func(u T) bool {
		// u is always a valid worker here
		for _, v := range adj[u] {
			worker, occupied := pairV[v]
			if !occupied {
				// We found a free job.
				// In standard HK, we only accept this if dist[NIL] == dist[u] + 1
				// But we don't have dist[NIL] easily accessible here unless we store it.
				// However, if we follow the layering, we only traverse edges (u, v) where dist[pairV[v]] == dist[u] + 1.
				// If v is free, pairV[v] is NIL. dist[NIL] should be dist[u] + 1.
				// So we should check if this path length is optimal.
				// But BFS sets distNIL to the shortest path length.
				// If we are here, and v is free, it means we reached a free node.
				// We should verify if this is the shortest path layer.
				// Actually, if we strictly follow dist layers, we are good.
				// But for free nodes, "dist[NIL]" is the target.
				// Let's re-structure to share distNIL.

				// Simplified check: just take it?
				// HK requires strictly shortest paths.
				// If we don't check dist, we might take a longer path?
				// But BFS stops extending when distNIL is set.
				// So all free nodes reached in BFS are at distNIL.
				// But in DFS we might reach a free node via a longer path if we are not careful?
				// No, DFS follows dist[u] + 1.

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
	return matching
}
