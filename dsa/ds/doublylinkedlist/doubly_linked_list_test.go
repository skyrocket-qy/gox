package doublylinkedlist

import (
	"testing"
)

func TestDoublyLinkedList(t *testing.T) {
	l := New()

	if l.Len() != 0 {
		t.Errorf("Expected Len to be 0, got %d", l.Len())
	}

	// Add to head
	n1 := l.AddHead(1)
	if l.Len() != 1 || l.head != n1 || l.tail != n1 {
		t.Errorf("AddHead failed")
	}

	// Add to tail
	n2 := l.AddTail(2)
	if l.Len() != 2 || l.tail != n2 || l.head.next != n2 || n2.prev != l.head {
		t.Errorf("AddTail failed")
	}

	// Pop from head
	p1 := l.PopHead()
	if p1 != n1 || l.Len() != 1 || l.head != n2 || l.head.prev != nil {
		t.Errorf("PopHead failed")
	}

	// Pop from tail
	p2 := l.PopTail()
	if p2 != n2 || l.Len() != 0 || l.head != nil || l.tail != nil {
		t.Errorf("PopTail failed")
	}

	// Test removing a node
	n1 = l.AddHead(1)
	n2 = l.AddHead(2)
	n3 := l.AddHead(3)
	l.RemoveNode(n2)

	if l.Len() != 2 || l.head != n3 || l.tail != n1 || n3.next != n1 || n1.prev != n3 {
		t.Errorf("RemoveNode failed")
	}
}

func TestDoublyLinkedList_EdgeCases(t *testing.T) {
	// PopHead on an empty list
	l := New()
	if node := l.PopHead(); node != nil {
		t.Errorf("PopHead on empty list should return nil, but got %v", node)
	}

	// PopTail on an empty list
	if node := l.PopTail(); node != nil {
		t.Errorf("PopTail on empty list should return nil, but got %v", node)
	}

	// PopHead on a list with one element
	l.AddHead(1)

	if node := l.PopHead(); node.Val != 1 || l.Len() != 0 || l.head != nil || l.tail != nil {
		t.Errorf("PopHead on a list with one element failed")
	}

	// PopTail on a list with one element
	l.AddHead(1)

	if node := l.PopTail(); node.Val != 1 || l.Len() != 0 || l.head != nil || l.tail != nil {
		t.Errorf("PopTail on a list with one element failed")
	}

	// RemoveNode for the head node
	l = New()
	n1 := l.AddHead(1)
	n2 := l.AddHead(2)
	l.RemoveNode(n2)

	if l.Len() != 1 || l.head != n1 || l.tail != n1 || n1.prev != nil {
		t.Errorf("RemoveNode for head failed")
	}

	// RemoveNode for the tail node
	l = New()
	n1 = l.AddTail(1)
	n2 = l.AddTail(2)
	l.RemoveNode(n2)

	if l.Len() != 1 || l.head != n1 || l.tail != n1 || n1.next != nil {
		t.Errorf("RemoveNode for tail failed")
	}

	// AddTail on an empty list
	l = New()

	n1 = l.AddTail(1)
	if l.Len() != 1 || l.head != n1 || l.tail != n1 {
		t.Errorf("AddTail on empty list failed")
	}
}
