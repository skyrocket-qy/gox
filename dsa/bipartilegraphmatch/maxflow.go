package bipartilegraphmatch

type Edge struct {
	to, rev   int
	cap, flow int
}

// DinicMatching: Solves using Flow Network construction.
// Complexity: O(V^2 E) generally, but O(E * sqrt(V)) for unit networks (like this).
func DinicMatching[T comparable](adj map[T][]T, uCount, vCount int) map[T]T {
	// Map T to int IDs
	// Source = 0
	// Workers = 1..uCount (or just mapped)
	// Jobs = uCount+1..uCount+vCount (or just mapped)
	// Sink = last

	// We need to assign IDs to all nodes.
	// Since we don't have a list of all jobs (V nodes) explicitly, we must scan adj.

	idMap := make(map[T]int)
	revMap := make(map[int]T) // ID -> T
	nextID := 1

	getID := func(node T) int {
		if id, ok := idMap[node]; ok {
			return id
		}
		id := nextID
		nextID++
		idMap[node] = id
		revMap[id] = node
		return id
	}

	// We need to differentiate between U nodes and V nodes to avoid ID collision if they share values?
	// In bipartite matching, U and V are disjoint sets.
	// If T is int, and U={1}, V={1}, they are the same value.
	// But in the input `adj`, keys are U, values are V.
	// If the graph is bipartite, the sets are disjoint.
	// However, if the user passes `adj` where `1` is in U and `1` is in V (as a neighbor of some U),
	// does it mean the same node?
	// Usually in bipartite matching problems, U and V are distinct entities.
	// But if T is `string`, "worker1" and "job1" are distinct.
	// If T is `int`, `1` (worker) and `1` (job) are the same value.
	// If the user intends them to be distinct, they should use different T values or we need to handle them as distinct.
	// The original code `jobNode := v + uCount + 1` implies that `v` (0..vCount-1) is shifted.
	// This implies the input `adj` used 0-based indices for V nodes, which overlap with U nodes 0..uCount-1.
	// So `adj` was `[][]int` where `adj[u]` contains indices of V.
	// Now with `map[T][]T`, `adj[u]` contains the *actual* V nodes (of type T).
	// If `T` is `string`, "A" -> ["X", "Y"]. "A" and "X" are distinct.
	// If `T` is `int`, 0 -> [0, 1]. Worker 0 connects to Job 0 and Job 1.
	// If Worker 0 and Job 0 are the same integer, are they the same node?
	// In a bipartite graph, edges are between disjoint sets.
	// If the user uses `int` for both, they might mean "Worker 0" and "Job 0".
	// But if they are the same value, a map `idMap` will give them the same ID.
	// This would create a self-loop or internal edge if we are not careful.
	// BUT, Dinic builds a flow network: Source -> U -> V -> Sink.
	// If U node `x` and V node `x` are the same `T` value, `idMap` gives same ID.
	// Then we have Source -> ID(x) and ID(x) -> Sink.
	// And edges U->V. If U=x connects to V=y.
	// If x and y are distinct, fine.
	// If x connects to x (Worker x connects to Job x), then ID(x) -> ID(x). Self loop.
	// This is valid in flow network but useless for matching.
	// However, the structure Source->U->V->Sink requires U and V to be *layers*.
	// If a node is both U and V, it can receive flow from Source and send to Sink.
	// But can it flow to itself?
	// The issue is if we treat them as the same node in the flow network.
	// In the original code, `jobNode` was shifted. This explicitly separated U and V.
	// With `map[T][]T`, if the user provides `1 -> [1]`, do they mean Worker 1 -> Job 1?
	// If so, they are distinct entities.
	// If `T` is `int`, we cannot distinguish `1` from `1` unless we know context.
	// BUT, usually with generic maps, the keys and values *are* the unique identifiers.
	// If Worker 1 and Job 1 are distinct, they should have different values (e.g. "W1", "J1").
	// If the user uses `int` and expects implicit separation, `map[T][]T` is ambiguous.
	// Given `T` is `comparable`, I will assume `T` values are unique identifiers for nodes in the graph.
	// i.e. U and V are disjoint sets of values.

	source := 0
	sink := -1 // Will set later

	// We need to construct the graph.
	// We can't know the full set of V nodes without iterating.
	// And we need to assign IDs.

	// Let's build the edges list first, then assign IDs.
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
		getID(u)
	}
	// V nodes
	for v := range vNodes {
		getID(v)
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
		uid := getID(u)
		addEdge(source, uid, 1)
	}

	// 2. Connect Workers -> Jobs
	for _, e := range edges {
		uid := getID(e.u)
		vid := getID(e.v)
		addEdge(uid, vid, 1)
	}

	// 3. Connect Jobs -> Sink
	for v := range vNodes {
		vid := getID(v)
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
		uid := getID(u)
		for _, e := range graph[uid] {
			// Check if edge is to a V node and has flow 1
			// V nodes are not source or sink.
			// e.to must be a V node ID.
			// And flow must be 1.
			if e.flow == 1 && e.to != source && e.to != sink {
				// Verify e.to is a V node (it should be if graph construction is correct)
				if v, ok := revMap[e.to]; ok {
					if vNodes[v] {
						result[u] = v
						break // One match per worker
					}
				}
			}
		}
	}

	return result
}
