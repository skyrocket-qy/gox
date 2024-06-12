package queue

import "errors"

type node[T any] struct {
	value T
	next  *node[T]
}

type Queue[T any] interface {
	IsEmpty() bool
	Pop() (T, error)
	Push(element T)
	ToSlice() []T
	Len() int
}

type queueImpl[T any] struct {
	front *node[T]
	rear  *node[T]
	size  int
}

func NewQueue[T any]() Queue[T] {
	return &queueImpl[T]{}
}

func (q *queueImpl[T]) Len() int {
	return q.size
}

func (q *queueImpl[T]) IsEmpty() bool {
	return q.size == 0
}

func (q *queueImpl[T]) Pop() (T, error) {
	if q.IsEmpty() {
		var zeroValue T
		return zeroValue, errors.New("queue is empty")
	}
	element := q.front.value
	q.front = q.front.next
	q.size--
	return element, nil
}

func (q *queueImpl[T]) Push(element T) {
	newNode := &node[T]{value: element, next: nil}
	if q.IsEmpty() {
		q.front = newNode
		q.rear = newNode
	} else {
		q.rear.next = newNode
		q.rear = newNode
	}
	q.size++
}

func (q *queueImpl[T]) ToSlice() []T {
	slice := make([]T, 0, q.Len())
	current := q.front
	for current != nil {
		slice = append(slice, current.value)
		current = current.next
	}
	return slice
}
