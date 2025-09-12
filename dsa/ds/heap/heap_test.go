package heap

import (
	"testing"
)

type Int int

func (i Int) Less(e Element) bool {
	return i < e.(Int)
}

func TestHeap(t *testing.T) {
	// Test Push and Pop
	h := &Heap{Int(0)} // Initialize with a dummy element at index 0
	h.Push(Int(3))
	h.Push(Int(1))
	h.Push(Int(4))
	h.Push(Int(1))
	h.Push(Int(5))
	h.Push(Int(9))

	expectedOrder := []Int{9, 5, 4, 3, 1, 1}
	for _, expected := range expectedOrder {
		popped := h.Pop().(Int)
		if popped != expected {
			t.Errorf("Pop() = %v, want %v", popped, expected)
		}
	}

	// Test Init
	elements := []Element{Int(0), Int(3), Int(1), Int(4), Int(1), Int(5), Int(9)}
	h = Init(elements)
	for _, expected := range expectedOrder {
		popped := h.Pop().(Int)
		if popped != expected {
			t.Errorf("Pop() after Init() = %v, want %v", popped, expected)
		}
	}
}
