package cuckoofilter

import (
	"testing"
)

func TestCuckooFilter_AddContainsDelete(t *testing.T) {
	filter := New(100, 4) // Capacity for 100 items, 4 entries per bucket

	itemsToAdd := []string{"apple", "banana", "cherry"}
	for _, item := range itemsToAdd {
		if !filter.Add([]byte(item)) {
			t.Fatalf("Failed to add %s to the filter", item)
		}
	}

	// Check for items that should be present
	for _, item := range itemsToAdd {
		if !filter.Contains([]byte(item)) {
			t.Errorf("Expected %s to be in the filter, but it was not", item)
		}
	}

	// Delete an item
	if !filter.Delete([]byte("apple")) {
		t.Errorf("Failed to delete apple from the filter")
	}

	// Check if deleted item is no longer present
	if filter.Contains([]byte("apple")) {
		t.Errorf("Expected apple NOT to be in the filter after deletion, but it was")
	}

	// Check for items that should still be present
	if !filter.Contains([]byte("banana")) {
		t.Errorf("Expected banana to be in the filter, but it was not")
	}

	// Check for items that should definitely not be present
	itemsNotPresent := []string{"grape", "kiwi"}
	for _, item := range itemsNotPresent {
		if filter.Contains([]byte(item)) {
			t.Errorf("Expected %s NOT to be in the filter (false positive), but it was", item)
		}
	}
}
