package randomizedset

import "testing"

func TestRandomizedSet(t *testing.T) {
	rs := Constructor()

	if !rs.Insert(1) {
		t.Error("Insert(1) should return true")
	}

	if rs.Insert(1) {
		t.Error("Insert(1) again should return false")
	}

	if !rs.Insert(2) {
		t.Error("Insert(2) should return true")
	}

	// GetRandom is hard to test deterministically, so we just check if it returns a valid element
	randomVal := rs.GetRandom()
	if randomVal != 1 && randomVal != 2 {
		t.Errorf("GetRandom() returned an unexpected value: %d", randomVal)
	}

	if !rs.Remove(1) {
		t.Error("Remove(1) should return true")
	}

	if rs.Remove(1) {
		t.Error("Remove(1) again should return false")
	}

	if rs.Insert(2) {
		t.Error("Insert(2) again should return false")
	}

	if rs.GetRandom() != 2 {
		t.Errorf("GetRandom() should return 2, got %d", rs.GetRandom())
	}
}
