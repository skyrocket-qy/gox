package queue

type Queue struct {
	list []interface{}
}

func (Q *Queue) IsEmpty() bool {
	return len(Q.list) == 0
}

func (Q *Queue) Pop() interface{} {
	tmp := Q.list[0]
	Q.list = Q.list[1:]
	return tmp
}

func (Q *Queue) Push(element interface{}) {
	Q.list = append(Q.list, element)
}
