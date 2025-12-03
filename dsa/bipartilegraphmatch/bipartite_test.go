package bipartilegraphmatch

import (
	"math/rand"
	"testing"
)

func TestMatchingInt(t *testing.T) {
	// U: {1, 2, 3}
	// V: {4, 5, 6}
	// Edges: 1-4, 1-5, 2-4, 3-6
	adj := map[int][]int{
		1: {4, 5},
		2: {4},
		3: {6},
	}
	uCount := 3
	vCount := 3

	expected := 3 // 1-5, 2-4, 3-6

	if res := BacktrackingMatching(adj, uCount, vCount); res != expected {
		t.Errorf("BacktrackingMatching failed: expected %d, got %d", expected, res)
	}
	if res := HopcroftKarp(adj, uCount, vCount); res != expected {
		t.Errorf("HopcroftKarp failed: expected %d, got %d", expected, res)
	}
	if res := KuhnsAlgorithm(adj, uCount, vCount); res != expected {
		t.Errorf("KuhnsAlgorithm failed: expected %d, got %d", expected, res)
	}
	if res := DinicMatching(adj, uCount, vCount); res != expected {
		t.Errorf("DinicMatching failed: expected %d, got %d", expected, res)
	}
}

func TestMatchingString(t *testing.T) {
	// U: {"A", "B"}
	// V: {"X", "Y"}
	// Edges: A-X, A-Y, B-X
	adj := map[string][]string{
		"A": {"X", "Y"},
		"B": {"X"},
	}
	uCount := 2
	vCount := 2

	expected := 2 // A-Y, B-X

	if res := BacktrackingMatching(adj, uCount, vCount); res != expected {
		t.Errorf("BacktrackingMatching failed: expected %d, got %d", expected, res)
	}
	if res := HopcroftKarp(adj, uCount, vCount); res != expected {
		t.Errorf("HopcroftKarp failed: expected %d, got %d", expected, res)
	}
	if res := KuhnsAlgorithm(adj, uCount, vCount); res != expected {
		t.Errorf("KuhnsAlgorithm failed: expected %d, got %d", expected, res)
	}
	if res := DinicMatching(adj, uCount, vCount); res != expected {
		t.Errorf("DinicMatching failed: expected %d, got %d", expected, res)
	}
}

// Benchmarks

func generateGraph(uCount, vCount int, density float64) map[int][]int {
	rng := rand.New(rand.NewSource(12345))
	adj := make(map[int][]int)
	// Ensure U and V are disjoint
	// U: 0..uCount-1
	// V: uCount..uCount+vCount-1

	for u := 0; u < uCount; u++ {
		var neighbors []int
		for v := 0; v < vCount; v++ {
			if rng.Float64() < density {
				neighbors = append(neighbors, v+uCount)
			}
		}
		if len(neighbors) > 0 {
			adj[u] = neighbors
		}
	}
	return adj
}

func runBenchmark(b *testing.B, u, v int, density float64, includeBacktracking bool) {
	adj := generateGraph(u, v, density)

	if includeBacktracking {
		b.Run("Backtracking", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				BacktrackingMatching(adj, u, v)
			}
		})
	}

	b.Run("Kuhn", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			KuhnsAlgorithm(adj, u, v)
		}
	})

	b.Run("HopcroftKarp", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HopcroftKarp(adj, u, v)
		}
	})

	b.Run("Dinic", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			DinicMatching(adj, u, v)
		}
	})
}

func Benchmark_Small_V10_E25(b *testing.B) {
	runBenchmark(b, 10, 10, 0.5, true)
}

func Benchmark_Medium_V100_E1000(b *testing.B) {
	runBenchmark(b, 100, 100, 0.1, false)
}

func Benchmark_Large_V1000_E5000(b *testing.B) {
	runBenchmark(b, 1000, 1000, 0.005, false)
}

func Benchmark_Dense_V200_E20000(b *testing.B) {
	runBenchmark(b, 200, 200, 0.5, false)
}
