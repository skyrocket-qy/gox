package bipartilegraphmatch

import (
	"fmt"
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

	expectedSize := 3 // 1-5, 2-4, 3-6

	checkResult := func(name string, res map[int]int) {
		if len(res) != expectedSize {
			t.Errorf("%s failed: expected size %d, got %d", name, expectedSize, len(res))
		}
		// Verify validity (edges exist)
		for u, v := range res {
			found := false
			for _, neighbor := range adj[u] {
				if neighbor == v {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("%s returned invalid edge: %d-%d", name, u, v)
			}
		}
	}

	checkResult("BacktrackingMatching", BacktrackingMatching(adj, uCount, vCount))
	checkResult("HopcroftKarp", HopcroftKarp(adj, uCount, vCount))
	checkResult("KuhnsAlgorithm", KuhnsAlgorithm(adj, uCount, vCount))
	checkResult("DinicMatching", DinicMatching(adj, uCount, vCount))
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

	expectedSize := 2 // A-Y, B-X

	checkResult := func(name string, res map[string]string) {
		if len(res) != expectedSize {
			t.Errorf("%s failed: expected size %d, got %d", name, expectedSize, len(res))
		}
		for u, v := range res {
			found := false
			for _, neighbor := range adj[u] {
				if neighbor == v {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("%s returned invalid edge: %s-%s", name, u, v)
			}
		}
	}

	checkResult("BacktrackingMatching", BacktrackingMatching(adj, uCount, vCount))
	checkResult("HopcroftKarp", HopcroftKarp(adj, uCount, vCount))
	checkResult("KuhnsAlgorithm", KuhnsAlgorithm(adj, uCount, vCount))
	checkResult("DinicMatching", DinicMatching(adj, uCount, vCount))
}

func TestHopcroftKarpVisualizer(t *testing.T) {
	// Example from explanation
	// Workers: A, B
	// Jobs: a, b
	// Edges: A-a, A-b, B-a
	adj := map[string][]string{
		"A": {"a", "b"},
		"B": {"a"},
	}
	uCount := 2
	vCount := 2

	fmt.Println("\n=== Running Hopcroft-Karp Visualizer ===")
	res := HopcroftKarpWithLogging(adj, uCount, vCount)
	fmt.Printf("Result: %v\n", res)
}

func TestOverlappingIDs(t *testing.T) {
	// Worker IDs: 1, 2
	// Job IDs: 1, 2
	// Edges: Worker 1 -> Job 2, Worker 2 -> Job 1
	// This tests if the algorithms handle key collisions between U and V sets correctly.
	adj := map[int][]int{
		1: {2},
		2: {1},
	}
	uCount := 2
	vCount := 2

	expectedSize := 2

	checkResult := func(name string, res map[int]int) {
		if len(res) != expectedSize {
			t.Errorf("%s failed: expected size %d, got %d", name, expectedSize, len(res))
		}
		// Check specific matches
		if res[1] != 2 {
			t.Errorf("%s failed: Worker 1 should match Job 2, got %d", name, res[1])
		}
		if res[2] != 1 {
			t.Errorf("%s failed: Worker 2 should match Job 1, got %d", name, res[2])
		}
	}

	checkResult("BacktrackingMatching", BacktrackingMatching(adj, uCount, vCount))
	checkResult("HopcroftKarp", HopcroftKarp(adj, uCount, vCount))
	checkResult("KuhnsAlgorithm", KuhnsAlgorithm(adj, uCount, vCount))
	checkResult("DinicMatching", DinicMatching(adj, uCount, vCount))
}

func TestComplexCases(t *testing.T) {
	runTest := func(name string, adj map[int][]int, uCount, vCount, expectedSize int) {
		t.Run(name, func(t *testing.T) {
			check := func(algoName string, res map[int]int) {
				if len(res) != expectedSize {
					t.Errorf("%s failed: expected size %d, got %d", algoName, expectedSize, len(res))
				}
				// Verify validity
				for u, v := range res {
					found := false
					for _, neighbor := range adj[u] {
						if neighbor == v {
							found = true
							break
						}
					}
					if !found {
						t.Errorf("%s returned invalid edge: %d-%d", algoName, u, v)
					}
				}
			}
			check("Backtracking", BacktrackingMatching(adj, uCount, vCount))
			check("HopcroftKarp", HopcroftKarp(adj, uCount, vCount))
			check("Kuhn", KuhnsAlgorithm(adj, uCount, vCount))
			check("Dinic", DinicMatching(adj, uCount, vCount))
		})
	}

	// Case 1: Disconnected
	runTest("Disconnected", map[int][]int{
		1: {},
		2: {},
	}, 2, 2, 0)

	// Case 2: Fully Connected (2x2)
	runTest("FullyConnected", map[int][]int{
		1: {3, 4},
		2: {3, 4},
	}, 2, 2, 2)

	// Case 3: Chain/Path (1-4, 1-5, 2-5, 2-6, 3-6) -> Match: 1-4, 2-5, 3-6
	runTest("Chain", map[int][]int{
		1: {4, 5},
		2: {5, 6},
		3: {6},
	}, 3, 3, 3)

	// Case 4: Star (All workers want Job 4)
	runTest("Star", map[int][]int{
		1: {4},
		2: {4},
		3: {4},
	}, 3, 1, 1)
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
