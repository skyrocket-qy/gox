package doublylinkedlist

/* @tags: linked_list,doubly_linked_list */

type Node struct {
	Val  int
	prev *Node
	next *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
	len  int
}

func New() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (l *DoublyLinkedList) Len() int {
	return l.len
}

func (l *DoublyLinkedList) AddHead(val int) *Node {
	node := &Node{Val: val}
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		l.head.prev = node
		node.next = l.head
		l.head = node
	}
	l.len++
	return node
}

func (l *DoublyLinkedList) AddTail(val int) *Node {
	node := &Node{Val: val}
	if l.tail == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail.next = node
		node.prev = l.tail
		l.tail = node
	}
	l.len++
	return node
}

func (l *DoublyLinkedList) PopHead() *Node {
	if l.head == nil {
		return nil
	}
	node := l.head
	l.head = l.head.next
	if l.head != nil {
		l.head.prev = nil
	} else {
		l.tail = nil
	}
	node.next = nil
	l.len--
	return node
}

func (l *DoublyLinkedList) PopTail() *Node {
	if l.tail == nil {
		return nil
	}
	node := l.tail
	l.tail = l.tail.prev
	if l.tail != nil {
		l.tail.next = nil
	} else {
		l.head = nil
	}
	node.prev = nil
	l.len--
	return node
}

func (l *DoublyLinkedList) RemoveNode(node *Node) {
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		l.head = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	} else {
		l.tail = node.prev
	}
	l.len--
}
