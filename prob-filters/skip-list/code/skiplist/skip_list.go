package skiplist

import (
	"math/rand"
	"time"
)

const (
	MaxLevel = 16 // Maximum level for a skip list node
	P        = 0.5 // Probability factor for determining node level
)

// Node represents a node in the Skip List.
type Node struct {
	Value int
	// Forward pointers for each level
	Forward []*Node
}

// SkipList represents the Skip List data structure.
type SkipList struct {
	Header *Node
	Level  int // Current maximum level of the skip list
	rand   *rand.Rand
}

// New creates a new SkipList.
func New() *SkipList {
	header := &Node{
		Forward: make([]*Node, MaxLevel),
	}
	return &SkipList{
		Header: header,
		Level:  0,
		rand:   rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// randomLevel generates a random level for a new node.
func (sl *SkipList) randomLevel() int {
	level := 0
	for sl.rand.Float64() < P && level < MaxLevel-1 {
		level++
	}
	return level
}

// Insert inserts a value into the Skip List.
func (sl *SkipList) Insert(value int) {
	update := make([]*Node, MaxLevel)
	current := sl.Header

	// Find insertion point at each level
	for i := sl.Level; i >= 0; i-- {
		for current.Forward[i] != nil && current.Forward[i].Value < value {
			current = current.Forward[i]
		}
		update[i] = current
	}

	// Move to level 0 and check if value already exists
	current = current.Forward[0]
	if current != nil && current.Value == value {
		return // Value already exists
	}

	// Generate a random level for the new node
	newLevel := sl.randomLevel()

	// If new level is greater than current SkipList level, update header
	if newLevel > sl.Level {
		for i := sl.Level + 1; i <= newLevel; i++ {
			update[i] = sl.Header
		}
		sl.Level = newLevel
	}

	// Create new node
	newNode := &Node{
		Value:   value,
		Forward: make([]*Node, newLevel+1),
	}

	// Insert node at appropriate positions
	for i := 0; i <= newLevel; i++ {
		newNode.Forward[i] = update[i].Forward[i]
		update[i].Forward[i] = newNode
	}
}

// Search searches for a value in the Skip List.
func (sl *SkipList) Search(value int) bool {
	current := sl.Header
	for i := sl.Level; i >= 0; i-- {
		for current.Forward[i] != nil && current.Forward[i].Value < value {
			current = current.Forward[i]
		}
	}
	current = current.Forward[0]
	return current != nil && current.Value == value
}

// Delete deletes a value from the Skip List.
func (sl *SkipList) Delete(value int) bool {
	update := make([]*Node, MaxLevel)
	current := sl.Header

	// Find deletion point at each level
	for i := sl.Level; i >= 0; i-- {
		for current.Forward[i] != nil && current.Forward[i].Value < value {
			current = current.Forward[i]
		}
		update[i] = current
	}

	// Move to level 0 and check if value exists
	current = current.Forward[0]
	if current == nil || current.Value != value {
		return false // Value not found
	}

	// Delete node at each level
	for i := 0; i <= sl.Level; i++ {
		if update[i].Forward[i] != current {
			break // Node not found at this level
		}
		update[i].Forward[i] = current.Forward[i]
	}

	// Update SkipList level if necessary
	for sl.Level > 0 && sl.Header.Forward[sl.Level] == nil {
		sl.Level--
	}
	return true
}