package trie

type Trie struct {
	count    int
	children map[rune]*Trie
}

func Constructor() Trie {
	return Trie{children: make(map[rune]*Trie)}
}

func (t *Trie) Insert(word string) {
	for _, char := range word {
		if node, ok := t.children[char]; ok {
			t = node
		} else {
			t.children[char] = &Trie{children: make(map[rune]*Trie)}
			t = t.children[char]
		}
	}
	t.count++
}

func (t *Trie) Search(word string) bool {
	for _, char := range word {
		if node, ok := t.children[char]; ok {
			t = node
		} else {
			return false
		}
	}
	return t.count > 0
}

func (t *Trie) StartsWith(prefix string) bool {
	for _, char := range prefix {
		if node, ok := t.children[char]; ok {
			t = node
		} else {
			return false
		}
	}
	return true
}

func (t *Trie) Remove(word string) {
	for _, char := range word[:len(word)-1] {
		if node, ok := t.children[char]; ok {
			t = node
		} else {
			return
		}
	}
	//delete node or decrease count
	if node, ok := t.children[rune(word[len(word)-1])]; ok {
		if node.count == 1 {
			if len(node.children) == 0 {
				delete(t.children, rune(word[len(word)-1]))
				return
			}
			node.count >>= 1
		} else if node.count > 1 {
			node.count--
		}
	}
}
