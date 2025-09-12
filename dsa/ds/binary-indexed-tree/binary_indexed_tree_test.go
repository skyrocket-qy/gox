package binaryindexedtree

import "testing"

func TestBinaryIndexedTree(t *testing.T) {
	bit := NewBinaryIndexedTree(10)
	bit.Update(1, 1)
	bit.Update(2, 2)
	bit.Update(3, 3)

	if bit.QueryPrefixSum(1) != 1 {
		t.Errorf("QueryPrefixSum(1) failed, got %d, want 1", bit.QueryPrefixSum(1))
	}

	if bit.QueryPrefixSum(2) != 3 {
		t.Errorf("QueryPrefixSum(2) failed, got %d, want 3", bit.QueryPrefixSum(2))
	}

	if bit.QueryPrefixSum(3) != 6 {
		t.Errorf("QueryPrefixSum(3) failed, got %d, want 6", bit.QueryPrefixSum(3))
	}

	if bit.Query(1) != 1 {
		t.Errorf("Query(1) failed, got %d, want 1", bit.Query(1))
	}

	if bit.Query(2) != 2 {
		t.Errorf("Query(2) failed, got %d, want 2", bit.Query(2))
	}

	if bit.Query(3) != 3 {
		t.Errorf("Query(3) failed, got %d, want 3", bit.Query(3))
	}

	bit.Set(2, 5)

	if bit.Query(2) != 5 {
		t.Errorf("Set(2, 5) failed, Query(2) got %d, want 5", bit.Query(2))
	}

	if bit.QueryPrefixSum(2) != 6 {
		t.Errorf("QueryPrefixSum(2) after Set(2,5) failed, got %d, want 6", bit.QueryPrefixSum(2))
	}
}

func TestBinaryIndexedTreeQuery(t *testing.T) {
	bit := NewBinaryIndexedTree(20)
	for i := 1; i <= 20; i++ {
		bit.Update(i, i)
	}

	// Test case for (i-4)%8 == 0
	if bit.Query(4) != 4 {
		t.Errorf("Query(4) failed, got %d, want 4", bit.Query(4))
	}

	// Test case for (i-8)%16 == 0
	if bit.Query(8) != 8 {
		t.Errorf("Query(8) failed, got %d, want 8", bit.Query(8))
	}

	// Test case for the final return statement
	if bit.Query(6) != 6 {
		t.Errorf("Query(6) failed, got %d, want 6", bit.Query(6))
	}
}

func TestBinaryIndexedTreeSetToZero(t *testing.T) {
	bit := NewBinaryIndexedTree(10)
	bit.Update(1, 1)
	bit.Set(1, 0)
	if bit.Query(1) != 0 {
		t.Errorf("Set(1, 0) failed, Query(1) got %d, want 0", bit.Query(1))
	}
	if bit.QueryPrefixSum(1) != 0 {
		t.Errorf("QueryPrefixSum(1) after Set(1,0) failed, got %d, want 0", bit.QueryPrefixSum(1))
	}
}
