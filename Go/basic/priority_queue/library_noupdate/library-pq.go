package lib

type Node struct {
	Priority uint
}

type PriorityQueue []Node

func NewPriorityQueue() PriorityQueue { return PriorityQueue{} }

func (q PriorityQueue) Len() int { return len(q) }

func (q PriorityQueue) Less(i, j int) bool { return q[i].Priority < q[j].Priority }

func (q PriorityQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (h *PriorityQueue) Push(x interface{}) {
	*h = append(*h, x.(Node))
}

func (h *PriorityQueue) Pop() interface{} {
	l := len(*h) - 1
	x := (*h)[l]
	*h = (*h)[:l]
	return x
}
