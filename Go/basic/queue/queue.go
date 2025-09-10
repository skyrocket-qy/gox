package queue

type Queue struct {
	list []any
}

func New() *Queue {
	return &Queue{
		list: make([]any, 0),
	}
}

func (q *Queue) IsEmpty() bool {
	return len(q.list) == 0
}

func (q *Queue) Pop() any {
	if q.IsEmpty() {
		return nil
	}

	tmp := q.list[0]
	q.list = q.list[1:]

	return tmp
}

func (q *Queue) Push(element any) {
	q.list = append(q.list, element)
}

func (q *Queue) ToSlice() []any {
	return q.list
}
