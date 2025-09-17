package lib

import (
	"container/heap"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	// Create a priority queue
	pq := NewPriorityQueue()

	// Test Len on an empty queue
	if pq.Len() != 0 {
		t.Errorf("Expected length of empty queue to be 0, but got %d", pq.Len())
	}

	// Push items into the priority queue
	items := []Node{
		{Priority: 3},
		{Priority: 1},
		{Priority: 4},
		{Priority: 2},
	}

	for _, item := range items {
		heap.Push(&pq, item)
	}

	// Test Len after pushing items
	if pq.Len() != len(items) {
		t.Errorf("Expected length of queue to be %d, but got %d", len(items), pq.Len())
	}

	// Test Pop in order of priority
	expectedPriorities := []uint{1, 2, 3, 4}
	for i, priority := range expectedPriorities {
		item := heap.Pop(&pq).(Node)
		if item.Priority != priority {
			t.Errorf("Expected item %d to have priority %d, but got %d", i, priority, item.Priority)
		}
	}

	// Test Len after popping all items
	if pq.Len() != 0 {
		t.Errorf("Expected length of queue to be 0 after popping all items, but got %d", pq.Len())
	}

	// Test Pop from an empty queue
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic when popping from an empty queue")
		}
	}()

	_ = heap.Pop(&pq)
}
