package circularqueue_test

import (
	"testing"

	"github.com/skyrocket-qy/gox/dsa/circularqueue"
)

func TestMyCircularQueue(t *testing.T) {
	q := circularqueue.Constructor(3)
	if !q.IsEmpty() {
		t.Error("New queue should be empty")
	}

	if q.IsFull() {
		t.Error("New queue should not be full")
	}

	if !q.EnQueue(1) {
		t.Error("EnQueue should succeed")
	}

	if !q.EnQueue(2) {
		t.Error("EnQueue should succeed")
	}

	if !q.EnQueue(3) {
		t.Error("EnQueue should succeed")
	}

	if q.EnQueue(4) {
		t.Error("EnQueue should fail on full queue")
	}

	if q.Rear() != 3 {
		t.Errorf("Rear failed, expected 3, got %d", q.Rear())
	}

	if q.Front() != 1 {
		t.Errorf("Front failed, expected 1, got %d", q.Front())
	}

	if !q.IsFull() {
		t.Error("Queue should be full")
	}

	if !q.DeQueue() {
		t.Error("DeQueue should succeed")
	}

	if !q.EnQueue(4) {
		t.Error("EnQueue should succeed")
	}

	if q.Rear() != 4 {
		t.Errorf("Rear failed, expected 4, got %d", q.Rear())
	}
}

func TestMyCircularQueue_Empty(t *testing.T) {
	q := circularqueue.Constructor(1)
	if !q.EnQueue(1) {
		t.Error("EnQueue should succeed")
	}

	if !q.DeQueue() {
		t.Error("DeQueue should succeed")
	}

	if !q.IsEmpty() {
		t.Error("Queue should be empty")
	}

	if q.Front() != -1 {
		t.Errorf("Front on empty queue should be -1, got %d", q.Front())
	}

	if q.Rear() != -1 {
		t.Errorf("Rear on empty queue should be -1, got %d", q.Rear())
	}
}
