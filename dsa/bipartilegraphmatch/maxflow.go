package bipartilegraphmatch

type Edge struct {
	to, rev   int
	cap, flow int
}

// DinicMatching: Solves using Flow Network construction.
// Complexity: O(V^2 E) generally, but O(E * sqrt(V)) for unit networks (like this).
func DinicMatching[T comparable](adj map[T][]T, uCount, vCount int) map[T]T {
	// Map T to int IDs
	// We must distinguish between U nodes and V nodes.
	// Source = 0
	// U nodes = 1..N
	// V nodes = N+1..M
	// Sink = M+1

	uIDMap := make(map[T]int)
	vIDMap := make(map[T]int)
	revVMap := make(map[int]T) // ID -> T (for V nodes only)

	nextID := 1

	getUID := func(node T) int {
		if id, ok := uIDMap[node]; ok {
			return id
		}
		id := nextID
		nextID++
		uIDMap[node] = id
		return id
	}

	getVID := func(node T) int {
		if id, ok := vIDMap[node]; ok {
			return id
		}
		id := nextID
		nextID++
		vIDMap[node] = id
		revVMap[id] = node
		return id
	}

	source := 0
	sink := -1 // Will set later

	// Collect edges and nodes
	type edgeReq struct {
		u, v T
	}
	var edges []edgeReq

	uNodes := make(map[T]bool)
	vNodes := make(map[T]bool)

	for u, neighbors := range adj {
		uNodes[u] = true
		for _, v := range neighbors {
			vNodes[v] = true
			edges = append(edges, edgeReq{u, v})
		}
	}

	// Assign IDs
	// Source = 0
	// U nodes
	for u := range uNodes {
		getUID(u)
	}
	// V nodes
	for v := range vNodes {
		getVID(v)
	}

	sink = nextID
	n := sink + 1

	graph := make([][]Edge, n)

	addEdge := func(from, to, cap int) {
		graph[from] = append(graph[from], Edge{to, len(graph[to]), cap, 0})
		graph[to] = append(graph[to], Edge{from, len(graph[from]) - 1, 0, 0})
	}

	// 1. Connect Source -> Workers
	for u := range uNodes {
		uid := getUID(u)
		addEdge(source, uid, 1)
	}

	// 2. Connect Workers -> Jobs
	for _, e := range edges {
		uid := getUID(e.u)
		vid := getVID(e.v)
		addEdge(uid, vid, 1)
	}

	// 3. Connect Jobs -> Sink
	for v := range vNodes {
		vid := getVID(v)
		addEdge(vid, sink, 1)
	}

	// Dinic Implementation (same as before, just using graph)
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

	for bfs() {
		for i := range ptr {
			ptr[i] = 0
		}
		for {
			pushed := dfs(source, 999999999)
			if pushed == 0 {
				break
			}
		}
	}

	// Extract matching
	result := make(map[T]T)
	// Iterate over U nodes and find edges with flow=1 to V nodes
	for u := range uNodes {
		uid := getUID(u)
		for _, e := range graph[uid] {
			// Check if edge is to a V node and has flow 1
			// V nodes are not source or sink.
			// e.to must be a V node ID.
			// And flow must be 1.
			if e.flow == 1 && e.to != source && e.to != sink {
				// Verify e.to is a V node (it should be if graph construction is correct)
				if v, ok := revVMap[e.to]; ok {
					result[u] = v
					break // One match per worker
				}
			}
		}
	}

	return result
}
