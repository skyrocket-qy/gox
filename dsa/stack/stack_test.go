package stack_test

import (
	"testing"

	"github.com/skyrocket-qy/gox/dsa/stack"
)

func TestStack(t *testing.T) {
	stack := stack.New[int]()

	if !stack.IsEmpty() {
		t.Error("New stack should be empty")
	}

	if _, ok := stack.Peek(); ok {
		t.Error("should not be able to peek an empty stack")
	}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	if stack.IsEmpty() {
		t.Error("Stack should not be empty after push")
	}

	if val, ok := stack.Peek(); !ok || val != 3 {
		t.Errorf("peek failed, expected 3, got %d", val)
	}
	// ensure that peek does not pop the element
	if val, ok := stack.Peek(); !ok || val != 3 {
		t.Errorf("peek failed, expected 3, got %d", val)
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
