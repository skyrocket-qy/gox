package bfs_test

import (
	"reflect"
	"testing"

	"github.com/skyrocket-qy/gox/dsa/graph/bfs"
)

func TestBfs(t *testing.T) {
	// A graph with a cycle
	graphWithCycle := map[int][]int{
		0: {1, 2},
		1: {2},
		2: {0, 3},
		3: {3},
	}

	testCases := []struct {
		name     string
		graph    map[int][]int
		root     int
		expected []int
	}{
		{
			name: "Complex Graph",
			graph: map[int][]int{
				0: {1, 2},
				1: {3, 4},
				2: {4},
				3: {5},
				4: {5},
				5: {},
			},
			root:     0,
			expected: []int{0, 1, 2, 3, 4, 5},
		},
		{
			name:     "Single Node Graph",
			graph:    map[int][]int{0: {}},
			root:     0,
			expected: []int{0},
		},
		{
			name: "Linear Graph",
			graph: map[int][]int{
				0: {1},
				1: {2},
				2: {3},
				3: {},
			},
			root:     0,
			expected: []int{0, 1, 2, 3},
		},
		{
			name:     "Empty graph",
			graph:    map[int][]int{},
			root:     0,
			expected: []int{0},
		},
		{
			name:     "Graph with cycle",
			graph:    graphWithCycle,
			root:     0,
			expected: []int{0, 1, 2, 3},
		},
		{
			name: "Disconnected graph",
			graph: map[int][]int{
				0: {1},
				1: {},
				2: {3},
			},
			root:     0,
			expected: []int{0, 1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name+" Recursive", func(t *testing.T) {
			result := bfs.BfsRecursive(tc.graph, tc.root)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("BfsRecursive failed, expected %v, got %v", tc.expected, result)
			}
		})

		t.Run(tc.name+" Iterative", func(t *testing.T) {
			result := bfs.BfsIterative(tc.graph, tc.root)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("BfsIterative failed, expected %v, got %v", tc.expected, result)
			}
		})
	}
}
