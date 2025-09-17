package topologicalsort_test

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/skyrocket-qy/gox/dsa/graph/topologicalsort"
)

// Helper function to create a worker that records the processing order.
func newOrderTrackingWorker(order *[]string, mu *sync.Mutex) topologicalsort.WorkerFunc {
	return func(ctx context.Context, node any) error {
		// Simulate some work
		time.Sleep(10 * time.Millisecond)

		mu.Lock()
		defer mu.Unlock()

		*order = append(*order, node.(string))

		return nil
	}
}

// Helper to check if a is a valid topological order for graph dependencies.
// It ensures that for any node `n` in the order, all of its dependencies appeared before it.
func isValidOrder(order []string, graph map[any][]any) (bool, string) {
	positions := make(map[string]int)
	for i, node := range order {
		positions[node] = i
	}

	for nodeAny, depsAny := range graph {
		node := nodeAny.(string)

		nodePos, exists := positions[node]
		if !exists {
			// This can happen if the worker fails mid-process. It's not an invalid order, just
			// incomplete.
			continue
		}

		for _, depAny := range depsAny {
			dep := depAny.(string)

			depPos, exists := positions[dep]
			if !exists {
				continue
			}

			if nodePos < depPos {
				return false, fmt.Sprintf(
					"invalid order: node '%s' appeared at position %d before its dependency '%s' at position %d",
					node,
					nodePos,
					dep,
					depPos,
				)
			}
		}
	}

	return true, ""
}

func TestConcurrentTopologicalSort_Success(t *testing.T) {
	t.Run("simple linear graph", func(t *testing.T) {
		// A -> B -> C (A depends on B, B depends on C)
		// Shutdown order should be A, then B, then C
		graph := map[any][]any{
			"A": {"B"},
			"B": {"C"},
			"C": {},
		}

		var (
			order []string
			mu    sync.Mutex
		)

		worker := newOrderTrackingWorker(&order, &mu)

		err := topologicalsort.ConcurrentTopologicalSort(context.Background(), graph, worker)
		if err != nil {
			t.Fatalf("Expected no error, but got: %v", err)
		}

		if len(order) != 3 {
			t.Fatalf("Expected 3 items to be processed, but got %d: %v", len(order), order)
		}

		// The only valid order is C, B, A
		expectedOrder := "C,B,A"

		actualOrder := strings.Join(order, ",")
		if actualOrder != expectedOrder {
			t.Errorf("Expected order %s, but got %s", expectedOrder, actualOrder)
		}
	})

	t.Run("complex DAG", func(t *testing.T) {
		// Graph:
		// A -> B, C
		// B -> D
		// C -> D, E
		// Expected shutdown: D and E can run first, then B and C, finally A.
		graph := map[any][]any{
			"A": {"B", "C"},
			"B": {"D"},
			"C": {"D", "E"},
			"D": {},
			"E": {},
		}

		var (
			order []string
			mu    sync.Mutex
		)

		worker := newOrderTrackingWorker(&order, &mu)

		err := topologicalsort.ConcurrentTopologicalSort(context.Background(), graph, worker)
		if err != nil {
			t.Fatalf("Expected no error, but got: %v", err)
		}

		if len(order) != 5 {
			t.Fatalf("Expected 5 items to be processed, but got %d: %v", len(order), order)
		}

		valid, reason := isValidOrder(order, graph)
		if !valid {
			t.Errorf(
				"Topological sort produced an invalid order: %s. Full order: %v",
				reason,
				order,
			)
		}
	})
}

func TestConcurrentTopologicalSort_ErrorHandling(t *testing.T) {
	t.Run("worker returns error", func(t *testing.T) {
		graph := map[any][]any{
			"A": {"B"},
			"B": {"C"}, // B will fail
			"C": {},
		}

		expectedErr := errors.New("worker failed on B")

		var processedCount int32

		worker := func(ctx context.Context, node any) error {
			atomic.AddInt32(&processedCount, 1)

			if node.(string) == "B" {
				return expectedErr
			}

			return nil
		}

		err := topologicalsort.ConcurrentTopologicalSort(context.Background(), graph, worker)
		if err == nil {
			t.Fatal("Expected an error, but got nil")
		}

		if !strings.Contains(err.Error(), expectedErr.Error()) {
			t.Errorf("Expected error to contain '%v', but got '%v'", expectedErr, err)
		}

		// C should process, B should fail, A should not start because the error stops the chain.
		// It's possible A runs concurrently with B, so count can be 2 or 3. But it can't be all
		// nodes.
		// A better check is that not all nodes were processed successfully.
		if atomic.LoadInt32(&processedCount) == int32(len(graph)) {
			t.Errorf(
				"Expected processing to stop or be incomplete after error, but all %d nodes were processed",
				len(graph),
			)
		}
	})

	t.Run("cycle detection", func(t *testing.T) {
		// A -> B -> C -> A
		graph := map[any][]any{
			"A": {"B"},
			"B": {"C"},
			"C": {"A"},
		}

		var processedCount int32

		worker := func(ctx context.Context, node any) error {
			atomic.AddInt32(&processedCount, 1)

			return nil
		}

		err := topologicalsort.ConcurrentTopologicalSort(context.Background(), graph, worker)
		if err == nil {
			t.Fatal("Expected a cycle detection error, but got nil")
		}

		if !strings.Contains(strings.ToLower(err.Error()), "cycle detected") {
			t.Errorf("Expected error message to contain 'cycle detected', but got: %v", err)
		}

		if atomic.LoadInt32(&processedCount) > 0 {
			t.Logf(
				"Note: %d nodes were processed before cycle was detected. This is acceptable.",
				atomic.LoadInt32(&processedCount),
			)
		}
	})
}

func TestConcurrentTopologicalSort_EdgeCases(t *testing.T) {
	t.Run("empty graph", func(t *testing.T) {
		graph := map[any][]any{}
		worker := func(ctx context.Context, node any) error {
			t.Fatal("Worker should not be called for an empty graph")

			return nil
		}

		err := topologicalsort.ConcurrentTopologicalSort(context.Background(), graph, worker)
		if err != nil {
			t.Fatalf("Expected no error for an empty graph, but got: %v", err)
		}
	})

	t.Run("disconnected components", func(t *testing.T) {
		// A -> B and C -> D
		graph := map[any][]any{
			"A": {"B"},
			"B": {},
			"C": {"D"},
			"D": {},
		}

		var (
			order []string
			mu    sync.Mutex
		)

		worker := newOrderTrackingWorker(&order, &mu)

		err := topologicalsort.ConcurrentTopologicalSort(context.Background(), graph, worker)
		if err != nil {
			t.Fatalf("Expected no error, but got: %v", err)
		}

		if len(order) != 4 {
			t.Fatalf("Expected all 4 nodes to be processed, but got %d: %v", len(order), order)
		}

		valid, reason := isValidOrder(order, graph)
		if !valid {
			t.Errorf(
				"Topological sort produced an invalid order: %s. Full order: %v",
				reason,
				order,
			)
		}
	})

	t.Run("no dependencies", func(t *testing.T) {
		// A, B, C can all run in parallel
		graph := map[any][]any{
			"A": {},
			"B": {},
			"C": {},
		}

		var processedCount int32

		worker := func(ctx context.Context, node any) error {
			atomic.AddInt32(&processedCount, 1)

			return nil
		}

		err := topologicalsort.ConcurrentTopologicalSort(context.Background(), graph, worker)
		if err != nil {
			t.Fatalf("Expected no error, but got: %v", err)
		}

		if atomic.LoadInt32(&processedCount) != 3 {
			t.Fatalf(
				"Expected 3 items to be processed, but got %d",
				atomic.LoadInt32(&processedCount),
			)
		}
	})
}
