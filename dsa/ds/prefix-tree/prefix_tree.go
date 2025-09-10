package prefixtree

/* @tags: prefix tree,tree */

type Trie struct {
	children map[byte]*Trie
	isWord   bool
}

func NewTrie() *Trie {
	return &Trie{
		children: make(map[byte]*Trie),
	}
}

func (t *Trie) Insert(s string) {
	var c byte
	for i := 0; i < len(s); i++ {
		c = s[i]
		if _, ok := t.children[c]; !ok {
			t.children[c] = NewTrie()
		}
		t = t.children[c]
	}
	t.isWord = true
}

func (t *Trie) Remove(s string) {
	if len(s) == 0 {
		return
	}

	// if do not find, return
	// if find in middle layer, change isWord = false
	// if find in leaf, delete node and back to delete leaf

	path := []*Trie{t}
	var c byte
	for i := 0; i < len(s); i++ {
		c = s[i]
		if child, ok := t.children[c]; ok {
			path = append(path, child)
			t = child
		} else {
			return
		}
	}

	if len(t.children) != 0 {
		t.isWord = false
		return
	}

	i, j := len(s)-1, len(path)-2
	for {
		delete(path[j].children, s[i])
		if len(path[j].children) == 0 && !path[j].isWord {
			i--
			j--
		} else {
			break
		}
	}
}

func (t *Trie) Search(s string) bool {
	for i := 0; i < len(s); i++ {
		if child, ok := t.children[s[i]]; ok {
			t = child
		} else {
			return false
		}
	}

	return t.isWord
}
