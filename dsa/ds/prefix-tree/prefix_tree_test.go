package prefixtree

import "testing"

func TestTrie(t *testing.T) {
	trie := NewTrie()
	trie.Insert("apple")
	if !trie.Search("apple") {
		t.Error("expected to find apple")
	}
	if trie.Search("app") {
		t.Error("expected to not find app")
	}
	trie.Insert("app")
	if !trie.Search("app") {
		t.Error("expected to find app")
	}
	trie.Remove("apple")
	if trie.Search("apple") {
		t.Error("expected to not find apple")
	}
	if !trie.Search("app") {
		t.Error("expected to still find app")
	}
}

func TestTrieRemove(t *testing.T) {
	trie := NewTrie()
	trie.Insert("a")
	trie.Remove("a")
	if trie.Search("a") {
		t.Error("expected to not find a")
	}

	trie.Insert("apple")
	trie.Remove("apple")
	if trie.Search("apple") {
		t.Error("expected to not find apple")
	}

	trie.Insert("apple")
	trie.Insert("apply")
	trie.Remove("apple")
	if trie.Search("apple") {
		t.Error("expected to not find apple")
	}
	if !trie.Search("apply") {
		t.Error("expected to still find apply")
	}

	trie.Insert("apple")
	trie.Insert("app")
	trie.Remove("app")
	if trie.Search("app") {
		t.Error("expected to not find app")
	}
	if !trie.Search("apple") {
		t.Error("expected to still find apple")
	}
}

func TestTrieSearchEmpty(t *testing.T) {
	trie := NewTrie()
	if trie.Search("") {
		t.Error("expected to not find empty string")
	}
	trie.Insert("")
	if !trie.Search("") {
		t.Error("expected to find empty string")
	}
}

func TestTrieRemoveEmpty(t *testing.T) {
	trie := NewTrie()
	trie.Remove("")
}

func TestTrieRemoveNotFound(t *testing.T) {
	trie := NewTrie()
	trie.Insert("apple")
	trie.Remove("apply")
	if !trie.Search("apple") {
		t.Error("expected to find apple")
	}
}
