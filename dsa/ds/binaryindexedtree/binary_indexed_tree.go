package binaryindexedtree

// BinaryIndexedTree is a generic binary indexed tree (also known as a Fenwick tree).
// It supports efficient prefix sum queries and point updates.
// All indices for this data structure are 1-based.
type BinaryIndexedTree[T any] struct {
	tree     []T
	add      func(a, b T) T
	subtract func(a, b T) T
	zero     T
}

// New creates a new generic BinaryIndexedTree.
// length: the size of the array the BIT is based on.
// add: a function that performs addition for type T.
// subtract: a function that performs subtraction for type T.
// zero: the zero value for type T.
// Time complexity: O(n) to initialize the tree.
// Space complexity: O(n)
func New[T any](length int, add func(a, b T) T, subtract func(a, b T) T, zero T) *BinaryIndexedTree[T] {
	tree := make([]T, length+1)
	for i := range tree {
		tree[i] = zero
	}
	return &BinaryIndexedTree[T]{
		tree:     tree,
		add:      add,
		subtract: subtract,
		zero:     zero,
	}
}

// Update adds a value 'v' to the element at index 'i' (1-based).
// Time complexity: O(log n)
func (bit *BinaryIndexedTree[T]) Update(i int, v T) {
	n := len(bit.tree)
	for i < n {
		bit.tree[i] = bit.add(bit.tree[i], v)
		i += i & -i
	}
}

// Set sets the element at index 'i' (1-based) to a new value 'v'.
// Time complexity: O(log n)
func (bit *BinaryIndexedTree[T]) Set(i int, v T) {
	oldV := bit.Query(i)
	diff := bit.subtract(v, oldV)
	bit.Update(i, diff)
}

// QueryPrefixSum calculates the prefix sum up to index 'i' (1-based)
// (i.e., sum of elements from index 1 to i).
// Time complexity: O(log n)
func (bit *BinaryIndexedTree[T]) QueryPrefixSum(i int) T {
	res := bit.zero
	for i > 0 {
		res = bit.add(res, bit.tree[i])
		i -= i & -i
	}
	return res
}

// Query returns the value of the element at index 'i' (1-based).
// Time complexity: O(log n)
func (bit *BinaryIndexedTree[T]) Query(i int) T {
	return bit.subtract(bit.QueryPrefixSum(i), bit.QueryPrefixSum(i-1))
}
