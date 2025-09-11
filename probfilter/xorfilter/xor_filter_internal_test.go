package xorfilter

import (
	"testing"
)

func TestHashKey(t *testing.T) {
	filter := &XorFilter{}
	hash1 := filter.hashKey(1, 1)

	hash2 := filter.hashKey(1, 1)
	if hash1 != hash2 {
		t.Errorf("Expected same hash for same key and seed, but got %d and %d", hash1, hash2)
	}

	hash3 := filter.hashKey(1, 2)
	if hash1 == hash3 {
		t.Errorf("Expected different hash for different seed, but got same hash %d", hash1)
	}

	hash4 := filter.hashKey(2, 1)
	if hash1 == hash4 {
		t.Errorf("Expected different hash for different key, but got same hash %d", hash1)
	}
}
