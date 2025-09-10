package radixtree

/* @tags: tree */

type radixTree struct {
	children map[string]*radixTree
}

func Init(strs []string) *radixTree {
	t := &radixTree{
		children: make(map[string]*radixTree),
	}

	for _, s := range strs {
		t.Insert(s)
	}

	return t
}

func (t *radixTree) Insert(s string) {
	root := t
	preI := 0
	for i := 1; i <= len(s); i++ {
		curS := s[preI:i]
		if child, ok := t.children[curS]; ok {
			root = child
			preI = i
		}
	}

	if preI == len(s) {
		return
	}
	root.children[s[preI:]] = &radixTree{
		children: make(map[string]*radixTree),
	}
}

func (t *radixTree) Remove(s string) {
	root := t
	preI := 0
	for i := 1; i <= len(s); i++ {
		curS := s[preI:i]
		if child, ok := root.children[curS]; ok {
			root = child
			preI = i
		}
	}
}
func (t *radixTree) Search(s string) bool {

	return true
}
