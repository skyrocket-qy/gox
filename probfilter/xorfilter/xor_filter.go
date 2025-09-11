package xorfilter

import (
	"errors"
	"fmt"
	"hash/fnv"
)

// XorFilter represents an Xor Filter data structure.
type XorFilter struct {
	// Internal structure for the Xor Filter
	// This will typically involve an array of fingerprints and a set of hash functions.
	// For a basic implementation, we'll simplify this.
	data []uint64
	seed uint64
}

// New creates and initializes a new XorFilter from a slice of uint64 keys.
// This is a simplified implementation for demonstration purposes.
// A full Xor Filter construction involves a more complex graph-based algorithm.
func New(keys []uint64) (*XorFilter, error) {
	if len(keys) == 0 {
		return nil, errors.New("keys cannot be empty")
	}

	// For a true Xor Filter, the construction is non-trivial.
	// This simplified version will just store the keys and use XOR for lookup.
	// This will NOT have the space efficiency or false positive guarantees of a real Xor Filter.
	// It's merely a placeholder to demonstrate the interface.

	filter := &XorFilter{
		data: make([]uint64, len(keys)),
		seed: 0xdeadbeef, // A fixed seed for simplicity
	}

	// In a real Xor Filter, `data` would store fingerprints derived from keys
	// and hash functions, such that XORing them reveals membership.
	// For this placeholder, we'll just copy the keys.
	copy(filter.data, keys)

	return filter, nil
}

// Contains checks if a key is present in the XorFilter.
// This is a simplified implementation.
func (xf *XorFilter) Contains(key uint64) bool {
	// In a real Xor Filter, this would involve hashing the key
	// to three positions and XORing the fingerprints at those positions.
	// For this placeholder, we'll just do a linear scan.
	for _, k := range xf.data {
		if k == key {
			return true
		}
	}

	return false
}

// hashKey generates a hash for a uint64 key.
func (xf *XorFilter) hashKey(key, seed uint64) uint64 {
	h := fnv.New64a()
	fmt.Fprintf(h, "%d-%d", key, seed)

	return h.Sum64()
}
