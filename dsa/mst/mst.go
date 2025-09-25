package mst

import (
	"math"
	"sort"
)

// Edge represents a weighted edge in a graph, suitable for Kruskal's algorithm.
// It's generic and can be used with any underlying graph representation.
type Edge struct {
	U, V int // The two vertices connected by the edge.
	Cost int // The cost or weight of the edge.
}

// Kruskal finds the Minimum Spanning Tree (MST) cost for a given graph.
// It's best for sparse graphs.
//
// Parameters:
//   - n: The total number of vertices in the graph.
//   - edges: A slice of all edges in the graph.
//
// Returns:
//
//	The total cost of the MST.
func Kruskal(n int, edges []Edge) int {
	if n <= 1 {
		return 0
	}

	// 1. Sort all edges by cost in non-decreasing order.
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Cost < edges[j].Cost
	})

	// 2. Build the MST using a Union-Find data structure.
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}

	var find func(i int) int

	find = func(i int) int {
		if parent[i] == i {
			return i
		}

		parent[i] = find(parent[i])

		return parent[i]
	}

	union := func(i, j int) bool {
		rootI := find(i)

		rootJ := find(j)
		if rootI != rootJ {
			parent[rootI] = rootJ

			return true
		}

		return false
	}

	totalCost := 0
	edgesUsed := 0

	for _, edge := range edges {
		// 3. If the vertices of the edge are not already connected, add the edge to the MST.
		if union(edge.U, edge.V) {
			totalCost += edge.Cost
			edgesUsed++
			// 4. Stop when we have N-1 edges, completing the tree.
			if edgesUsed == n-1 {
				break
			}
		}
	}

	return totalCost
}

// Prims finds the Minimum Spanning Tree (MST) cost for a given graph.
// It's efficient for dense graphs.
//
// Parameters:
//   - n: The total number of vertices in the graph.
//   - adj: An adjacency list where adj[i] is a list of pairs [neighbor, cost].
//
// Returns:
//
//	The total cost of the MST.
func Prims(n int, adj map[int][][2]int) int {
	if n <= 1 {
		return 0
	}

	// 1. Initialization.
	// `visited` tracks vertices already in the MST.
	visited := make([]bool, n)
	// `minCost[i]` stores the minimum cost to connect vertex `i` to the MST.
	minCost := make([]int, n)
	for i := range minCost {
		minCost[i] = math.MaxInt32
	}

	// Start with the first vertex (vertex 0).
	minCost[0] = 0
	totalCost := 0

	// 2. Grow the tree until all `n` vertices are included.
	for range n {
		// Find the next closest, unvisited vertex to add to the MST.
		u := -1
		for v := range n {
			if !visited[v] && (u == -1 || minCost[v] < minCost[u]) {
				u = v
			}
		}

		// If no connected, unvisited vertex is found, the graph is disconnected.
		if u == -1 || minCost[u] == math.MaxInt32 {
			// Depending on the problem, you might return an error or a specific value.
			// For many LeetCode problems, this indicates it's impossible to connect all points.
			return -1 // Or handle as per problem requirements.
		}

		// Add the chosen vertex to the MST.
		visited[u] = true
		totalCost += minCost[u]

		// 3. Update the `minCost` for all neighbors of the newly added vertex.
		if neighbors, ok := adj[u]; ok {
			for _, edge := range neighbors {
				v, cost := edge[0], edge[1]
				if !visited[v] && cost < minCost[v] {
					minCost[v] = cost
				}
			}
		}
	}

	return totalCost
}
