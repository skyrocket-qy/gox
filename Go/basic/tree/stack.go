package stack

type Stack struct {
	list []interface{}
}

func (stk *Stack) IsEmpty() bool {
	return len(stk.list) == 0
}

func (stk *Stack) Pop() interface{} {
	tmp := stk.list[len(stk.list)-1]
	stk.list = stk.list[:len(stk.list)-1]
	return tmp
}

func (stk *Stack) Push(element interface{}) {
	stk.list = append(stk.list, element)
}
