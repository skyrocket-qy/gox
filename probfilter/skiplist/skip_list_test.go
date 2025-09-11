package skiplist_test

import (
	"testing"

	"github.com/skyrocket-qy/gox/probfilter/skiplist"
)

func TestSkipListBasic(t *testing.T) {
	sl := skiplist.NewSkipList()

	// Test Insert and Search
	valuesToInsert := []int{10, 20, 5, 15, 25}
	for _, val := range valuesToInsert {
		sl.Insert(val)
	}

	for _, val := range valuesToInsert {
		if !sl.Search(val) {
			t.Errorf("Expected to find %d, but didn't", val)
		}
	}

	// Test Search for non-existent values
	valuesToNotFind := []int{1, 12, 30}
	for _, val := range valuesToNotFind {
		if sl.Search(val) {
			t.Errorf("Did not expect to find %d, but did", val)
		}
	}

	// Test Delete
	sl.Delete(15)
	if sl.Search(15) {
		t.Errorf("Expected 15 to be deleted, but it's still found")
	}

	sl.Delete(5)
	if sl.Search(5) {
		t.Errorf("Expected 5 to be deleted, but it's still found")
	}

	// Test deleting a non-existent value
	sl.Delete(100)

	// Verify remaining values
	if !sl.Search(10) || !sl.Search(20) || !sl.Search(25) {
		t.Errorf("Remaining values not found after deletion")
	}

	if sl.Search(15) || sl.Search(5) {
		t.Errorf("Deleted values still found")
	}
}
