package mst

import (
	"sort"
)

type Edge struct {
	U, V, Weight int
}

type Graph struct {
	V     int
	Edges []Edge
}

func NewGraph(vertices int) *Graph {
	return &Graph{
		V: vertices,
	}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.Edges = append(g.Edges, Edge{U: u, V: v, Weight: w})
}

func (g *Graph) find(parent []int, i int) int {
	if parent[i] == i {
		return i
	}
	parent[i] = g.find(parent, parent[i]) // Path compression
	return parent[i]
}

func (g *Graph) union(parent []int, rank []int, x, y int) {
	xroot := g.find(parent, x)
	yroot := g.find(parent, y)

	if rank[xroot] < rank[yroot] {
		parent[xroot] = yroot
	} else if rank[xroot] > rank[yroot] {
		parent[yroot] = xroot
	} else {
		parent[yroot] = xroot
		rank[xroot]++
	}
}

func (g *Graph) Kruskal() []Edge {
	result := []Edge{}
	i, e := 0, 0 // i is used to iterate through sorted edges, e is used for result array index

	// Sort edges by weight
	sort.Slice(g.Edges, func(i, j int) bool {
		return g.Edges[i].Weight < g.Edges[j].Weight
	})

	parent := make([]int, g.V)
	rank := make([]int, g.V)

	for node := 0; node < g.V; node++ {
		parent[node] = node
		rank[node] = 0
	}

	for e < g.V-1 && i < len(g.Edges) {
		edge := g.Edges[i]
		i++
		u, v, w := edge.U, edge.V, edge.Weight

		x := g.find(parent, u)
		y := g.find(parent, v)

		if x != y {
			e++
			result = append(result, Edge{U: u, V: v, Weight: w})
			g.union(parent, rank, x, y)
		}
	}

	return result
}


