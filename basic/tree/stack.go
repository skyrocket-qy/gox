package tree

type Stack struct {
	list []any
}

func (stk *Stack) IsEmpty() bool {
	return len(stk.list) == 0
}

func (stk *Stack) Pop() any {
	tmp := stk.list[len(stk.list)-1]
	stk.list = stk.list[:len(stk.list)-1]

	return tmp
}

func (stk *Stack) Push(element any) {
	stk.list = append(stk.list, element)
}
