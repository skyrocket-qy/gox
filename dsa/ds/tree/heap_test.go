package tree

import "testing"

func TestMinHeap(t *testing.T) {
	list := []int{11, 2, 10, 4, 5, 7, 9, 8, 6, 1}

	mh := MinHeap{}
	mh.Build()

	for _, val := range list {
		mh.Put(val)
	}

	// After building the heap, the order is not guaranteed to be sorted,
	// but it must satisfy the heap property.
	// The smallest element should be at the root.
	if mh.Get() != 1 {
		t.Error("expected to get 1 from heap")
	}
	if mh.Get() != 2 {
		t.Error("expected to get 2 from heap")
	}
	if mh.Get() != 4 {
		t.Error("expected to get 4 from heap")
	}
}
