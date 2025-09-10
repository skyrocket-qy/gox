package tree

type Queue struct {
	list []any
}

func (Q *Queue) IsEmpty() bool {
	return len(Q.list) == 0
}

func (Q *Queue) Pop() any {
	tmp := Q.list[0]
	Q.list = Q.list[1:]

	return tmp
}

func (Q *Queue) Push(element any) {
	Q.list = append(Q.list, element)
}
