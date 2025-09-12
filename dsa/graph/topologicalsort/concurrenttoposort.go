package topologicalsort

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
)

// WorkerFunc is a function that processes a single node.
type WorkerFunc func(ctx context.Context, node any) error

// ConcurrentTopologicalSort executes a worker function on each node of a directed acyclic graph
// in a topologically sorted order, concurrently.
// The graph is represented as an adjacency list where graph[node] -> dependencies.
func ConcurrentTopologicalSort(ctx context.Context, graph map[any][]any, worker WorkerFunc) error {
	// --- 1. Calculate In-Degrees (number of dependencies) ---
	inDegrees := make(map[any]int)
	allNodes := make(map[any]struct{})

	for node, deps := range graph {
		allNodes[node] = struct{}{}
		inDegrees[node] = len(deps) // Node's in-degree is the count of its dependencies
		for _, dep := range deps {
			allNodes[dep] = struct{}{} // Ensure all dependencies are also in allNodes
		}
	}

	// --- 2. Initialize Queue with Zero In-Degree Nodes ---
	// These are the starting points (leaves of the dependency graph).
	readyCh := make(chan any, len(allNodes))
	for node := range allNodes {
		if inDegrees[node] == 0 {
			readyCh <- node
		}
	}

	// --- 3. Process Nodes Concurrently ---
	var (
		wg        sync.WaitGroup
		mu        sync.Mutex // Protects inDegrees map
		firstErr  atomic.Value
		processed int64
	)

	// We start with the number of items initially in the queue.
	// The WaitGroup will be incremented as new items become ready.
	wg.Add(len(readyCh))

	go func() {
		wg.Wait()
		close(readyCh)
	}()

	for node := range readyCh {
		go func(node any) {
			defer wg.Done()

			// If an error has already occurred, stop processing new items.
			if firstErr.Load() != nil {
				return
			}

			// Execute the specific work for this node.
			if err := worker(ctx, node); err != nil {
				firstErr.Store(err)
				return
			}

			atomic.AddInt64(&processed, 1)

			// --- 4. Update In-Degrees of Neighbors ---
			// Since this node is done, decrement the in-degree of the nodes
			// that depended on it.
			for dependent, deps := range graph {
				for _, dep := range deps {
					if dep == node {
						mu.Lock()
						inDegrees[dependent]--
						isReady := inDegrees[dependent] == 0
						mu.Unlock()

						if isReady {
							wg.Add(1)
							readyCh <- dependent
						}
					}
				}
			}
		}(node)
	}

	// --- 5. Final Error and Cycle Check ---
	if loadedErr := firstErr.Load(); loadedErr != nil {
		if e, ok := loadedErr.(error); ok {
			return e
		}
		// If for some reason it's not an error, return a standard Go error.
		return fmt.Errorf("shutdown failed: unexpected error type: %v", loadedErr)
	}

	// If not all nodes were processed, there must be a cycle in the graph.
	if int(processed) != len(allNodes) {
		return errors.New("cycle detected in dependency graph")
	}

	return nil
}
