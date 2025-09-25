package mst_test

import (
	"testing"

	"github.com/skyrocket-qy/gox/dsa/mst"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func calculateEdges(points [][]int) (int, []mst.Edge, map[int][][2]int) {
	n := len(points)
	if n == 0 {
		return 0, nil, nil
	}

	var edges []mst.Edge

	adj := make(map[int][][2]int)

	for i := range n {
		for j := i + 1; j < n; j++ {
			dist := abs(points[i][0]-points[j][0]) + abs(points[i][1]-points[j][1])
			edges = append(edges, mst.Edge{U: i, V: j, Cost: dist})

			adj[i] = append(adj[i], [2]int{j, dist})
			adj[j] = append(adj[j], [2]int{i, dist})
		}
	}

	return n, edges, adj
}

func TestMST(t *testing.T) {
	testCases := []struct {
		name   string
		points [][]int
		want   int
	}{
		{
			name:   "no points",
			points: [][]int{},
			want:   0,
		},
		{
			name:   "one point",
			points: [][]int{{0, 0}},
			want:   0,
		},
		{
			name:   "two points",
			points: [][]int{{0, 0}, {1, 1}},
			want:   2,
		},
		{
			name:   "three points",
			points: [][]int{{0, 0}, {1, 1}, {2, 0}},
			want:   4,
		},
		{
			name:   "five points",
			points: [][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}},
			want:   20,
		},
	}

	for _, tc := range testCases {
		n, edges, adj := calculateEdges(tc.points)
		t.Run(tc.name, func(t *testing.T) {
			t.Run("kruskal", func(t *testing.T) {
				got := mst.Kruskal(n, edges)
				if got != tc.want {
					t.Errorf("Kruskal() = %v, want %v", got, tc.want)
				}
			})

			t.Run("prims", func(t *testing.T) {
				got := mst.Prims(n, adj)
				if got != tc.want {
					t.Errorf("Prims() = %v, want %v", got, tc.want)
				}
			})
		})
	}
}
