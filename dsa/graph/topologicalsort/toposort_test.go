package topologicalsort_test

import (
	"testing"

	"github.com/skyrocket-qy/gox/dsa/graph/topologicalsort"
	"github.com/stretchr/testify/assert"
)

func TestTopoSortRemoveVertix(t *testing.T) {
	// Simple graph
	graph1 := map[int][]int{
		1: {2, 3},
		2: {4},
		3: {4},
		4: {},
	}
	// The result is not unique, so we need to check if it's a valid topological sort
	isValidTopoSort(t, graph1, topologicalsort.TopoSortRemoveVertix(graph1))

	// Complex graph
	graph2 := map[int][]int{
		5: {2, 0},
		4: {0, 1},
		2: {3},
		3: {1},
		0: {},
		1: {},
	}
	isValidTopoSort(t, graph2, topologicalsort.TopoSortRemoveVertix(graph2))

	// Graph with a cycle
	graph3 := map[int][]int{
		1: {2},
		2: {3},
		3: {1},
	}
	assert.Nil(t, topologicalsort.TopoSortRemoveVertix(graph3))
}

func TestTopoSortDfs(t *testing.T) {
	// Simple graph
	graph1 := map[int][]int{
		1: {2, 3},
		2: {4},
		3: {4},
		4: {},
	}
	isValidTopoSort(t, graph1, topologicalsort.TopoSortDfs(graph1))

	// Complex graph
	graph2 := map[int][]int{
		5: {2, 0},
		4: {0, 1},
		2: {3},
		3: {1},
		0: {},
		1: {},
	}
	isValidTopoSort(t, graph2, topologicalsort.TopoSortDfs(graph2))

	// Graph with a cycle
	graph3 := map[int][]int{
		1: {2},
		2: {3},
		3: {1},
	}
	assert.Nil(t, topologicalsort.TopoSortDfs(graph3))
}

// isValidTopoSort checks if a given sequence is a valid topological sort of a graph.
func isValidTopoSort(t *testing.T, graph map[int][]int, seq []int) {
	assert.True(t, isValidTopoSortHelper(graph, seq))
}

func isValidTopoSortHelper(graph map[int][]int, seq []int) bool {
	if seq == nil {
		return false
	}
	pos := make(map[int]int)
	for i, node := range seq {
		pos[node] = i
	}

	if len(pos) != len(graph) {
		return false
	}

	for node, edges := range graph {
		for _, edge := range edges {
			if pos[node] > pos[edge] {
				return false
			}
		}
	}

	return true
}
