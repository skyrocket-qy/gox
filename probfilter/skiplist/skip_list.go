package skiplist

import (
	"math/rand"
	"time"
)

const (
	maxLevel    = 16  // Maximum level for a skip list
	probability = 0.5 // Probability of increasing level
)

// Node represents a node in the skip list.
type Node struct {
	value  int
	levels []*Node // Array of pointers to next nodes at different levels
}

// SkipList represents the skip list data structure.
type SkipList struct {
	header *Node // Pointer to the header node
	level  int   // Current maximum level of the skip list
	rand   *rand.Rand
}

// NewSkipList creates and initializes a new SkipList.
func NewSkipList() *SkipList {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source) //nolint:gosec
	header := &Node{value: 0, levels: make([]*Node, maxLevel)}

	return &SkipList{
		header: header,
		level:  0,
		rand:   r,
	}
}

// randomLevel generates a random level for a new node.
func (sl *SkipList) randomLevel() int {
	lvl := 0
	for sl.rand.Float64() < probability && lvl < maxLevel-1 {
		lvl++
	}

	return lvl
}

// Insert inserts a value into the skip list.
func (sl *SkipList) Insert(value int) {
	update := make([]*Node, maxLevel)
	current := sl.header

	for i := sl.level; i >= 0; i-- {
		for current.levels[i] != nil && current.levels[i].value < value {
			current = current.levels[i]
		}

		update[i] = current
	}

	current = current.levels[0]

	if current == nil || current.value != value {
		// Value not present, insert new node
		lvl := sl.randomLevel()

		if lvl > sl.level {
			for i := sl.level + 1; i <= lvl; i++ {
				update[i] = sl.header
			}

			sl.level = lvl
		}

		newNode := &Node{value: value, levels: make([]*Node, lvl+1)}

		for i := 0; i <= lvl; i++ {
			newNode.levels[i] = update[i].levels[i]
			update[i].levels[i] = newNode
		}
	}
}

// Search searches for a value in the skip list.
func (sl *SkipList) Search(value int) bool {
	current := sl.header

	for i := sl.level; i >= 0; i-- {
		for current.levels[i] != nil && current.levels[i].value < value {
			current = current.levels[i]
		}
	}

	current = current.levels[0]

	return current != nil && current.value == value
}

// Delete deletes a value from the skip list.
func (sl *SkipList) Delete(value int) {
	update := make([]*Node, maxLevel)
	current := sl.header

	for i := sl.level; i >= 0; i-- {
		for current.levels[i] != nil && current.levels[i].value < value {
			current = current.levels[i]
		}

		update[i] = current
	}

	current = current.levels[0]

	if current != nil && current.value == value {
		for i := 0; i <= sl.level; i++ {
			if update[i].levels[i] != current {
				break
			}

			update[i].levels[i] = current.levels[i]
		}

		for sl.level > 0 && sl.header.levels[sl.level] == nil {
			sl.level--
		}
	}
}
