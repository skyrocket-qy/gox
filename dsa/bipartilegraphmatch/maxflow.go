package bipartilegraphmatch

type Edge struct {
	to, rev   int
	cap, flow int
}

// DinicMatching: Max Flow based matching.
// Complexity: O(E * sqrt(V))
func DinicMatching[W, J comparable](adj map[W][]J) map[W]J {
	// Map W and J to integer IDs for the flow network
	// 0: Source, 1: Sink
	// Workers: 2 ... 2+len(W)-1
	// Jobs: 2+len(W) ... 2+len(W)+len(J)-1

	uIDMap := make(map[W]int)
	vIDMap := make(map[J]int)
	revUMap := make(map[int]W)
	revVMap := make(map[int]J)

	nextID := 2 // Start after Source(0) and Sink(1)

	// Collect all unique workers and jobs and assign IDs
	// First assign IDs to all workers
	for u := range adj {
		if _, exists := uIDMap[u]; !exists {
			uIDMap[u] = nextID
			revUMap[nextID] = u
			nextID++
		}
	}
	// Then assign IDs to all jobs
	for _, neighbors := range adj {
		for _, v := range neighbors {
			if _, exists := vIDMap[v]; !exists {
				vIDMap[v] = nextID
				revVMap[nextID] = v
				nextID++
			}
		}
	}

	source := 0
	sink := 1
	n := nextID // Total number of nodes in the flow network

	// Build Flow Network
	graph := make([][]Edge, n)
	addEdge := func(from, to, cap int) {
		graph[from] = append(graph[from], Edge{to: to, rev: len(graph[to]), cap: cap, flow: 0})
		graph[to] = append(graph[to], Edge{to: from, rev: len(graph[from]) - 1, cap: 0, flow: 0})
	}

	// Edges from Source to Workers
	for _, id := range uIDMap {
		addEdge(source, id, 1)
	}

	// Edges from Jobs to Sink
	for _, id := range vIDMap {
		addEdge(id, sink, 1)
	}

	// Edges from Workers to Jobs
	for u, neighbors := range adj {
		uID := uIDMap[u]
		for _, v := range neighbors {
			vID := vIDMap[v]
			addEdge(uID, vID, 1)
		}
	}

	// Dinic's Algorithm
	level := make([]int, n)
	ptr := make([]int, n)

	bfs := func() bool {
		for i := range level {
			level[i] = -1
		}
		level[source] = 0
		q := []int{source}
		for len(q) > 0 {
			u := q[0]
			q = q[1:]
			for _, e := range graph[u] {
				if e.cap-e.flow > 0 && level[e.to] == -1 {
					level[e.to] = level[u] + 1
					q = append(q, e.to)
				}
			}
		}
		return level[sink] != -1
	}

	var dfs func(u, pushed int) int
	dfs = func(u, pushed int) int {
		if pushed == 0 || u == sink {
			return pushed
		}
		for ; ptr[u] < len(graph[u]); ptr[u]++ {
			e := &graph[u][ptr[u]]
			if level[u]+1 != level[e.to] || e.cap-e.flow == 0 {
				continue
			}
			tr := dfs(e.to, min(pushed, e.cap-e.flow))
			if tr == 0 {
				continue
			}
			e.flow += tr
			graph[e.to][e.rev].flow -= tr
			return tr
		}
		return 0
	}

	maxFlow := 0
	for bfs() {
		for i := range ptr {
			ptr[i] = 0
		}
		for {
			pushed := dfs(source, 999999999)
			if pushed == 0 {
				break
			}
			maxFlow += pushed
		}
	}

	// Extract matching
	result := make(map[W]J)
	for u, uID := range uIDMap {
		for _, e := range graph[uID] {
			// Check if this edge goes from a worker node to a job node
			// and has a flow of 1 (indicating a match).
			// e.to must be a job node ID.
			if e.flow == 1 {
				if v, ok := revVMap[e.to]; ok {
					result[u] = v
					break // One match per worker
				}
			}
		}
	}

	return result
}
