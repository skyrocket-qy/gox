package binaryindexedtree

/* @tags: tree,bit operation,prefix sum */

/*
Use: Compute dynamic prefix sum
Complexity:
	Time: O(logN)
	Space: O(N)
*/

type binaryIndexedTree []int

func NewBinaryIndexedTree(length int) binaryIndexedTree {
	return make(binaryIndexedTree, length+1)
}

// Set A[i] += v
func (t binaryIndexedTree) Update(i, v int) {
	n := len(t)
	for i <= n {
		t[i] += v
		i += i & -i
	}
}

// Set A[i] = v
func (t binaryIndexedTree) Set(i, v int) {
	oldV := t.Query(i)
	diff := v - oldV
	t.Update(i, diff)
}

// A[1] + ... + A[i]
func (t binaryIndexedTree) QueryPrefixSum(i int) int {
	res := 0
	for i > 0 {
		res += t[i]
		i -= i & -i
	}
	return res
}

// A[i] = PreSum[i] - PreSum[i-1] = B[i] -
func (t binaryIndexedTree) Query(i int) int {
	if i&1 == 1 {
		return t[i]
	}
	if (i-2)%4 == 0 {
		return t[i] - t[i-1]
	}
	if (i-4)%8 == 0 {
		return t[i] - t[i-1] - t[i-2]
	}
	if (i-8)%16 == 0 {
		return t[i] - t[i-1] - t[i-2] - t[i-4]
	}
	return t.QueryPrefixSum(i) - t.QueryPrefixSum(i-1)
}
