package tree

import "testing"

func TestTreeStack(t *testing.T) {
	s := &Stack{}
	if !s.IsEmpty() {
		t.Error("expected stack to be empty")
	}

	s.Push(1)
	s.Push(2)

	if s.IsEmpty() {
		t.Error("expected stack not to be empty")
	}

	if val := s.Pop(); val != 2 {
		t.Errorf("expected pop to return 2, but got %v", val)
	}

	if val := s.Pop(); val != 1 {
		t.Errorf("expected pop to return 1, but got %v", val)
	}

	if !s.IsEmpty() {
		t.Error("expected stack to be empty")
	}
}
