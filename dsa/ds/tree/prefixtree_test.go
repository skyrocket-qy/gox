package tree

import "testing"

func TestTreeTrie(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")

	if !trie.Search("apple") {
		t.Error("expected to find apple")
	}

	if trie.Search("app") {
		t.Error("expected to not find app")
	}

	if !trie.StartsWith("app") {
		t.Error("expected to find prefix app")
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

	trie.Remove("app")

	if trie.Search("app") {
		t.Error("expected to not find app")
	}
}

func TestTreeTrieRemove(t *testing.T) {
	trie := Constructor()
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
}

func TestTreeTrieRemoveNonExistent(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")
	trie.Remove("apply")

	if !trie.Search("apple") {
		t.Error("expected to find apple")
	}
}
