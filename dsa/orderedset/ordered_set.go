package orderedset

import "github.com/skyrocket-qy/gox/dsa/heap"

type OrderedSet[T comparable] struct {
	h *heap.Heap[T]
	s map[T]struct{}
}

func New[T comparable](eles []T, less func(T, T) bool) *OrderedSet[T] {
	h := heap.New(eles, less)
	s := make(map[T]struct{}, len(eles))
	for _, e := range eles {
		s[e] = struct{}{}
	}

	return &OrderedSet[T]{
		h: h,
		s: s,
	}
}

func (s *OrderedSet[T]) Pop() T {
	if len(s.s) == 0 {
		var zeroValue T
		return zeroValue
	}

	v := s.h.Pop()
	delete(s.s, v)
	return v
}

func (s *OrderedSet[T]) Push(v T) {
	if _, ok := s.s[v]; ok {
		return
	}

	s.h.Push(v)
	s.s[v] = struct{}{}
}
