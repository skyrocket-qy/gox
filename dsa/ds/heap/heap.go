package heap

// Heap is a min-heap or max-heap implementation using generics.
// It is 1-indexed, meaning the first element (index 0) is a dummy value.
type Heap[T any] struct {
	data []T
	less func(T, T) bool // less(a, b) returns true if a < b (for min-heap) or a > b (for max-heap)
}

// New creates and initializes a new Heap with the given elements and a less function.
// The less function determines the heap's order:
// - For a min-heap, less(a, b) should return true if a < b.
// - For a max-heap, less(a, b) should return true if a > b.
func New[T any](eles []T, less func(T, T) bool) *Heap[T] {
	h := &Heap[T]{
		// 1. Create the slice with length 1 (for the dummy element)
		//    and capacity for all future elements.
		data: make([]T, 1, len(eles)+1),
		less: less,
	}

	// 2. Append the initial elements. This correctly sets the slice's length.
	h.data = append(h.data, eles...)

	// 3. The heapify logic remains the same and is now safe to run.
	// h.Len() correctly returns n here.
	for i := h.Len() / 2; i > 0; i-- {
		h.down(i)
	}

	return h
}

// Len returns the number of elements in the heap.
func (h *Heap[T]) Len() int { return len(h.data) - 1 }

func (h *Heap[T]) swap(i, j int) { h.data[i], h.data[j] = h.data[j], h.data[i] }

// Push adds an element to the heap.
func (h *Heap[T]) Push(e T) {
	h.data = append(h.data, e)
	h.up(h.Len())
}

// Pop removes and returns the top element from the heap.
func (h *Heap[T]) Pop() T {
	if h.Len() == 0 {
		var zeroValue T
		return zeroValue // Or panic, depending on desired behavior for empty heap
	}

	n := h.Len()
	res := h.data[1]
	h.data[1] = h.data[n]
	h.data = h.data[:n] // Shrink the slice
	h.down(1)
	return res
}

func (h *Heap[T]) up(i int) {
	for j := i >> 1; i > 1 && h.less(h.data[i], h.data[j]); i, j = j, j>>1 {
		h.swap(i, j)
	}
}

func (h *Heap[T]) down(i int) {
	for j := i << 1; j <= h.Len(); i, j = j, j<<1 {
		if j+1 <= h.Len() && h.less(h.data[j+1], h.data[j]) {
			j++
		}
		if h.less(h.data[j], h.data[i]) {
			h.swap(i, j)
		} else {
			break
		}
	}
}
