package prefixtree

// use array instead of struct-map
type Trie_array struct {
	children [26]*Trie_array
}

func Init(strs []string) *Trie_array {
	t := &Trie_array{
		children: [26]*Trie_array{},
	}

	for _, s := range strs {
		t.Insert(s)
	}

	return t
}

func (t *Trie_array) Insert(s string) {
	root := t
	for i := 0; i < len(s); i++ {
		if root.children[s[i]] == nil {
			root.children[s[i]] = &Trie_array{
				children: [26]*Trie_array{},
			}
		}
		root = root.children[s[i]]
	}
}

func (t *Trie_array) Remove(s string) {
	root := t
	for i := 0; i < len(s)-1; i++ {
		if root.children[s[i]] == nil {
			return
		}
		root = root.children[s[i]]
	}
	root.children[s[len(s)-1]] = nil
}

func (t *Trie_array) Search(s string) bool {
	root := t
	for i := 0; i < len(s); i++ {
		if root.children[s[i]] == nil {
			return false
		}
		root = root.children[s[i]]
	}
	return true
}
