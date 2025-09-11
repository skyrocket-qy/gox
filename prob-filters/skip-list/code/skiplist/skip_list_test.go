package skiplist

import (
	"testing"
)

func TestSkipList_InsertSearchDelete(t *testing.T) {
	sl := New()

	itemsToAdd := []int{10, 5, 20, 15, 7, 25, 12}
	for _, item := range itemsToAdd {
		sl.Insert(item)
	}

	// Check for items that should be present
	for _, item := range itemsToAdd {
		if !sl.Search(item) {
			t.Errorf("Expected %d to be in the skip list, but it was not", item)
		}
	}

	// Check for items that should not be present
	itemsNotPresent := []int{1, 100, 0}
	for _, item := range itemsNotPresent {
		if sl.Search(item) {
			t.Errorf("Expected %d NOT to be in the skip list, but it was", item)
		}
	}

	// Delete an item that exists
	if !sl.Delete(15) {
		t.Errorf("Failed to delete 15 from the skip list")
	}
	if sl.Search(15) {
		t.Errorf("Expected 15 NOT to be in the skip list after deletion, but it was")
	}

	// Delete an item that does not exist
	if sl.Delete(99) {
		t.Errorf("Expected 99 NOT to be deleted (not found), but it was")
	}

	// Check remaining items
	if !sl.Search(10) {
		t.Errorf("Expected 10 to be in the skip list, but it was not")
	}
}
