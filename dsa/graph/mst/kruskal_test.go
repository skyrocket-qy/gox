package mst_test

import (
	"reflect"
	"testing"

	"github.com/skyrocket-qy/gox/dsa/graph/mst"
)

func TestKruskal(t *testing.T) {
	g := mst.NewGraph(4)
	g.AddEdge(0, 1, 10)
	g.AddEdge(0, 2, 6)
	g.AddEdge(0, 3, 5)
	g.AddEdge(1, 3, 15)
	g.AddEdge(2, 3, 4)

	expectedMST := []mst.Edge{
		{U: 2, V: 3, Weight: 4},
		{U: 0, V: 3, Weight: 5},
		{U: 0, V: 1, Weight: 10},
	}

	result := g.Kruskal()

	if !reflect.DeepEqual(result, expectedMST) {
		t.Errorf("Kruskal() = %v, want %v", result, expectedMST)
	}
}
