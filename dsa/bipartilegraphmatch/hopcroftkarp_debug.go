package bipartilegraphmatch

import (
	"fmt"
)

// HopcroftKarpWithLogging: Same as HopcroftKarp but prints execution steps.
func HopcroftKarpWithLogging[T comparable](adj map[T][]T, uCount, vCount int) map[T]T {
	pairU := make(map[T]T)  // Worker -> Job
	pairV := make(map[T]T)  // Job -> Worker
	dist := make(map[T]int) // Distances for BFS layering

	fmt.Println("Starting Hopcroft-Karp Algorithm...")

	phase := 1

	// BFS: Builds level graph, returns true if an augmenting path exists
	bfs := func() bool {
		fmt.Printf("\n--- Phase %d: BFS (Layering) ---\n", phase)
		q := []T{}
		for u := range adj {
			if _, ok := pairU[u]; !ok { // If worker is free, add to queue
				dist[u] = 0
				q = append(q, u)
				fmt.Printf("  Worker %v is free, adding to queue (Dist: 0)\n", u)
			} else {
				dist[u] = 999999999 // Infinite
			}
		}
		fmt.Println("dist:", dist)

		distNIL := 999999999
		fmt.Println("  BFS Queue processing...")

		for len(q) > 0 {
			u := q[0]
			q = q[1:]

			if dist[u] < distNIL {
				for _, v := range adj[u] {
					if worker, ok := pairV[v]; !ok {
						// v is free
						if distNIL == 999999999 {
							distNIL = dist[u] + 1
							fmt.Printf("  Found free Job %v at distance %d! (Shortest Path Length set)\n", v, distNIL)
						}
					} else {
						// v is matched to worker
						if d, exists := dist[worker]; !exists || d == 999999999 {
							dist[worker] = dist[u] + 1
							q = append(q, worker)
							fmt.Printf("  Job %v is matched to Worker %v. Setting Dist[%v] = %d\n", v, worker, worker, dist[worker])
						}
					}
				}
			}
		}
		if distNIL != 999999999 {
			fmt.Printf("  BFS finished. Shortest augmenting path length: %d\n", distNIL)
			return true
		}
		fmt.Println("  BFS finished. No augmenting paths found.")
		return false
	}

	// DFS: Finds augmenting paths using the levels from BFS
	var dfs func(u T, path []T) bool
	dfs = func(u T, path []T) bool {
		path = append(path, u)
		for _, v := range adj[u] {
			worker, occupied := pairV[v]
			if !occupied {
				// Found free job
				// We rely on BFS layers implicitly or check dist if needed.
				// But standard HK logic with distNIL check in BFS ensures we only find shortest paths if we follow dist.
				// However, here we don't have distNIL easily available unless we store it or just trust the flow.
				// Let's print the augmentation.
				fmt.Printf("    DFS reached free Job %v via path %v. Augmenting!\n", v, path)
				pairV[v] = u
				pairU[u] = v
				return true
			} else {
				if dist[worker] == dist[u]+1 {
					if dfs(worker, append(path, v)) {
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
		fmt.Printf("--- Phase %d: DFS (Augmenting) ---\n", phase)
		for u := range adj {
			if _, ok := pairU[u]; !ok {
				if dfs(u, []T{}) {
					matching++
					fmt.Printf("  Matched Worker %v. Total Matching: %d\n", u, matching)
				}
			}
		}
		phase++
	}

	fmt.Println("\nAlgorithm Finished.")
	fmt.Printf("Final Matching Size: %d\n", matching)
	return pairU
}
