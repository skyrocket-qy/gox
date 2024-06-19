package QueueImpl

import "errors"

type node[T any] struct {
	value T
	next  *node[T]
}

type queue[T any] struct {
	head *node[T]
	tail *node[T]
	size int
}

func NewQueue[T any]() *queue[T] {
	return &queue[T]{}
}

func (q *queue[T]) Len() int {
	return q.size
}

func (q *queue[T]) IsEmpty() bool {
	return q.size == 0
}

func (q *queue[T]) Front() (T, error) {
	if q.IsEmpty() {
		var zeroValue T
		return zeroValue, errors.New("queue is empty")
	}
	return q.head.value, nil
}

func (q *queue[T]) Pop() (T, error) {
	if q.IsEmpty() {
		var zeroValue T
		return zeroValue, errors.New("queue is empty")
	}
	oldHead := q.head
	elem := oldHead.value
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
	oldHead.next = nil
	oldHead = nil
	q.size--
	return elem, nil
}

func (q *queue[T]) Push(elem T) {
	newNode := &node[T]{value: elem, next: nil}
	if q.IsEmpty() {
		q.head = newNode
		q.tail = newNode
	} else {
		q.tail.next = newNode
		q.tail = newNode
	}
	q.size++
}

func (q *queue[T]) ToSlice() []T {
	slice := make([]T, 0, q.Len())
	current := q.head
	for current != nil {
		slice = append(slice, current.value)
		current = current.next
	}
	return slice
}
