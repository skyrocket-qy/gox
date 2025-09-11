package countingbloomfilter

import (
	"testing"
)

func TestCountingBloomFilter_AddContainsRemove(t *testing.T) {
	filter := New(100, 0.01) // Capacity for 100 items, 1% false positive rate

	itemsToAdd := []string{"apple", "banana", "cherry"}
	for _, item := range itemsToAdd {
		filter.Add([]byte(item))
	}

	// Check for items that should be present
	for _, item := range itemsToAdd {
		if !filter.Contains([]byte(item)) {
			t.Errorf("Expected %s to be in the filter, but it was not", item)
		}
	}

	// Remove an item
	filter.Remove([]byte("apple"))

	// Check if removed item is no longer present (or less likely to be)
	if filter.Contains([]byte("apple")) {
		t.Errorf("Expected apple NOT to be in the filter after removal, but it was")
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
