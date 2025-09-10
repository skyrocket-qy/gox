package queue

/* @tags: queue,circular list */

// QueueCircularList represents a circular queue implemented using a slice
type QueueCircularList struct {
	items  []interface{}
	front  int
	rear   int
	length int
}

// Enqueue adds an element to the rear of the queue
func (q *QueueCircularList) Enqueue(value interface{}) {
	if q.length == len(q.items) {
		// Resize the slice if it is full
		q.resize()
	}

	q.items[q.rear] = value
	q.rear = (q.rear + 1) % len(q.items)
	q.length++
}

// Dequeue removes and returns an element from the front of the queue
func (q *QueueCircularList) Dequeue() interface{} {
	if q.length == 0 {
		return nil // QueueCircularList is empty
	}

	value := q.items[q.front]
	q.front = (q.front + 1) % len(q.items)
	q.length--

	return value
}

// IsEmpty checks if the queue is empty
func (q *QueueCircularList) IsEmpty() bool {
	return q.length == 0
}

// Size returns the number of elements in the queue
func (q *QueueCircularList) Size() int {
	return q.length
}

// resize doubles the capacity of the slice when it becomes full
func (q *QueueCircularList) resize() {
	capacity := len(q.items)
	newCapacity := capacity * 2

	newItems := make([]interface{}, newCapacity)

	if q.front < q.rear {
		copy(newItems, q.items[q.front:q.rear])
	} else {
		n := copy(newItems, q.items[q.front:])
		copy(newItems[n:], q.items[:q.rear])
	}

	q.items = newItems
	q.front = 0
	q.rear = q.length
}
