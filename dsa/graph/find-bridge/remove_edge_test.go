package findbridge

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveEdge(t *testing.T) {
	// Create a graph given in the above diagram
	g1 := NewGraphRemoveEdge(5)
	g1.AddEdge(1, 0)
	g1.AddEdge(0, 2)
	g1.AddEdge(2, 1)
	g1.AddEdge(0, 3)
	g1.AddEdge(3, 4)

	fmt.Println("Bridges in first graph (Remove Edge Method):")
	bridges := g1.FindBridgesRemoveEdge()
	fmt.Println(bridges)

	expectedBridges := [][]int{{
		0, 3,
	}, {
		3, 4,
	}}

	// Sort both actual and expected bridges for consistent comparison
	sortBridgesRemoveEdge(bridges)
	sortBridgesRemoveEdge(expectedBridges)

	assert.Equal(t, expectedBridges, bridges, "Bridges should match")
}

// Helper function to sort bridges for consistent comparison
func sortBridgesRemoveEdge(bridges [][]int) {
	sort.Slice(bridges, func(i, j int) bool {
		if bridges[i][0] != bridges[j][0] {
			return bridges[i][0] < bridges[j][0]
		}
		return bridges[i][1] < bridges[j][1]
	})
	for i := range bridges {
		if bridges[i][0] > bridges[i][1] {
			bridges[i][0], bridges[i][1] = bridges[i][1], bridges[i][0]
		}
	}
}
