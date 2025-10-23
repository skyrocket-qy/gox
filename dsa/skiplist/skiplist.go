package skiplist

import (
	"math/rand"
	"time"
)

const (
	// MaxLevel is the maximum level of the skip list.
	MaxLevel = 32
	// Probability is the probability of increasing the level of a new node.
	Probability = 0.25
)

// Node represents a node in the skip list.
type Node struct {
	Value interface{}
	// Forward[i] is the node at level i.
	Forward []*Node
}

// SkipList represents a skip list.
type SkipList struct {
	Header    *Node
	Level     int
	Length    int
	rand      *rand.Rand
	Comparator Comparator
}

// Comparator is a function that compares two interface{} values.
// It returns -1 if a < b, 0 if a == b, and 1 if a > b.
type Comparator func(a, b interface{}) int

func (sl *SkipList) GetRangeByValue(min, max interface{}) []interface{} {
	var result []interface{}
	current := sl.Header

	// Find the first element >= min
	for i := sl.Level - 1; i >= 0; i-- {
		for current.Forward[i] != nil && sl.Comparator(current.Forward[i].Value, min) < 0 {
			current = current.Forward[i]
		}
	}
	current = current.Forward[0]

	// Collect elements <= max
	for current != nil && sl.Comparator(current.Value, max) <= 0 {
		result = append(result, current.Value)
		current = current.Forward[0]
	}

	return result
}

func (sl *SkipList) GetByIndex(index int) interface{} {
	if index < 0 || index >= sl.Length {
		return nil
	}

	current := sl.Header
	count := -1 // Header is at -1 index

	// Traverse at level 0 to find the element at the given index.
	for current.Forward[0] != nil && count < index {
		current = current.Forward[0]
		count++
	}

	if count == index {
		return current.Value
	}

	return nil
}

func (sl *SkipList) Len() int {
	return sl.Length
}

func (sl *SkipList) Delete(value interface{}) {
	update := make([]*Node, MaxLevel)
	current := sl.Header

	// Find the node to be deleted and store the update path.
	for i := sl.Level - 1; i >= 0; i-- {
		for current.Forward[i] != nil && sl.Comparator(current.Forward[i].Value, value) < 0 {
			current = current.Forward[i]
		}
		update[i] = current
	}

	// Move to the next node at level 0.
	current = current.Forward[0]

	// If the node is found, delete it.
	if current != nil && sl.Comparator(current.Value, value) == 0 {
		for i := 0; i < sl.Level; i++ {
			if update[i].Forward[i] != current {
				break
			}
			update[i].Forward[i] = current.Forward[i]
		}

		// Decrease the skip list's level if the highest level becomes empty.
		for sl.Level > 1 && sl.Header.Forward[sl.Level-1] == nil {
			sl.Level--
		}

		sl.Length--
	}
}

func (sl *SkipList) Search(value interface{}) interface{} {
	current := sl.Header

	// Traverse from the highest level down to the lowest.
	for i := sl.Level - 1; i >= 0; i-- {
		for current.Forward[i] != nil && sl.Comparator(current.Forward[i].Value, value) < 0 {
			current = current.Forward[i]
		}
	}

	// Move to the next node at level 0.
	current = current.Forward[0]

	// Check if the current node's value matches the target value.
	if current != nil && sl.Comparator(current.Value, value) == 0 {
		return current.Value
	}

	return nil
}

func (sl *SkipList) Insert(value interface{}) {
	update := make([]*Node, MaxLevel)
	current := sl.Header

	// Find the insertion point at each level.
	for i := sl.Level - 1; i >= 0; i-- {
		for current.Forward[i] != nil && sl.Comparator(current.Forward[i].Value, value) < 0 {
			current = current.Forward[i]
		}
		update[i] = current
	}

	// Move to the next node at level 0.
	current = current.Forward[0]

	// If the value already exists, update it.
	if current != nil && sl.Comparator(current.Value, value) == 0 {
		current.Value = value
		return
	}

	// Generate a random level for the new node.
	newLevel := sl.randomLevel()

	// If the new level is greater than the current skip list's level, update the skip list's level.
	if newLevel > sl.Level {
		for i := sl.Level; i < newLevel; i++ {
			update[i] = sl.Header
		}
		sl.Level = newLevel
	}

	// Create a new node.
	newNode := &Node{
		Value:   value,
		Forward: make([]*Node, newLevel),
	}

	// Insert the new node.
	for i := 0; i < newLevel; i++ {
		newNode.Forward[i] = update[i].Forward[i]
		update[i].Forward[i] = newNode
	}

	sl.Length++
}

func (sl *SkipList) randomLevel() int {
	level := 1
	for sl.rand.Float64() < Probability && level < MaxLevel {
		level++
	}
	return level
}

// New creates a new skip list.
func New(comp Comparator) *SkipList {
	header := &Node{
		Forward: make([]*Node, MaxLevel),
	}
	return &SkipList{
		Header:    header,
		Level:     0,
		Length:    0,
		rand:      rand.New(rand.NewSource(time.Now().UnixNano())),
		Comparator: comp,
	}
}
