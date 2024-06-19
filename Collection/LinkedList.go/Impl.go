package LinkedList

import (
	"errors"
	"fmt"
)

// Node represents a node in a singly linked list.
type Node[T any] struct {
	value T
	next  *Node[T]
}

// SinglyLinkedList is a singly linked list that implements the List interface.
type SinglyLinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

// Append adds a value to the end of the list.
func (list *SinglyLinkedList[T]) Append(value T) {
	newNode := &Node[T]{value: value}
	if list.tail == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		list.tail.next = newNode
		list.tail = newNode
	}
	list.size++
}

// Prepend adds a value to the beginning of the list.
func (list *SinglyLinkedList[T]) Prepend(value T) {
	newNode := &Node[T]{value: value, next: list.head}
	if list.head == nil {
		list.tail = newNode
	}
	list.head = newNode
	list.size++
}

// Insert inserts a value at a specific index.
func (list *SinglyLinkedList[T]) Insert(index int, value T) error {
	if index < 0 || index > list.size {
		return errors.New("index out of bounds")
	}

	if index == 0 {
		list.Prepend(value)
		return nil
	}

	newNode := &Node[T]{value: value}
	current := list.head
	for i := 0; i < index-1; i++ {
		current = current.next
	}
	newNode.next = current.next
	current.next = newNode

	if newNode.next == nil {
		list.tail = newNode
	}
	list.size++
	return nil
}

// Remove removes a value from the list.
func (list *SinglyLinkedList[T]) Remove(value T) error {
	if list.head == nil {
		return errors.New("list is empty")
	}

	if list.head.value == value {
		list.head = list.head.next
		if list.head == nil {
			list.tail = nil
		}
		list.size--
		return nil
	}

	prev, current := list.head, list.head.next
	for current != nil && current.value != value {
		prev = current
		current = current.next
	}

	if current == nil {
		return errors.New("value not found in list")
	}

	prev.next = current.next
	if current.next == nil {
		list.tail = prev
	}
	list.size--
	return nil
}

// Pop removes and returns the first element.
func (list *SinglyLinkedList[T]) Pop() (T, error) {
	if list.head == nil {
		var zeroValue T
		return zeroValue, errors.New("list is empty")
	}

	value := list.head.value
	list.head = list.head.next
	if list.head == nil {
		list.tail = nil
	}
	list.size--
	return value, nil
}

// Get returns the value at a specific index.
func (list *SinglyLinkedList[T]) Get(index int) (T, error) {
	if index < 0 || index >= list.size {
		var zeroValue T
		return zeroValue, errors.New("index out of bounds")
	}

	current := list.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	return current.value, nil
}

// Size returns the current size of the list.
func (list *SinglyLinkedList[T]) Size() int {
	return list.size
}

// PrintList prints all values in the list.
func (list *SinglyLinkedList[T]) PrintList() {
	current := list.head
	for current != nil {
		fmt.Printf("%v -> ", current.value)
		current = current.next
	}
	fmt.Println("nil")
}
