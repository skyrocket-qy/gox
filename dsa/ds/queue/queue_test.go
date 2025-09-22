package queue_test

import (
	"testing"

	"github.com/skyrocket-qy/gox/dsa/ds/queue"
)

func TestQueue(t *testing.T) {
	queue := queue.New[int]()

	if !queue.IsEmpty() {
		t.Error("New queue should be empty")
	}

	queue.Push(1)
	queue.Push(2)
	queue.Push(3)

	if queue.IsEmpty() {
		t.Error("Queue should not be empty after push")
	}

	val, ok := queue.Pop()
	if !ok || val != 1 {
		t.Errorf("Pop failed, expected 1, got %d", val)
	}

	val, ok = queue.Pop()
	if !ok || val != 2 {
		t.Errorf("Pop failed, expected 2, got %d", val)
	}

	val, ok = queue.Pop()
	if !ok || val != 3 {
		t.Errorf("Pop failed, expected 3, got %d", val)
	}

	if !queue.IsEmpty() {
		t.Error("Queue should be empty after popping all elements")
	}

	_, ok = queue.Pop()
	if ok {
		t.Error("Pop on empty queue should return false")
	}
}
