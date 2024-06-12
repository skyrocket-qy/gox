package Collection

type Set[T comparable] struct {
	data map[T]bool
}

func NewSet[T comparable]() Set[T] {
	return Set[T]{
		data: make(map[T]bool),
	}
}

func (s *Set[T]) Add(element T) {
	s.data[element] = true
}

func (s *Set[T]) Remove(element T) {
	delete(s.data, element)
}

func (s *Set[T]) Exist(element T) bool {
	if _, ok := s.data[element]; ok {
		return true
	}
	return false
}

func (s *Set[T]) ToSlice() []T {
	eles := []T{}
	for key := range s.data {
		eles = append(eles, key)
	}
	return eles
}
