package findbridge_test

import (
	"log"
	"testing"

	"github.com/skyrocket-qy/gox/dsa/findbridge"
	"github.com/stretchr/testify/assert"
)

func TestRemoveEdge(t *testing.T) {
	// Create a graph given in the above diagram
	g1 := findbridge.NewGraphRemoveEdge(5)
	g1.AddEdge(1, 0)
	g1.AddEdge(0, 2)
	g1.AddEdge(2, 1)
	g1.AddEdge(0, 3)
	g1.AddEdge(3, 4)

	log.Println("Bridges in first graph (Remove Edge Method):")

	bridges := g1.FindBridgesRemoveEdge()
	log.Println(bridges)

	expectedBridges := [][]int{{
		0, 3,
	}, {
		3, 4,
	}}

	// Sort both actual and expected bridges for consistent comparison
	findbridge.SortBridges(bridges)
	findbridge.SortBridges(expectedBridges)

	assert.Equal(t, expectedBridges, bridges, "Bridges should match")
}
