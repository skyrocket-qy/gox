package stack

import "testing"

func TestStack(t *testing.T) {
	stack := NewStack[int]()

	if !stack.IsEmpty() {
		t.Error("New stack should be empty")
	}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	if stack.IsEmpty() {
		t.Error("Stack should not be empty after push")
	}

	val, ok := stack.Pop()
	if !ok || val != 3 {
		t.Errorf("Pop failed, expected 3, got %d", val)
	}

	val, ok = stack.Pop()
	if !ok || val != 2 {
		t.Errorf("Pop failed, expected 2, got %d", val)
	}

	val, ok = stack.Pop()
	if !ok || val != 1 {
		t.Errorf("Pop failed, expected 1, got %d", val)
	}

	if !stack.IsEmpty() {
		t.Error("Stack should be empty after popping all elements")
	}

	_, ok = stack.Pop()
	if ok {
		t.Error("Pop on empty stack should return false")
	}
}
