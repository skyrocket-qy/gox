package queue

import "errors"

type Queue[T any] interface {
	IsEmpty() bool
	Pop() (T, error)
	Push(element T)
	ToSlice() []T
	Len() int
}

type queueImpl[T any] struct {
	list []T
}

func NewQueue[T any]() Queue[T] {
	return &queueImpl[T]{
		list: make([]T, 0),
	}
}

func (q *queueImpl[T]) Len() int {
	return len(q.list)
}

func (q *queueImpl[T]) IsEmpty() bool {
	return len(q.list) == 0
}

func (q *queueImpl[T]) Pop() (T, error) {
	if q.IsEmpty() {
		var zeroValue T
		return zeroValue, errors.New("queue is empty")
	}
	tmp := q.list[0]
	q.list = q.list[1:]
	return tmp, nil
}

func (q *queueImpl[T]) Push(element T) {
	q.list = append(q.list, element)
}

func (q *queueImpl[T]) ToSlice() []T {
	return q.list
}
