package singlesourceshortestpath

import (
	"testing"
)

func TestDijkstraAlgorithm(t *testing.T) {
	// A B C D E
	// 0 1 2 3 4
	graph := [][]int{
		{0, 6, -1, 1, -1}, // A
		{6, 0, 5, 2, 2},   // B
		{-1, 5, 0, -1, 5}, // C
		{1, 2, -1, 0, 1},  // D
		{-1, 2, 5, 1, 0},  // E
	}
	// A -> D -> B -> C : 1 + 2 + 5 = 8
	// A -> D -> E -> B -> C : 1 + 1 + 2 + 5 = 9
	// A -> D -> E -> C : 1 + 1 + 5 = 7
	start := 0
	end := 2
	expected := 7

	result := DijkstraAlgorithm(graph, start, end)
	if result != expected {
		t.Errorf("DijkstraAlgorithm() = %d; want %d", result, expected)
	}
}

func TestLabelSettingAlgorithm(t *testing.T) {
	graph := [][]int{
		{0, 6, -1, 1, -1},
		{6, 0, 5, 2, 2},
		{-1, 5, 0, -1, 5},
		{1, 2, -1, 0, 1},
		{-1, 2, 5, 1, 0},
	}
	start := 0
	end := 2
	expected := 7

	result := LabelSettingAlgorithm(graph, start, end)
	if result != expected {
		t.Errorf("LabelSettingAlgorithm() = %d; want %d", result, expected)
	}
}
