package queue

/* @tags: queue,linked list */

// Node represents a node in the linked list
type Node struct {
	value interface{}
	next  *Node
}

// Queue represents a queue implemented using a linked list
type QueueLinkedList struct {
	front *Node
	rear  *Node
}

// Enqueue adds an element to the rear of the queue
func (q *QueueLinkedList) Enqueue(value interface{}) {
	newNode := &Node{value: value, next: nil}

	if q.front == nil {
		q.front = newNode
		q.rear = newNode
	} else {
		q.rear.next = newNode
		q.rear = newNode
	}
}

// Dequeue removes and returns an element from the front of the queue
func (q *QueueLinkedList) Dequeue() interface{} {
	if q.front == nil {
		return nil // Queue is empty
	}

	value := q.front.value
	q.front = q.front.next

	if q.front == nil {
		q.rear = nil // Last element removed, update rear
	}

	return value
}

// IsEmpty checks if the queue is empty
func (q *QueueLinkedList) IsEmpty() bool {
	return q.front == nil
}

// Size returns the number of elements in the queue
func (q *QueueLinkedList) Size() int {
	count := 0
	current := q.front

	for current != nil {
		count++
		current = current.next
	}

	return count
}
