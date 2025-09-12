package prefixtree

import "testing"

func TestTrieArray(t *testing.T) {
	trie := Init([]string{"apple", "apply"})
	if !trie.Search("apple") {
		t.Error("expected to find apple")
	}
	if !trie.Search("apply") {
		t.Error("expected to find apply")
	}
	if trie.Search("app") {
		// This search should be true as it is a prefix search
	} else {
		t.Error("expected to find prefix app")
	}
	if trie.Search("apples") {
		t.Error("expected to not find apples")
	}
	trie.Remove("apple")
	if trie.Search("apple") {
		t.Error("expected to not find apple")
	}
	if !trie.Search("apply") {
		t.Error("expected to find apply after removing apple")
	}
}

func TestTrieArrayRemoveNotFound(t *testing.T) {
	trie := Init([]string{"apple"})
	trie.Remove("apply")
	if !trie.Search("apple") {
		t.Error("expected to find apple")
	}
}

func TestTrieArrayInsert(t *testing.T) {
	trie := &Trie_array{}
	trie.Insert("word")
	if !trie.Search("word") {
		t.Error("expected to find word")
	}
}
