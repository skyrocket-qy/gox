package tree

import "testing"

func TestTreeQueue(t *testing.T) {
	q := &Queue{}
	if !q.IsEmpty() {
		t.Error("expected queue to be empty")
	}
	q.Push(1)
	q.Push(2)
	if q.IsEmpty() {
		t.Error("expected queue not to be empty")
	}
	if val := q.Pop(); val != 1 {
		t.Errorf("expected pop to return 1, but got %v", val)
	}
	if val := q.Pop(); val != 2 {
		t.Errorf("expected pop to return 2, but got %v", val)
	}
	if !q.IsEmpty() {
		t.Error("expected queue to be empty")
	}
}
