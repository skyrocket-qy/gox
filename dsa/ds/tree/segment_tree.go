package tree

import "github.com/skyrocket-qy/gox/dsa/ds/stack"

/*Linked list.*/
type Node struct {
	l, r, sum   int
	left, right *Node
}

func NewNode(array []int) *Node {
	if len(array) == 0 {
		return nil
	}

	var i int

	array, i = replenish(array)

	// initial leaf node
	arrayN := make([]*Node, 1<<i)
	for j, val := range array {
		arrayN[j] = &Node{j, j, val, nil, nil}
	}

	// recursive construct the node bottom-up
	for i > 0 {
		i--

		new := make([]*Node, 1<<i)
		for j := range new {
			new[j] = &Node{
				arrayN[j<<1].l,
				arrayN[j<<1+1].r,
				arrayN[j<<1].sum + arrayN[j<<1+1].sum,
				arrayN[j<<1],
				arrayN[j<<1+1],
			}
		}

		arrayN = new
	}

	return arrayN[0]
}

func replenish(array []int) ([]int, int) {
	i := 0
	for n := len(array); n > 1; n >>= 1 {
		i++
	}

	if 1<<i < len(array) {
		i++
		n := 1<<i - len(array)
		re_array := make([]int, n)
		array = append(array, re_array...)
	}

	return array, i
}

func (node *Node) Update(pos, value int) {
	if node == nil {
		return
	}

	stk := stack.NewStack[*Node]()

	var mid int
	// find the leaf
	for pos != node.l || pos != node.r {
		stk.Push(node)

		mid = (node.l + node.r) >> 1
		if pos > mid {
			node = node.right
		} else {
			node = node.left
		}
	}

	node.sum = value

	for !stk.IsEmpty() {
		val, _ := stk.Pop()
		node = val

		node.sum = node.left.sum + node.right.sum
	}
}

func (node *Node) GetSum(left, right int) int {
	if node == nil {
		return 0
	}

	sum := 0
	p := &sum
	node.query(left, right, p)

	return sum
}

func (node *Node) query(l, r int, p *int) {
	// If the node's range is completely outside the query range, do nothing.
	if node == nil || l > node.r || r < node.l {
		return
	}

	// If the node's range is completely inside the query range, add its sum.
	if l <= node.l && node.r <= r {
		*p += node.sum

		return
	}

	// Partially overlapping, recurse on children
	node.left.query(l, r, p)
	node.right.query(l, r, p)
}
