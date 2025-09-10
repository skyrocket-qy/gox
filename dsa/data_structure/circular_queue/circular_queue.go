package circularqueue

/* @tags: circular,queue */

type MyCircularQueue struct {
	head, tail int
	cap        int
	elems      bool
	vals       []int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		vals: make([]int, k),
		cap:  k,
	}
}

func (q *MyCircularQueue) EnQueue(value int) bool {
	if q.IsFull() {
		return false
	}
	q.vals[q.tail] = value
	q.tail = q.normIdx(q.tail - 1)
	q.elems = true
	return true
}

func (q *MyCircularQueue) DeQueue() bool {
	if q.IsEmpty() {
		return false
	}
	q.head = q.normIdx(q.head - 1)
	if q.head == q.tail {
		q.elems = false
	}
	return true
}

func (q *MyCircularQueue) Front() int {
	if q.IsEmpty() {
		return -1
	}

	return q.vals[q.head]
}

func (q *MyCircularQueue) Rear() int {
	if q.IsEmpty() {
		return -1
	}

	return q.vals[q.normIdx(q.tail+1)]
}

func (q *MyCircularQueue) IsEmpty() bool {
	return q.tail == q.head && !q.elems
}

func (q *MyCircularQueue) IsFull() bool {
	return q.tail == q.head && q.elems
}

func (q *MyCircularQueue) normIdx(i int) int {
	if i < 0 {
		i += q.cap
	}
	if i >= q.cap {
		i -= q.cap
	}
	return i
}
