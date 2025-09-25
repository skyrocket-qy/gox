package binaryindexedtree

import (
	"testing"
)

func TestBinaryIndexedTree_Int(t *testing.T) {
	length := 10

	bit := New[int](length)

	// Initial array: [0, 0, 0, 0, 0, 0, 0, 0, 0, 0]

	// Update index 3 (1-based) with value 5
	// Array becomes: [0, 0, 5, 0, 0, 0, 0, 0, 0, 0]
	bit.Update(3, 5)

	// Query prefix sum up to 5
	// Sum(1..5) = 5
	if sum := bit.QueryPrefixSum(5); sum != 5 {
		t.Errorf("QueryPrefixSum(5) = %d; want 5", sum)
	}
	// Query prefix sum up to 2
	// Sum(1..2) = 0
	if sum := bit.QueryPrefixSum(2); sum != 0 {
		t.Errorf("QueryPrefixSum(2) = %d; want 0", sum)
	}

	// Update index 5 with value 2
	// Array becomes: [0, 0, 5, 0, 2, 0, 0, 0, 0, 0]
	bit.Update(5, 2)

	// Query prefix sum up to 5
	// Sum(1..5) = 5 + 2 = 7
	if sum := bit.QueryPrefixSum(5); sum != 7 {
		t.Errorf("QueryPrefixSum(5) = %d; want 7", sum)
	}

	// Query value at index 3
	if val := bit.Query(3); val != 5 {
		t.Errorf("Query(3) = %d; want 5", val)
	}

	// Query value at index 4
	if val := bit.Query(4); val != 0 {
		t.Errorf("Query(4) = %d; want 0", val)
	}

	// Set value at index 3 to 1
	// Array becomes: [0, 0, 1, 0, 2, 0, 0, 0, 0, 0]
	bit.Set(3, 1)

	// Query prefix sum up to 5
	// Sum(1..5) = 1 + 2 = 3
	if sum := bit.QueryPrefixSum(5); sum != 3 {
		t.Errorf("After Set, QueryPrefixSum(5) = %d; want 3", sum)
	}

	// Query value at index 3
	if val := bit.Query(3); val != 1 {
		t.Errorf("After Set, Query(3) = %d; want 1", val)
	}
}

func TestBinaryIndexedTree_Float64(t *testing.T) {
	length := 10

	bit := New[float64](length)

	bit.Update(3, 5.5)
	bit.Update(5, 2.2)

	if sum := bit.QueryPrefixSum(5); sum != 7.7 {
		t.Errorf("QueryPrefixSum(5) = %f; want 7.7", sum)
	}

	bit.Set(3, 1.1)

	// Expected sum is 1.1 + 2.2 = 3.3
	if sum := bit.QueryPrefixSum(5); sum > 3.3+1e-9 || sum < 3.3-1e-9 {
		t.Errorf("After Set, QueryPrefixSum(5) = %f; want 3.3", sum)
	}
}