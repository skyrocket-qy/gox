package segment_tree

import "src/stack"


/*Linked list*/
type Node struct {
	l, r, sum   int
	left, right *Node
}

func Build(array []int) *Node {
	var i int
	array, i = replenish(array)

	//initial leaf node
	arrayN := make([]*Node, 1<<i, 1<<i)
	for j, val := range array {
		arrayN[j] = &Node{j, j, val, nil, nil}
	}

	//recursive construct the node bottom-up
	for i > 0 {
		i >>= 1
		new := make([]*Node, i, i)
		for j, _ := range new {
			new[j] = &Node{arrayN[j<<1].l, arrayN[j<<1+1].r, arrayN[j<<1].sum + arrayN[j<<1+1].sum, arrayN[j<<1], arrayN[j<<1+1]}
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
		re_array := make([]int, n, n)
		array = append(array, re_array...)
	}
	return array, i
}

func (node *Node) Update(pos, value int) {
	stk := stack.Stack{}
	var mid int
	//find the leaf
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
	for stk.IsEmpty() == false {
		node = stk.Pop().(*Node)
		node.sum = node.left.sum + node.right.sum
	}
}

func (node *Node) GetSum(left, right int) int {
	sum := 0
	p := &sum
	node.query(left, right, p)
	return sum
}

func (node *Node) query(l, r int, p *int) {
	if l > r {
		return
	}
	if l == node.l && r == node.r {
		*p += node.sum
		return
	}
	mid := (node.l + node.r) >> 1
	node.left.query(l, mid, p)
	node.right.query(mid+1, r, p)
}
