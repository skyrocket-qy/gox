package findbridge

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTarjan(t *testing.T) {
	// Create a graph given in the above diagram
	g1 := NewGraph(5)
	g1.AddEdge(1, 0)
	g1.AddEdge(0, 2)
	g1.AddEdge(2, 1)
	g1.AddEdge(0, 3)
	g1.AddEdge(3, 4)

	log.Println("Bridges in first graph ")

	bridges := g1.FindBridges()
	log.Println(bridges)

	expectedBridges := [][]int{{
		0, 3,
	}, {
		3, 4,
	}}

	// Sort both actual and expected bridges for consistent comparison
	sortBridges(bridges)
	sortBridges(expectedBridges)

	assert.Equal(t, expectedBridges, bridges, "Bridges should match")
}
