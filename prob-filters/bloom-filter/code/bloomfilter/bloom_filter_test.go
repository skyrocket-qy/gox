package bloomfilter

import (
	"testing"
)

func TestBloomFilter_AddContains(t *testing.T) {
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

	// Check for items that should definitely not be present
	itemsNotPresent := []string{"grape", "kiwi", "orange"}
	for _, item := range itemsNotPresent {
		if filter.Contains([]byte(item)) {
			t.Errorf("Expected %s NOT to be in the filter (false positive), but it was", item)
		}
	}
}