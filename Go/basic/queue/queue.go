package queue

type Queue struct {
	list []interface{}
}

func New() *Queue {
	return &Queue{
		list: make([]interface{}, 0),
	}
}

func (q *Queue) IsEmpty() bool {
	return len(q.list) == 0
}

func (q *Queue) Pop() interface{} {
	if q.IsEmpty() {
		return nil
	}
	tmp := q.list[0]
	q.list = q.list[1:]
	return tmp
}

func (q *Queue) Push(element interface{}) {
	q.list = append(q.list, element)
}

func (q *Queue) ToSlice() []interface{} {
	return q.list
}
