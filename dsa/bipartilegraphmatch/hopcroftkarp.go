package bipartilegraphmatch

// HopcroftKarp: BFS for layering + DFS for augmenting.
// Complexity: O(E * sqrt(V))
func HopcroftKarp(adj [][]int, uCount, vCount int) int {
	pairU := make([]int, uCount)  // Worker -> Job
	pairV := make([]int, vCount)  // Job -> Worker
	dist := make([]int, uCount+1) // Distances for BFS layering

	for i := range pairU {
		pairU[i] = -1
	}
	for i := range pairV {
		pairV[i] = -1
	}

	// BFS: Builds level graph, returns true if an augmenting path exists
	bfs := func() bool {
		q := []int{}
		for u := 0; u < uCount; u++ {
			if pairU[u] == -1 { // If worker is free, add to queue
				dist[u] = 0
				q = append(q, u)
			} else {
				dist[u] = 999999999 // Infinite
			}
		}
		dist[uCount] = 999999999 // Distance to NIL node

		for len(q) > 0 {
			u := q[0]
			q = q[1:]
			if dist[u] < dist[uCount] {
				for _, v := range adj[u] {
					if pairV[v] == -1 {
						if dist[uCount] == 999999999 {
							dist[uCount] = dist[u] + 1
						}
					} else {
						if dist[pairV[v]] == 999999999 {
							dist[pairV[v]] = dist[u] + 1
							q = append(q, pairV[v])
						}
					}
				}
			}
		}
		return dist[uCount] != 999999999
	}

	// DFS: Finds augmenting paths using the levels from BFS
	var dfs func(u int) bool
	dfs = func(u int) bool {
		if u != -1 {
			for _, v := range adj[u] {
				// Follow the BFS levels
				if pairV[v] == -1 || (dist[pairV[v]] == dist[u]+1 && dfs(pairV[v])) {
					pairV[v] = u
					pairU[u] = v
					return true
				}
			}
			dist[u] = 999999999 // Mark as visited/useless for this phase
			return false
		}
		return true
	}

	matching := 0
	for bfs() {
		for u := 0; u < uCount; u++ {
			if pairU[u] == -1 {
				if dfs(u) {
					matching++
				}
			}
		}
	}
	return matching
}
