package set

import (
	"testing"
)

func TestNewSet(t *testing.T) {
	s := NewSet[int]()
	if s.data == nil {
		t.Error("NewSet() should initialize the data map")
	}
}

func TestAdd(t *testing.T) {
	s := NewSet[int]()
	s.Add(1)
	s.Add(2)
	s.Add(1) // Add a duplicate

	if len(s.data) != 2 {
		t.Errorf("Add() failed, expected size 2, got %d", len(s.data))
	}

	if !s.Exist(1) || !s.Exist(2) {
		t.Error("Add() failed, elements not added correctly")
	}
}

func TestRemove(t *testing.T) {
	s := NewSet[int]()
	s.Add(1)
	s.Add(2)

	s.Remove(1)
	if s.Exist(1) {
		t.Error("Remove() failed, element still exists")
	}

	if len(s.data) != 1 {
		t.Errorf("Remove() failed, expected size 1, got %d", len(s.data))
	}

	// Test removing a non-existent element
	s.Remove(3)
	if len(s.data) != 1 {
		t.Errorf("Remove() of non-existent element failed, expected size 1, got %d", len(s.data))
	}
}

func TestExist(t *testing.T) {
	s := NewSet[string]()
	s.Add("a")

	if !s.Exist("a") {
		t.Error("Exist() failed, expected true, got false")
	}
	if s.Exist("b") {
		t.Error("Exist() failed, expected false, got true")
	}
}

func TestToSlice(t *testing.T) {
	s := NewSet[int]()
	s.Add(1)
	s.Add(2)
	s.Add(3)

	slice := s.ToSlice()
	if len(slice) != 3 {
		t.Errorf("ToSlice() failed, expected slice of length 3, got %d", len(slice))
	}

	// Check for existence of elements, order doesn't matter
	seen := make(map[int]bool)
	for _, v := range slice {
		seen[v] = true
	}

	if !seen[1] || !seen[2] || !seen[3] {
		t.Error("ToSlice() failed, missing elements in slice")
	}
}
