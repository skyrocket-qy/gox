package radixtree_test

import (
	"testing"

	"github.com/skyrocket-qy/gox/dsa/radixtree"
)

func TestRadixTree_InsertAndSearch(t *testing.T) {
	tree := radixtree.New([]string{"apple", "app", "application"})

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

	// Test for deleting a string
	tree.Remove("app")

	if tree.Search("app") {
		t.Error("Expected not to find 'app' after deletion, but it was found.")
	}

	if !tree.Search("apple") {
		t.Error("Expected to find 'apple' after deleting 'app', but it was not found.")
	}
}

func TestRadixTree_Remove(t *testing.T) {
	tree := radixtree.New([]string{"apple", "app", "application", "apply"})

	tree.Remove("app")

	if tree.Search("app") {
		t.Error("Expected not to find 'app' after deletion, but it was found.")
	}

	tree.Remove("apple")

	if tree.Search("apple") {
		t.Error("Expected not to find 'apple' after deletion, but it was found.")
	}

	tree.Remove("application")

	if tree.Search("application") {
		t.Error("Expected not to find 'application' after deletion, but it was found.")
	}

	tree.Remove("apply")

	if tree.Search("apply") {
		t.Error("Expected not to find 'apply' after deletion, but it was found.")
	}

	// Test deleting a non-existing string
	tree.Remove("non-existing")

	if tree.Search("non-existing") {
		t.Error("Expected not to find 'non-existing', but it was found.")
	}
}
