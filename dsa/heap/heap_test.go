package heap_test

import (
	"container/heap"
	"math/rand"
	"testing"

	heapx "github.com/skyrocket-qy/gox/dsa/heap"
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

type IntHeap []int

func (h *IntHeap) Len() int           { return len(*h) }
func (h *IntHeap) Less(i, j int) bool { return (*h)[i] < (*h)[j] }
func (h *IntHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *IntHeap) Push(x any) {
	intVal, ok := x.(int)
	if !ok {
		panic("Push received a non-int value") // Should not happen in this test context
	}

	*h = append(*h, intVal)
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

// ============================================================================
// 3. Benchmark Setup
// ============================================================================

const benchmarkSize = 100000

var testData []int

// generateData creates a slice of random integers.
// It's called once to ensure all benchmarks run on the exact same data.
func init() {
	// #nosec G404 -- math/rand is acceptable for test data generation, not security-sensitive.
	rng := rand.New(rand.NewSource(42)) // Fixed seed for reproducible benchmarks

	testData = make([]int, benchmarkSize)
	for i := range benchmarkSize {
		testData[i] = rng.Int()
	}
}

// less function for both heaps.
var lessFunc = func(a, b int) bool { return a < b }

// ============================================================================
// 4. Benchmark Functions
// ============================================================================

// --- Initialization Benchmarks ---
// These test creating a heap from an existing slice of N elements.
// This is where we expect to see the biggest difference.

func BenchmarkCustomHeap_Init(b *testing.B) {
	for range b.N {
		_ = heapx.New(testData, lessFunc)
	}
}

func BenchmarkStdHeap_Init(b *testing.B) {
	for range b.N {
		dataCopy := make(IntHeap, benchmarkSize)
		copy(dataCopy, testData)
		heap.Init(&dataCopy)
	}
}

// --- Push Benchmarks ---
// These test adding one element to a full heap.

func BenchmarkCustomHeap_Push(b *testing.B) {
	h := heapx.New(testData, lessFunc)

	b.ResetTimer() // Start timing after the heap is already built

	for i := range b.N {
		h.Push(i)
	}
}

func BenchmarkStdHeap_Push(b *testing.B) {
	dataCopy := make(IntHeap, benchmarkSize)
	copy(dataCopy, testData)
	h := &dataCopy
	heap.Init(h)
	b.ResetTimer() // Start timing after the heap is already built

	for i := range b.N {
		heap.Push(h, i)
	}
}

// --- Pop Benchmarks ---
// These test removing the top element from a full heap.

func BenchmarkCustomHeap_Pop(b *testing.B) {
	h := heapx.New(testData, lessFunc)

	b.ResetTimer()

	for i := range b.N {
		// To prevent the heap from emptying, we push an element back.
		// This keeps the benchmark focused on the Pop() operation
		// on a consistently-sized heap.
		h.Pop()
		h.Push(i)
	}
}

func BenchmarkStdHeap_Pop(b *testing.B) {
	dataCopy := make(IntHeap, benchmarkSize)
	copy(dataCopy, testData)
	h := &dataCopy
	heap.Init(h)
	b.ResetTimer()

	for i := range b.N {
		heap.Pop(h)
		heap.Push(h, i)
	}
}
