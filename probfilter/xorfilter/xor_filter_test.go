package xorfilter_test

import (
	"testing"

	"xorfilter"
)

func TestXorFilterBasic(t *testing.T) {
	keys := []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	filter, err := xorfilter.New(keys)
	if err != nil {
		t.Fatalf("Failed to create XorFilter: %v", err)
	}

	// Test Contains for existing keys
	for _, key := range keys {
		if !filter.Contains(key) {
			t.Errorf("XorFilter should contain %d, but it doesn't", key)
		}
	}

	// Test Contains for non-existing keys
	nonExistingKeys := []uint64{11, 12, 13, 14, 15}
	for _, key := range nonExistingKeys {
		if filter.Contains(key) {
			t.Errorf("XorFilter should NOT contain %d, but it does (false positive)", key)
		}
	}
}
