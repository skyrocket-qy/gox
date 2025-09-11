package countingbloomfilter

import (
	"hash"
	"hash/fnv"
	"math"
)

// CountingBloomFilter represents a probabilistic data structure that supports deletions.
type CountingBloomFilter struct {
	counts     []uint8 // Using uint8 for counters, assuming max count of 255
	hashFuncs  []hash.Hash64
	numHashes  int
	capacity   int
	falsePosRate float64
}

// New creates a new CountingBloomFilter with a given capacity and desired false positive rate.
func New(capacity int, falsePosRate float64) *CountingBloomFilter {
	// Optimal number of bits (m) and hash functions (k)
	// m = -(n * ln(p)) / (ln(2)^2)
	// k = (m / n) * ln(2)
	m := int(-float64(capacity) * (math.Log(falsePosRate) / (math.Log(2) * math.Log(2))))
	k := int(float64(m) / float64(capacity) * math.Log(2))

	if m < 1 {
		m = 1
	}
	if k < 1 {
		k = 1
	}

	hashFuncs := make([]hash.Hash64, k)
	for i := 0; i < k; i++ {
		hashFuncs[i] = fnv.New64a() // Using FNV-1a for simplicity
	}

	return &CountingBloomFilter{
		counts:     make([]uint8, m),
		hashFuncs:  hashFuncs,
		numHashes:  k,
		capacity:   capacity,
		falsePosRate: falsePosRate,
	}
}

// Add inserts an element into the Counting Bloom filter.
func (cbf *CountingBloomFilter) Add(data []byte) {
	for _, hf := range cbf.hashFuncs {
		hf.Reset()
		hf.Write(data)
		index := hf.Sum64() % uint64(len(cbf.counts))
		if cbf.counts[index] < math.MaxUint8 { // Prevent overflow
			cbf.counts[index]++
		}
	}
}

// Remove deletes an element from the Counting Bloom filter.
func (cbf *CountingBloomFilter) Remove(data []byte) {
	// Before removing, check if the item is present to avoid underflow
	// This is a simplification; a true removal would require knowing if the item was actually added.
	// For a robust implementation, one might check if Contains returns true first.
	for _, hf := range cbf.hashFuncs {
		hf.Reset()
		hf.Write(data)
		index := hf.Sum64() % uint64(len(cbf.counts))
		if cbf.counts[index] > 0 {
			cbf.counts[index]--
		}
	}
}

// Contains checks if an element might be in the Counting Bloom filter.
// Returns true if the element might be present, false if definitely not.
func (cbf *CountingBloomFilter) Contains(data []byte) bool {
	for _, hf := range cbf.hashFuncs {
		hf.Reset()
		hf.Write(data)
		index := hf.Sum64() % uint64(len(cbf.counts))
		if cbf.counts[index] == 0 {
			return false // Definitely not in the set
		}
	}
	return true // Might be in the set (possible false positive)
}