package custom

import (
	"testing"
)

func TestCustomPriorityQueue(t *testing.T) {
	// Create a priority queue
	pq := NewPriorityQueue()

	// Push items into the priority queue
	items := []Node{
		{Priority: 3},
		{Priority: 1},
		{Priority: 4},
		{Priority: 2},
	}

	for _, item := range items {
		pq.Push(item)
	}

	// Test Pop in order of priority
	expectedPriorities := []uint{1, 2, 3, 4}
	for i, priority := range expectedPriorities {
		item := pq.Pop()
		if item.Priority != priority {
			t.Errorf("Expected item %d to have priority %d, but got %d", i, priority, item.Priority)
		}
	}
}

func TestPopFromEmptyQueue(t *testing.T) {
	pq := NewPriorityQueue()
	// Test Pop from an empty queue
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic when popping from an empty queue")
		}
	}()

	pq.Pop()
}

func TestDown(t *testing.T) {
	// Manually create a queue to test the down function
	pq := NewPriorityQueue()
	pq = append(pq, Node{Priority: 10})
	pq = append(pq, Node{Priority: 5})
	pq = append(pq, Node{Priority: 6})
	pq.down(1)

	if pq[1].Priority != 5 {
		t.Errorf("down failed, expected 5, got %d", pq[1].Priority)
	}

	pq = NewPriorityQueue()
	pq = append(pq, Node{Priority: 10})
	pq = append(pq, Node{Priority: 6})
	pq = append(pq, Node{Priority: 5})
	pq.down(1)

	if pq[1].Priority != 5 {
		t.Errorf("down failed, expected 5, got %d", pq[1].Priority)
	}

	pq = NewPriorityQueue()
	pq = append(pq, Node{Priority: 1})
	pq = append(pq, Node{Priority: 6})
	pq = append(pq, Node{Priority: 5})
	pq.down(1)

	if pq[1].Priority != 1 {
		t.Errorf("down failed, expected 1, got %d", pq[1].Priority)
	}
}
