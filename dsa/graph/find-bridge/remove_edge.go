package findbridge

import (
	"fmt"
	"sort"
)

// Graph represents an undirected graph using adjacency list representation
type GraphRemoveEdge struct {
	V     int                 // No. of vertices
	Adj   map[int]map[int]bool // Adjacency list
}

// NewGraphRemoveEdge creates a new GraphRemoveEdge instance
func NewGraphRemoveEdge(vertices int) *GraphRemoveEdge {
	adj := make(map[int]map[int]bool)
	for i := 0; i < vertices; i++ {
		adj[i] = make(map[int]bool)
	}
	return &GraphRemoveEdge{
		V:   vertices,
		Adj: adj,
	}
}

// AddEdge adds an edge to the graph
func (g *GraphRemoveEdge) AddEdge(u, v int) {
	g.Adj[u][v] = true
	g.Adj[v][u] = true
}

func (g *GraphRemoveEdge) dfs(u int, visited []bool) {
	visited[u] = true
	for v := range g.Adj[u] {
		if !visited[v] {
			g.dfs(v, visited)
		}
	}
}

func (g *GraphRemoveEdge) ComputeGroup() int {
	visited := make([]bool, g.V)
	group := 0
	for i := 0; i < g.V; i++ {
		if !visited[i] {
			group++
			g.dfs(i, visited)
		}
	}
	return group
}

func (g *GraphRemoveEdge) GetEdges() [][]int {
	edges := make(map[string]bool)
	var result [][]int
	for u := 0; u < g.V; u++ {
		for v := range g.Adj[u] {
			edge := []int{u, v}
			sort.Ints(edge)
			edgeStr := fmt.Sprintf("%d-%d", edge[0], edge[1])
			if !edges[edgeStr] {
				edges[edgeStr] = true
				result = append(result, []int{u, v})
			}
		}
	}
	return result
}

func (g *GraphRemoveEdge) FindBridgesRemoveEdge() [][]int {
	bridges := [][]int{}

	// get the disconnect subgraph count
	initialGroup := g.ComputeGroup()

	// remove each edge, if subgraph count is increased, that means this edge is bridge
	edgesList := g.GetEdges()
	for _, edge := range edgesList {
		u, v := edge[0], edge[1]

		// Remove edge
		delete(g.Adj[u], v)
		delete(g.Adj[v], u)

		if g.ComputeGroup() > initialGroup {
			bridges = append(bridges, []int{u, v})
		}

		// Add edge back
		g.AddEdge(u, v)
	}

	return bridges
}
