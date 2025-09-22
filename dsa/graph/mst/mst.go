package mst

import (
	"math"
	"sort"

	"github.com/skyrocket-qy/gox/dsa/alg/unionfind"
)

// An Edge represents a connection between two points with a specific cost.
type Edge struct {
	point1 int
	point2 int
	cost   int
}

func Kruskal(points [][]int) int {
	n := len(points)
	if n <= 1 {
		return 0
	}

	// 1. Generate all possible edges and their costs (Manhattan distance).
	edges := []Edge{}

	for i := range n {
		for j := i + 1; j < n; j++ {
			cost := abs(points[i][0]-points[j][0]) + abs(points[i][1]-points[j][1])
			edges = append(edges, Edge{i, j, cost})
		}
	}

	// 2. Sort all edges by cost in non-decreasing order.
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].cost < edges[j].cost
	})

	// 3. Build the MST using Union-Find.
	uf := unionfind.New[int]()
	totalCost := 0
	edgesUsed := 0

	for _, edge := range edges {
		// If the points are not already connected, unite them.
		if uf.Union(edge.point1, edge.point2) {
			totalCost += edge.cost
			edgesUsed++
			// 4. Stop when we have N-1 edges, as the tree is complete.
			if edgesUsed == n-1 {
				break
			}
		}
	}

	return totalCost
}

// Helper function for absolute value.
func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func Prims(points [][]int) int {
	n := len(points)
	if n <= 1 {
		return 0
	}

	// 1. Initialization.
	// `visited` tracks which points are already in our MST.
	visited := make([]bool, n)
	// `minCost[i]` stores the minimum cost to connect point `i` to the MST.
	minCost := make([]int, n)
	for i := range minCost {
		minCost[i] = math.MaxInt32
	}

	// Start with the first point (point 0).
	minCost[0] = 0
	totalCost := 0

	// 2. Grow the tree until all `n` points are included.
	for range n {
		// Find the next closest, unvisited point to add to the MST.
		currPoint := -1
		currMinCost := math.MaxInt32

		for i := range n {
			if !visited[i] && minCost[i] < currMinCost {
				currMinCost = minCost[i]
				currPoint = i
			}
		}

		// If no unvisited point is found, something is wrong (should not happen in this problem).
		if currPoint == -1 {
			break
		}

		// Add the chosen point to the MST.
		visited[currPoint] = true
		totalCost += currMinCost

		// 3. Update the `minCost` for all neighbors of the newly added point.
		// For every unvisited point, check if connecting it via `currPoint` is cheaper.
		for i := range n {
			if !visited[i] {
				cost := abs(
					points[currPoint][0]-points[i][0],
				) + abs(
					points[currPoint][1]-points[i][1],
				)
				if cost < minCost[i] {
					minCost[i] = cost
				}
			}
		}
	}

	return totalCost
}
