package doublylinkedlist

/* @tags: linked_list,doubly_linked_list */

type doubly_linked_list struct {
}

func (l *doubly_linked_list) Len() int

func (l *doubly_linked_list) PopHead()

func (l *doubly_linked_list) PopTail()

func (l *doubly_linked_list) AddHead()

func (l *doubly_linked_list) AddTail()

func (l *doubly_linked_list) RemoveNode()
