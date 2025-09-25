package binaryindexedtree

import "golang.org/x/exp/constraints"

/* @tags: tree,bit operation,prefix sum */

type Number interface {
	constraints.Integer | constraints.Float
}

// BinaryIndexedTree is a generic binary indexed tree (also known as a Fenwick tree).
// It supports efficient prefix sum queries and point updates for numeric types.
// All indices for this data structure are 1-based.
type BinaryIndexedTree[T Number] struct {
	tree []T
}

// New creates a new generic BinaryIndexedTree for any numeric type.
// length: the size of the array the BIT is based on.
// Time complexity: O(n) to initialize the tree.
// Space complexity: O(n)
func New[T Number](length int) *BinaryIndexedTree[T] {
	return &BinaryIndexedTree[T]{
		tree: make([]T, length+1),
	}
}

// Update adds a value 'v' to the element at index 'i' (1-based).
// Time complexity: O(log n)
func (bit *BinaryIndexedTree[T]) Update(i int, v T) {
	n := len(bit.tree)
	for i < n {
		bit.tree[i] += v
		i += i & -i
	}
}

// Set sets the element at index 'i' (1-based) to a new value 'v'.
// Time complexity: O(log n)
func (bit *BinaryIndexedTree[T]) Set(i int, v T) {
	oldV := bit.Query(i)
	diff := v - oldV
	bit.Update(i, diff)
}

// QueryPrefixSum calculates the prefix sum up to index 'i' (1-based)
// (i.e., sum of elements from index 1 to i).
// Time complexity: O(log n)
func (bit *BinaryIndexedTree[T]) QueryPrefixSum(i int) T {
	var res T
	for i > 0 {
		res += bit.tree[i]
		i -= i & -i
	}
	return res
}

// Query returns the value of the element at index 'i' (1-based).
// Time complexity: O(log n)
func (bit *BinaryIndexedTree[T]) Query(i int) T {
	return bit.QueryPrefixSum(i) - bit.QueryPrefixSum(i-1)
}
