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
