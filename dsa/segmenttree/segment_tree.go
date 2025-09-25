package segmenttree

import (
	"fmt"
)

// SegmentTree represents a segment tree.
// It is a versatile data structure that can be used to answer range queries.
// The type of query it can answer depends on the merge function provided.
// For example, if the merge function is `func(a, b int) int { return a + b }`,
// then the segment tree can be used to find the sum of a range.
// If the merge function is `func(a, b int) int { if a < b { return a } else { return b } }`,
// then the segment tree can be used to find the minimum of a range.
type SegmentTree[T any] struct {
	data  []T
	tree  []T
	merge func(a, b T) T
}

// New creates a new SegmentTree with the given data and merge function.
// The merge function is used to combine two elements in the tree.
// Time complexity: O(n)
// Space complexity: O(n).
func New[T any](data []T, merge func(a, b T) T) *SegmentTree[T] {
	if len(data) == 0 {
		return nil
	}

	size := nextPowerOf2(len(data))
	treeSize := 2*size - 1

	st := &SegmentTree[T]{
		data:  make([]T, len(data)),
		tree:  make([]T, treeSize), // Use the calculated, optimized size
		merge: merge,
	}
	copy(st.data, data)

	st.build(0, 0, len(data)-1)

	return st
}

// nextPowerOf2 calculates the next power of 2 for a given integer n.
// This helps in calculating the minimum required size for the segment tree array.
func nextPowerOf2(n int) int {
	if n <= 0 {
		return 1
	}

	p := 1
	for p < n {
		p <<= 1
	}

	return p
}

// build is a private helper function that builds the segment tree recursively.
func (st *SegmentTree[T]) build(treeIndex, l, r int) {
	if l == r {
		st.tree[treeIndex] = st.data[l]

		return
	}

	mid := l + (r-l)/2
	leftTreeIndex := 2*treeIndex + 1
	rightTreeIndex := 2*treeIndex + 2

	st.build(leftTreeIndex, l, mid)
	st.build(rightTreeIndex, mid+1, r)

	st.tree[treeIndex] = st.merge(st.tree[leftTreeIndex], st.tree[rightTreeIndex])
}

// Query performs a query on the range [queryL, queryR].
// It returns the result of the query.
// If the query range is invalid, it returns the zero value of type T.
// Time complexity: O(log n)
// Space complexity: O(log n) for recursion stack.
func (st *SegmentTree[T]) Query(queryL, queryR int) T {
	var zero T
	if st == nil || queryL < 0 || queryR >= len(st.data) || queryL > queryR {
		return zero
	}

	return st.query(0, 0, len(st.data)-1, queryL, queryR)
}

// query is a private helper function that performs the query recursively.
func (st *SegmentTree[T]) query(treeIndex, l, r, queryL, queryR int) T {
	if l == queryL && r == queryR {
		return st.tree[treeIndex]
	}

	mid := l + (r-l)/2
	leftTreeIndex := 2*treeIndex + 1
	rightTreeIndex := 2*treeIndex + 2

	if queryR <= mid {
		return st.query(leftTreeIndex, l, mid, queryL, queryR)
	} else if queryL > mid {
		return st.query(rightTreeIndex, mid+1, r, queryL, queryR)
	}

	leftResult := st.query(leftTreeIndex, l, mid, queryL, mid)
	rightResult := st.query(rightTreeIndex, mid+1, r, mid+1, queryR)

	return st.merge(leftResult, rightResult)
}

// Update updates the value at the given index with the given value.
// Time complexity: O(log n)
// Space complexity: O(log n) for recursion stack.
func (st *SegmentTree[T]) Update(index int, val T) {
	if st == nil || index < 0 || index >= len(st.data) {
		return
	}

	st.data[index] = val
	st.update(0, 0, len(st.data)-1, index, val)
}

// update is a private helper function that performs the update recursively.
func (st *SegmentTree[T]) update(treeIndex, l, r, index int, val T) {
	if l == r {
		st.tree[treeIndex] = val

		return
	}

	mid := l + (r-l)/2
	leftTreeIndex := 2*treeIndex + 1
	rightTreeIndex := 2*treeIndex + 2

	if index <= mid {
		st.update(leftTreeIndex, l, mid, index, val)
	} else {
		st.update(rightTreeIndex, mid+1, r, index, val)
	}

	st.tree[treeIndex] = st.merge(st.tree[leftTreeIndex], st.tree[rightTreeIndex])
}

// String returns a string representation of the segment tree.
func (st *SegmentTree[T]) String() string {
	if st == nil {
		return "<nil>"
	}

	return fmt.Sprintf("SegmentTree with data: %v", st.data)
}

// QueryLeftmostIndex finds the index of the leftmost basket with capacity >= requiredCapacity.
// This is a specialized query, assuming the tree stores integers and the merge function
// calculates the maximum in a range.
// Returns -1 if no such basket is found.
// It will panic if the tree does not store integers.
// Time complexity: O(log n)
// Space complexity: O(log n) for recursion stack.
func (st *SegmentTree[T]) QueryLeftmostIndex(requiredCapacity int) int {
	if st == nil {
		return -1
	}

	// Check if the max of the whole range is sufficient.
	if v, ok := any(st.tree[0]).(int); !ok || v < requiredCapacity {
		return -1
	}

	return st.queryLeftmostIndex(0, 0, len(st.data)-1, requiredCapacity)
}

func (st *SegmentTree[T]) queryLeftmostIndex(treeIndex, l, r, requiredCapacity int) int {
	// The max in this range is not enough.
	if v, ok := any(st.tree[treeIndex]).(int); !ok || v < requiredCapacity {
		return -1
	}

	// We are at a leaf node, and we know its value is >= requiredCapacity.
	if l == r {
		return l
	}

	mid := l + (r-l)/2
	leftTreeIndex := 2*treeIndex + 1
	rightTreeIndex := 2*treeIndex + 2

	// Try to find in the left subtree first.
	res := st.queryLeftmostIndex(leftTreeIndex, l, mid, requiredCapacity)
	if res != -1 {
		return res
	}

	// If not in the left, try the right subtree.
	return st.queryLeftmostIndex(rightTreeIndex, mid+1, r, requiredCapacity)
}
