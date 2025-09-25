package heap_test

import (
	"sync"
	"testing"

	heapx "github.com/skyrocket-qy/gox/cdsa/heap"
)

func TestMinHeap(t *testing.T) {
	less := func(a, b int) bool {
		return a < b
	}

	h := heapx.New([]int{}, less)
	h.Push(3)
	h.Push(2)
	h.Push(1)

	if h.Len() != 3 {
		t.Errorf("Expected length 3, got %d", h.Len())
	}

	if h.Pop() != 1 {
		t.Errorf("Expected 1, got %d", h.Pop())
	}

	if h.Pop() != 2 {
		t.Errorf("Expected 2, got %d", h.Pop())
	}

	if h.Pop() != 3 {
		t.Errorf("Expected 3, got %d", h.Pop())
	}

	if h.Len() != 0 {
		t.Errorf("Expected length 0, got %d", h.Len())
	}
}

func TestMaxHeap(t *testing.T) {
	less := func(a, b int) bool {
		return a > b
	}

	h := heapx.New([]int{}, less)
	h.Push(1)
	h.Push(2)
	h.Push(3)

	if h.Len() != 3 {
		t.Errorf("Expected length 3, got %d", h.Len())
	}

	if h.Pop() != 3 {
		t.Errorf("Expected 3, got %d", h.Pop())
	}

	if h.Pop() != 2 {
		t.Errorf("Expected 2, got %d", h.Pop())
	}

	if h.Pop() != 1 {
		t.Errorf("Expected 1, got %d", h.Pop())
	}

	if h.Len() != 0 {
		t.Errorf("Expected length 0, got %d", h.Len())
	}
}

func TestHeapWithInitialElements(t *testing.T) {
	less := func(a, b int) bool {
		return a < b
	}

	elements := []int{5, 2, 8, 1, 9, 4}
	h := heapx.New(elements, less)

	expected := []int{1, 2, 4, 5, 8, 9}
	for _, val := range expected {
		if h.Pop() != val {
			t.Errorf("Expected %d, got %d", val, h.Pop())
		}
	}

	if h.Len() != 0 {
		t.Errorf("Expected length 0, got %d", h.Len())
	}
}

func TestConcurrentHeap(t *testing.T) {
	less := func(a, b int) bool {
		return a < b
	}
	h := heapx.New([]int{}, less)

	var wg sync.WaitGroup

	numGoroutines := 100
	numPushesPerGoroutine := 100

	// Concurrent pushes
	for i := range numGoroutines {
		wg.Add(1)

		go func(start int) {
			defer wg.Done()

			for j := range numPushesPerGoroutine {
				h.Push(start + j)
			}
		}(i * numPushesPerGoroutine)
	}

	wg.Wait()

	if h.Len() != numGoroutines*numPushesPerGoroutine {
		t.Errorf("Expected length %d, got %d", numGoroutines*numPushesPerGoroutine, h.Len())
	}

	// Concurrent pops
	poppedElements := make(chan int, numGoroutines*numPushesPerGoroutine)
	for range numGoroutines {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for range numPushesPerGoroutine {
				poppedElements <- h.Pop()
			}
		}()
	}

	wg.Wait()
	close(poppedElements)

	// Verify popped elements (order doesn't matter for min-heap, but all should be present)
	// For simplicity, we'll just check if the heap is empty and if the number of popped elements is
	// correct.
	if h.Len() != 0 {
		t.Errorf("Expected heap to be empty, got length %d", h.Len())
	}

	if len(poppedElements) != numGoroutines*numPushesPerGoroutine {
		t.Errorf(
			"Expected %d elements popped, got %d",
			numGoroutines*numPushesPerGoroutine,
			len(poppedElements),
		)
	}
}
