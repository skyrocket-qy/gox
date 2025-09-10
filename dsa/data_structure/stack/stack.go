package stack

/* @tags: stack */

type Stack[T any] struct {
	slice []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		slice: []T{},
	}
}

func (stk *Stack[T]) Push(element T) {
	stk.slice = append(stk.slice, element)
}

func (stk *Stack[T]) Pop() (T, bool) {
	var x T
	if len(stk.slice) > 0 {
		x, stk.slice = stk.slice[len(stk.slice)-1], stk.slice[:len(stk.slice)-1]
		return x, true
	}
	return x, false
}

func (stk *Stack[T]) IsEmpty() bool {
	return len(stk.slice) == 0
}
