package radixtree

import (
	"testing"
)

func TestRadixTree_InsertAndSearch(t *testing.T) {
	tree := Init([]string{"apple", "app", "application"})

	// Test for existing strings
	if !tree.Search("apple") {
		t.Error("Expected to find 'apple', but it was not found.")
	}
	if !tree.Search("app") {
		t.Error("Expected to find 'app', but it was not found.")
	}
	if !tree.Search("application") {
		t.Error("Expected to find 'application', but it was not found.")
	}

	// Test for non-existing string
	if tree.Search("apples") {
		t.Error("Expected not to find 'apples', but it was found.")
	}

	// Insert a new string and test for it
	tree.Insert("apply")
	if !tree.Search("apply") {
		t.Error("Expected to find 'apply' after insertion, but it was not found.")
	}
}
