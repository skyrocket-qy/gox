package tree

type Queue struct {
	list []any
}

func (q *Queue) IsEmpty() bool {
	return len(q.list) == 0
}

func (q *Queue) Pop() any {
	tmp := q.list[0]
	q.list = q.list[1:]

	return tmp
}

func (q *Queue) Push(element any) {
	q.list = append(q.list, element)
}
