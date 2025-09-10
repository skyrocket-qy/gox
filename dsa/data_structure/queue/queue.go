package queue

/* @tags: queue,array,linked list,circular list */

// array
type Queue[T any] struct {
	slice []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		slice: []T{},
	}
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.slice) == 0
}

func (q *Queue[T]) Pop() (T, bool) {
	var x T
	if len(q.slice) > 0 {
		x, q.slice = q.slice[0], q.slice[1:]
		return x, true
	}
	return x, false
}

func (q *Queue[T]) Push(element T) {
	q.slice = append(q.slice, element)
}
