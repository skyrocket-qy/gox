package custom

type Node struct {
	Priority uint
}

type PriorityQueue []Node

func NewPriorityQueue() PriorityQueue { return append(PriorityQueue{}, Node{}) }

func (q *PriorityQueue) Push(n Node) {
	*q = append(*q, n)
	q.up(len(*q) - 1)
}

func (q *PriorityQueue) Pop() Node {
	n := len(*q) - 1
	res := (*q)[1]
	(*q)[1] = (*q)[n]
	*q = (*q)[:n]
	q.down(1)
	return res
}

func (q PriorityQueue) up(i int) {
	for j := i >> 1; i > 1 && q[j].Priority > q[i].Priority; i, j = j, j>>1 {
		q[i], q[j] = q[j], q[i]
	}
}

func (q PriorityQueue) down(i int) {
	for j := i << 1; j < len(q); i, j = j, j<<1 {
		if j+1 < len(q) && q[j].Priority > q[j+1].Priority {
			j++
		}
		if q[i].Priority < q[j].Priority {
			break
		}
		q[i], q[j] = q[j], q[i]
	}
}
