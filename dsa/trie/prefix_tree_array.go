package prefixtree

// use array instead of struct-map.
// This implementation only supports lowercase English letters 'a'-'z'.
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
	for i := range len(s) {
		if root.children[s[i]-'a'] == nil {
			root.children[s[i]-'a'] = &Trie_array{
				children: [26]*Trie_array{},
			}
		}

		root = root.children[s[i]-'a']
	}
}

func (t *Trie_array) Remove(s string) {
	root := t
	for i := range len(s) - 1 {
		if root.children[s[i]-'a'] == nil {
			return
		}

		root = root.children[s[i]-'a']
	}

	root.children[s[len(s)-1]-'a'] = nil
}

func (t *Trie_array) Search(s string) bool {
	root := t
	for i := range len(s) {
		if root.children[s[i]-'a'] == nil {
			return false
		}

		root = root.children[s[i]-'a']
	}

	return true
}
