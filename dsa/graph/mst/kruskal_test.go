package mst_test

import (
	"testing"

	"github.com/skyrocket-qy/gox/dsa/graph/mst"
)

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
		t.Run(tc.name+"_kruskal", func(t *testing.T) {
			got := mst.Kruskal(tc.points)
			if got != tc.want {
				t.Errorf("Kruskal() = %v, want %v", got, tc.want)
			}
		})
	}

	for _, tc := range testCases {
		t.Run(tc.name+"_prims", func(t *testing.T) {
			got := mst.Prims(tc.points)
			if got != tc.want {
				t.Errorf("Prims() = %v, want %v", got, tc.want)
			}
		})
	}
}