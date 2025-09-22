package astar_test

import (
	"math"
	"testing"

	"github.com/skyrocket-qy/gox/dsa/graph/astar"
)

func TestAStar(t *testing.T) {
	// Test case 1: Simple path
	grid1 := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	path1, cost1 := astar.AStar(grid1, 0, 0, 2, 2, astar.ManhattanDistance)
	if path1 == nil || cost1 == 0 {
		t.Errorf("Test Case 1 Failed: Expected a path, got nil or 0 cost")
	}

	if len(path1) != 3 ||
		math.Abs(cost1-2*math.Sqrt(2)) > 0.001 { // Expected 2*sqrt(2) for 8-directional
		t.Errorf(
			"Test Case 1 Failed: Expected path length 3 and cost around 2.828, got %d and %f",
			len(path1),
			cost1,
		)
	}

	// Test case 2: Path with obstacle
	grid2 := [][]int{
		{0, 1, 0},
		{0, 1, 0},
		{0, 0, 0},
	}

	path2, cost2 := astar.AStar(grid2, 0, 0, 2, 2, astar.ManhattanDistance)
	if path2 == nil || cost2 == 0 {
		t.Errorf("Test Case 2 Failed: Expected a path, got nil or 0 cost")
	}

	if len(path2) != 4 ||
		math.Abs(cost2-(2+math.Sqrt(2))) > 0.001 { // Expected 2 + sqrt(2) for 8-directional
		t.Errorf(
			"Test Case 2 Failed: Expected path length 4 and cost around 3.414, got %d and %f",
			len(path2),
			cost2,
		)
	}

	// Test case 3: No path
	grid3 := [][]int{
		{0, 1, 0},
		{0, 1, 0},
		{0, 1, 0},
	}

	path3, cost3 := astar.AStar(grid3, 0, 0, 2, 2, astar.ManhattanDistance)
	if path3 != nil || cost3 != 0 {
		t.Errorf("Test Case 3 Failed: Expected no path, got path or non-zero cost")
	}

	// Test case 4: Start and goal are the same
	grid4 := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	path4, cost4 := astar.AStar(grid4, 0, 0, 0, 0, astar.ManhattanDistance)
	if path4 == nil || cost4 != 0 {
		t.Errorf("Test Case 4 Failed: Expected path with cost 0, got nil or non-zero cost")
	}

	if len(path4) != 1 {
		t.Errorf("Test Case 4 Failed: Expected path length 1, got %d", len(path4))
	}

	// Test case 5: Invalid start coordinates
	path5, cost5 := astar.AStar(grid1, -1, 0, 2, 2, astar.ManhattanDistance)
	if path5 != nil || cost5 != 0 {
		t.Errorf("Test Case 5 Failed: Expected no path, got path or non-zero cost")
	}

	// Test case 6: Invalid goal coordinates
	path6, cost6 := astar.AStar(grid1, 0, 0, 3, 3, astar.ManhattanDistance)
	if path6 != nil || cost6 != 0 {
		t.Errorf("Test Case 6 Failed: Expected no path, got path or non-zero cost")
	}

	// Test case 7: Start is an obstacle
	grid7 := [][]int{
		{1, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	path7, cost7 := astar.AStar(grid7, 0, 0, 2, 2, astar.ManhattanDistance)
	if path7 != nil || cost7 != 0 {
		t.Errorf("Test Case 7 Failed: Expected no path, got path or non-zero cost")
	}

	// Test case 8: Goal is an obstacle
	grid8 := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 1},
	}

	path8, cost8 := astar.AStar(grid8, 0, 0, 2, 2, astar.ManhattanDistance)
	if path8 != nil || cost8 != 0 {
		t.Errorf("Test Case 8 Failed: Expected no path, got path or non-zero cost")
	}

	// Test case 9: Empty grid
	grid9 := [][]int{}

	path9, cost9 := astar.AStar(grid9, 0, 0, 0, 0, astar.ManhattanDistance)
	if path9 != nil || cost9 != 0 {
		t.Errorf("Test Case 9 Failed: Expected no path, got path or non-zero cost")
	}

	// Test case 10: Single cell grid, start and goal same
	grid10 := [][]int{
		{0},
	}

	path10, cost10 := astar.AStar(grid10, 0, 0, 0, 0, astar.ManhattanDistance)
	if path10 == nil || cost10 != 0 || len(path10) != 1 {
		t.Errorf(
			"Test Case 10 Failed: Expected path with cost 0 and length 1, got nil or non-zero cost or wrong length",
		)
	}

	// Test case 11: Single cell grid, start and goal different (invalid)
	path11, cost11 := astar.AStar(grid10, 0, 0, 0, 1, astar.ManhattanDistance)
	if path11 != nil || cost11 != 0 {
		t.Errorf("Test Case 11 Failed: Expected no path, got path or non-zero cost")
	}

	// Test case 12: Test with EuclideanDistance heuristic
	grid12 := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	path12, cost12 := astar.AStar(grid12, 0, 0, 2, 2, astar.EuclideanDistance)
	if path12 == nil || cost12 == 0 {
		t.Errorf("Test Case 12 Failed: Expected a path, got nil or 0 cost")
	}
	// Path length and cost might differ slightly due to diagonal movement cost (sqrt(2))
	// For a 3x3 grid from (0,0) to (2,2), a direct diagonal path would be 2*sqrt(2) if allowed.
	// With 8-directional movement, the cost is 2*sqrt(2) + 2*1 = 4.828
	// The path length would be 3 nodes (start, middle, end) if direct diagonal is allowed.
	// Given our current implementation, it's 8-directional, so it should find a shorter path.
	// Let's check if the cost is reasonable.
	if cost12 < 2.8 || cost12 > 3.0 { // Expecting 2*sqrt(2) for direct diagonal path
		t.Errorf("Test Case 12 Failed: Expected cost around 2.828, got %f", cost12)
	}
}
