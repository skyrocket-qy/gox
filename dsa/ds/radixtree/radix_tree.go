package radixtree

import "strings"

/* @tags: tree */

// leafNode is used to represent a value.
type leafNode struct {
	key string
	val any
}

type radixTree struct {
	// leaf is used to store possible leaf
	leaf *leafNode

	// prefix is the common prefix we ignore
	prefix string

	// children should be stored in-order for iteration.
	children map[byte]*radixTree
}

func New(strs []string) *radixTree {
	t := &radixTree{
		children: make(map[byte]*radixTree),
	}

	for _, s := range strs {
		t.Insert(s)
	}

	return t
}

// longestPrefix finds the length of the shared prefix
// of two strings.
func longestPrefix(k1, k2 string) int {
	maxLength := len(k1)
	if l := len(k2); l < maxLength {
		maxLength = l
	}

	var i int
	for i = 0; i < maxLength; i++ {
		if k1[i] != k2[i] {
			break
		}
	}

	return i
}

func (t *radixTree) Insert(s string) {
	var parent *radixTree

	n := t

	search := s
	for {
		// Handle key exhaustion
		if len(search) == 0 {
			if n.leaf != nil {
				return
			}

			n.leaf = &leafNode{
				key: s,
				val: s,
			}

			return
		}

		// Look for the edge
		parent = n

		child, ok := n.children[search[0]]
		if !ok {
			// No edge, create one
			e := &radixTree{
				leaf: &leafNode{
					key: s,
					val: s,
				},
				prefix:   search,
				children: make(map[byte]*radixTree),
			}
			parent.children[search[0]] = e

			return
		}

		n = child

		// Determine longest prefix of the search key on match
		commonPrefix := longestPrefix(search, n.prefix)
		if commonPrefix == len(n.prefix) {
			search = search[commonPrefix:]

			continue
		}

		// Split the node
		child = &radixTree{
			prefix:   search[:commonPrefix],
			children: make(map[byte]*radixTree),
		}
		parent.children[search[0]] = child

		// Restore the existing node
		child.children[n.prefix[commonPrefix]] = n
		n.prefix = n.prefix[commonPrefix:]

		// Create a new leaf node
		leaf := &leafNode{
			key: s,
			val: s,
		}

		// If the new key is a subset, add to this node
		search = search[commonPrefix:]
		if len(search) == 0 {
			child.leaf = leaf

			return
		}

		// Create a new edge for the node
		child.children[search[0]] = &radixTree{
			leaf:     leaf,
			prefix:   search,
			children: make(map[byte]*radixTree),
		}

		return
	}
}

// Delete is used to delete a key, returning the previous
// value and if it was deleted.
func (t *radixTree) Remove(s string) {
	var (
		parent *radixTree
		label  byte
	)

	n := t

	search := s
	for {
		// Check for key exhaution
		if len(search) == 0 {
			if n.leaf == nil {
				break
			}

			goto DELETE
		}

		// Look for an edge
		parent = n
		label = search[0]

		child, ok := n.children[label]
		if !ok {
			break
		}

		n = child

		// Consume the search prefix
		if strings.HasPrefix(search, n.prefix) {
			search = search[len(n.prefix):]
		} else {
			break
		}
	}

	return

DELETE:
	// Delete the leaf
	n.leaf = nil

	// Check if we should delete this node from the parent
	if parent != nil && len(n.children) == 0 {
		delete(parent.children, label)
	}

	// Check if we should merge this node
	if n != t && len(n.children) == 1 && n.leaf == nil {
		for _, child := range n.children {
			n.prefix += child.prefix
			n.leaf = child.leaf
			n.children = child.children
		}
	}

	// Check if we should merge the parent's other child
	if parent != nil && parent != t && len(parent.children) == 1 && parent.leaf == nil {
		for _, child := range parent.children {
			parent.prefix += child.prefix
			parent.leaf = child.leaf
			parent.children = child.children
		}
	}
}

func (t *radixTree) Search(s string) bool {
	n := t

	search := s
	for {
		// Check for key exhaustion
		if len(search) == 0 {
			if n.leaf != nil {
				return true
			}

			break
		}

		// Look for an edge
		child, ok := n.children[search[0]]
		if !ok {
			break
		}

		n = child

		// Consume the search prefix
		if strings.HasPrefix(search, n.prefix) {
			search = search[len(n.prefix):]
		} else {
			break
		}
	}

	return false
}
