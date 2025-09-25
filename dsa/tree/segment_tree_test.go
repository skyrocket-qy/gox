package tree_test

import (
	"testing"

	"github.com/skyrocket-qy/gox/dsa/tree"
)

func TestSegmentTree(t *testing.T) {
	array := []int{1, 3, 5, 7, 9, 11}
	root := tree.NewNode(array)

	// Test GetSum before any updates
	if sum := root.GetSum(0, 5); sum != 36 {
		t.Errorf("Initial GetSum(0, 5) failed, expected 36, got %d", sum)
	}

	if sum := root.GetSum(0, 2); sum != 9 {
		t.Errorf("Initial GetSum(0, 2) failed, expected 9, got %d", sum)
	}

	if sum := root.GetSum(3, 5); sum != 27 {
		t.Errorf("Initial GetSum(3, 5) failed, expected 27, got %d", sum)
	}

	if sum := root.GetSum(2, 4); sum != 21 { // 5 + 7 + 9 = 21
		t.Errorf("Initial GetSum(2, 4) failed, expected 21, got %d", sum)
	}

	// Test Update
	root.Update(1, 4) // Update array[1] from 3 to 4. New array: [1, 4, 5, 7, 9, 11]

	// Test GetSum after update
	if sum := root.GetSum(0, 5); sum != 37 { // 1+4+5+7+9+11 = 37
		t.Errorf("GetSum(0, 5) after update failed, expected 37, got %d", sum)
	}

	if sum := root.GetSum(0, 2); sum != 10 { // 1 + 4 + 5 = 10
		t.Errorf("GetSum(0, 2) after update failed, expected 10, got %d", sum)
	}

	// Test another update
	root.Update(5, 1) // Update array[5] from 11 to 1. New array: [1, 4, 5, 7, 9, 1]

	if sum := root.GetSum(0, 5); sum != 27 { // 1 + 4 + 5 + 7 + 9 + 1 = 27
		t.Errorf("GetSum(0, 5) after second update failed, expected 27, got %d", sum)
	}

	if sum := root.GetSum(4, 5); sum != 10 { // 9 + 1 = 10
		t.Errorf("GetSum(4, 5) after second update failed, expected 10, got %d", sum)
	}

	// Test update on edge
	root.Update(0, 2) // New array: [2, 4, 5, 7, 9, 1]

	if sum := root.GetSum(0, 0); sum != 2 {
		t.Errorf("GetSum(0,0) after update failed, expected 2, got %d", sum)
	}

	// Test out-of-bounds query
	if sum := root.GetSum(100, 200); sum != 0 {
		t.Errorf("GetSum for out-of-bounds range failed, expected 0, got %d", sum)
	}
}

func TestSegmentTree_Empty(t *testing.T) {
	var array []int

	root := tree.NewNode(array)

	if root != nil {
		t.Errorf("NewNode with empty array should return nil, got %v", root)
	}

	if sum := root.GetSum(0, 5); sum != 0 {
		t.Errorf("GetSum on nil tree failed, expected 0, got %d", sum)
	}

	// Should not panic
	root.Update(0, 10)
}
