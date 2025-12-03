package bipartilegraphmatch

import (
	"fmt"
)

// HopcroftKarpWithLogging: Same as HopcroftKarp but prints execution steps.
func HopcroftKarpWithLogging[W, J comparable](adj map[W][]J) map[W]J {
	pairU := make(map[W]J)  // Worker -> Job
	pairV := make(map[J]W)  // Job -> Worker
	dist := make(map[W]int) // Distances for Quickest Paths

	fmt.Println("Starting Hopcroft-Karp Algorithm...")

	round := 1

	// BFS: Finds Quickest Paths, returns true if an improvement path exists
	bfs := func() bool {
		fmt.Printf("\n--- Round %d: Find Quickest Paths (BFS) ---\n", round)
		q := []W{}
		for u := range adj {
			if _, ok := pairU[u]; !ok { // If worker is Available, add to queue
				dist[u] = 0
				q = append(q, u)
				fmt.Printf("  [BFS Init] Worker %v is Available -> Add to Q (Dist: 0)\n", u)
			} else {
				dist[u] = Infinity // Infinite (Worker is Taken)
			}
		}
		fmt.Printf("  [BFS State] Initial Distances: %v\n", dist)

		distNIL := Infinity
		fmt.Println("  [BFS Loop] Starting Queue processing...")

		for len(q) > 0 {
			u := q[0]
			q = q[1:]
			fmt.Printf("  [BFS Dequeue] Processing Worker %v (Dist: %d)\n", u, dist[u])

			if dist[u] < distNIL { // worker u is available
				for _, v := range adj[u] {
					fmt.Printf("    -> Checking Neighbor Job %v... ", v)
					if worker, ok := pairV[v]; !ok { // job v is available
						// v is free
						fmt.Printf("Available! ")
						if distNIL == Infinity { // first available job
							distNIL = dist[u] + 1 // set distance to 1
							fmt.Printf("Found Quickest Path Length: %d\n", distNIL)
						} else {
							fmt.Printf("(Path Length %d already found)\n", distNIL)
						}
					} else { // job v is taken
						// v is matched to worker
						fmt.Printf("Taken by Worker %v. ", worker)
						if d, exists := dist[worker]; !exists || d == Infinity { // worker is available
							dist[worker] = dist[u] + 1
							q = append(q, worker)
							fmt.Printf("Update Dist[%v] = %d -> Add to Q\n", worker, dist[worker])
						} else {
							fmt.Printf("Dist[%v] already %d (No update)\n", worker, d)
						}
					}
				}
			} else {
				fmt.Printf("    Skipping (Dist %d >= Quickest Path Length %d)\n", dist[u], distNIL)
			}
		}
		if distNIL != Infinity {
			fmt.Printf("  [BFS Done] Quickest path length found: %d\n", distNIL)
			return true
		}
		fmt.Println("  [BFS Done] No paths found.")
		return false
	}

	// DFS: Finds Improvement Paths using the distances from BFS
	var dfs func(u W, path []W) bool
	dfs = func(u W, path []W) bool {
		indent := ""
		for i := 0; i < len(path); i++ {
			indent += "  "
		}

		path = append(path, u)
		fmt.Printf("    %s[DFS Visit] Worker %v (Path: %v)\n", indent, u, path)

		for _, v := range adj[u] {
			worker, occupied := pairV[v]
			if !occupied {
				// Found free job
				fmt.Printf("    %s-> Found Available Job %v! **SWAP** executed along path.\n", indent, v)
				pairV[v] = u
				pairU[u] = v
				fmt.Printf("    %s   (Match: %v <==> %v)\n", indent, u, v)
				return true
			} else {
				fmt.Printf("    %s-> Job %v is Taken by %v. Checking Dist...\n", indent, v, worker)
				if dist[worker] == dist[u]+1 {
					fmt.Printf("    %s   Dist[%v](%d) == Dist[%v](%d) + 1. Recursing to %v...\n", indent, worker, dist[worker], u, dist[u], worker)
					if dfs(worker, path) { // Pass current path, don't append J node
						// The path in explanation is U -> V -> U -> V.
						// But our DFS only tracks U nodes in recursion.
						// Let's just track U nodes in path for logging simplicity or convert V to string if needed.
						// Actually, the previous code appended v (type T) to path (type []T).
						// Now W and J are different. We can't append J to []W.
						// Let's just append the next worker to the path in the recursive call.
						// The path variable is just for logging.
						// We can't easily mix types in a slice in Go without interface{}.
						// Let's just log the worker path.
						pairV[v] = u
						pairU[u] = v
						fmt.Printf("    %s   (Match: %v <==> %v)\n", indent, u, v)
						return true
					}
				} else {
					fmt.Printf("    %s   Skipping (Distances don't match for shortest path)\n", indent)
				}
			}
		}
		dist[u] = Infinity // Mark as visited/useless for this phase
		fmt.Printf("    %s[DFS Backtrack] No path from %v. Mark Dist[%v] = Inf\n", indent, u, u)
		return false
	}

	matching := 0
	for bfs() {
		fmt.Printf("--- Round %d: Add Matches (DFS) ---\n", round)
		for u := range adj {
			if _, ok := pairU[u]; !ok {
				fmt.Printf("  [DFS Start] Trying to match Available Worker %v...\n", u)
				if dfs(u, []W{}) {
					matching++
					fmt.Printf("  [DFS Success] Matched Worker %v. Total Matching: %d\n", u, matching)
				} else {
					fmt.Printf("  [DFS Fail] Could not match Worker %v this round.\n", u)
				}
			}
		}
		round++
	}

	fmt.Println("\nAlgorithm Finished.")
	fmt.Printf("Final Matching Size: %d\n", matching)
	return pairU
}
