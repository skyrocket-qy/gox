package stack

type Stack struct {
	list []any
}

func New() *Stack {
	return &Stack{
		list: make([]any, 0),
	}
}

func (s *Stack) IsEmpty() bool {
	return len(s.list) == 0
}

func (s *Stack) Pop() any {
	if s.IsEmpty() {
		return nil
	}

	tmp := s.list[len(s.list)-1]
	s.list = s.list[:len(s.list)-1]

	return tmp
}

func (s *Stack) Push(element any) {
	s.list = append(s.list, element)
}

func (s *Stack) ToSlice() []any {
	return s.list
}
