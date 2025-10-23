package orderedset

import "github.com/skyrocket-qy/gox/dsa/heap"

// OrderedSet represents a collection of unique elements that maintains
// the order in which elements were added.
type OrderedSet interface {
	// Add inserts an element into the set if it is not already present.
	// Returns true if the element was added (i.e., it was new), and false otherwise.
	Add(element interface{}) bool

	// Remove deletes the specified element from the set.
	// Returns true if the element was removed (i.e., it existed), and false otherwise.
	Remove(element interface{}) bool

	// Contains checks if the set holds the specified element.
	Contains(element interface{}) bool

	// Size returns the total number of elements in the set.
	Size() int

	// IsEmpty returns true if the set contains no elements.
	IsEmpty() bool

	// Clear removes all elements from the set.
	Clear()

	// Values returns a slice of all elements in the set, maintaining the
	// order of insertion. This is the primary method to iterate over the set.
	Values() []interface{}

	// Get returns the element at the specified index (based on insertion order).
	// Returns the element and true if the index is valid, or nil and false otherwise.
	Get(index int) (interface{}, bool)

	// IndexOf returns the index (order of insertion) of the element, or -1
	// if the element is not found in the set.
	IndexOf(element interface{}) int
}

type OrderedSetImpl[T comparable] struct {
	h *heap.Heap[T]
	s map[T]struct{}
}

func New[T comparable](eles []T, less func(T, T) bool) *OrderedSetImpl[T] {
	h := heap.New(eles, less)
	s := make(map[T]struct{}, len(eles))
	for _, e := range eles {
		s[e] = struct{}{}
	}

	return &OrderedSetImpl[T]{
		h: h,
		s: s,
	}
}

func (s *OrderedSetImpl[T]) Pop() T {
	if len(s.s) == 0 {
		var zeroValue T
		return zeroValue
	}

	v := s.h.Pop()
	delete(s.s, v)
	return v
}

func (s *OrderedSetImpl[T]) Push(v T) {
	if _, ok := s.s[v]; ok {
		return
	}

	s.h.Push(v)
	s.s[v] = struct{}{}
}
