package bipartilegraphmatch

type Edge struct {
	to, rev   int
	cap, flow int
}

// DinicMaxFlow: Solves using Flow Network construction.
// Complexity: O(V^2 E) generally, but O(E * sqrt(V)) for unit networks (like this).
func DinicMatching(adj [][]int, uCount, vCount int) int {
	source := 0
	sink := uCount + vCount + 1
	n := sink + 1

	graph := make([][]Edge, n)

	// Helper to add directional edge with capacity
	addEdge := func(from, to, cap int) {
		graph[from] = append(graph[from], Edge{to, len(graph[to]), cap, 0})
		graph[to] = append(graph[to], Edge{from, len(graph[from]) - 1, 0, 0})
	}

	// 1. Connect Source -> Workers
	for i := 1; i <= uCount; i++ {
		addEdge(source, i, 1)
	}
	// 2. Connect Workers -> Jobs
	for u := 0; u < uCount; u++ {
		for _, v := range adj[u] {
			// Offset job index by uCount + 1 to avoid ID collision
			jobNode := v + uCount + 1
			addEdge(u+1, jobNode, 1)
		}
	}
	// 3. Connect Jobs -> Sink
	for j := 1; j <= vCount; j++ {
		jobNode := j + uCount
		addEdge(jobNode, sink, 1)
	}

	// Dinic Implementation
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
			idx := ptr[u]
			e := &graph[u][idx]
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
	return maxFlow
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
