package quotientfilter

import (
	"testing"
)

func TestQuotientFilter_AddContains(t *testing.T) {
	filter := New(100) // Capacity for 100 items

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

	// Check for items that should definitely not be present
	itemsNotPresent := []string{"grape", "kiwi", "orange"}
	for _, item := range itemsNotPresent {
		if filter.Contains([]byte(item)) {
			t.Errorf("Expected %s NOT to be in the filter (false positive), but it was", item)
		}
	}
}
