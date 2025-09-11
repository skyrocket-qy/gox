package queue

import "testing"

func TestQueueLinkedList(t *testing.T) {
	q := &QueueLinkedList{}

	if !q.IsEmpty() {
		t.Error("New queue should be empty")
	}

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	if q.Size() != 3 {
		t.Errorf("Expected queue size to be 3, but got %d", q.Size())
	}

	if q.IsEmpty() {
		t.Error("Queue should not be empty after enqueue")
	}

	val := q.Dequeue()
	if val != 1 {
		t.Errorf("Dequeue failed, expected 1, got %d", val)
	}

	val = q.Dequeue()
	if val != 2 {
		t.Errorf("Dequeue failed, expected 2, got %d", val)
	}

	val = q.Dequeue()
	if val != 3 {
		t.Errorf("Dequeue failed, expected 3, got %d", val)
	}

	if !q.IsEmpty() {
		t.Error("Queue should be empty after popping all elements")
	}

	val = q.Dequeue()
	if val != nil {
		t.Errorf("Dequeue on empty queue should return nil, got %v", val)
	}
}
