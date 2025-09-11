package bloomfilter

import (
	"hash/fnv"
	"math"
)

// BloomFilter represents a probabilistic data structure for set membership testing.
type BloomFilter struct {
	bits         []bool
	numHashes    int
	size         uint64
	capacity     int
	falsePosRate float64
}

// New creates a new BloomFilter with a given capacity and desired false positive rate.
func New(capacity int, falsePosRate float64) *BloomFilter {
	// Optimal number of bits (m) and hash functions (k)
	// m = -(n * ln(p)) / (ln(2)^2)
	// k = (m / n) * ln(2)
	mFloat := -float64(capacity) * math.Log(falsePosRate) / (math.Ln2 * math.Ln2)
	kFloat := mFloat / float64(capacity) * math.Ln2

	m := int(math.Ceil(mFloat))
	k := int(math.Ceil(kFloat))

	if m < 1 {
		m = 1
	}
	if k < 1 {
		k = 1
	}

	return &BloomFilter{
		bits:         make([]bool, m),
		numHashes:    k,
		size:         uint64(m),
		capacity:     capacity,
		falsePosRate: falsePosRate,
	}
}

// baseHashes computes the two base hashes for double hashing.
func (bf *BloomFilter) baseHashes(data []byte) (uint64, uint64) {
	h1 := fnv.New64a()
	h1.Write(data)
	hash1 := h1.Sum64()

	h2 := fnv.New64()
	h2.Write(data)
	hash2 := h2.Sum64()

	return hash1, hash2
}

// Add inserts an element into the Bloom filter using double hashing.
func (bf *BloomFilter) Add(data []byte) {
	h1, h2 := bf.baseHashes(data)
	for i := 0; i < bf.numHashes; i++ {
		index := (h1 + uint64(i)*h2) % bf.size
		bf.bits[index] = true
	}
}

// Contains checks if an element might be in the Bloom filter using double hashing.
// Returns true if the element might be present, false if definitely not.
func (bf *BloomFilter) Contains(data []byte) bool {
	h1, h2 := bf.baseHashes(data)
	for i := 0; i < bf.numHashes; i++ {
		index := (h1 + uint64(i)*h2) % bf.size
		if !bf.bits[index] {
			return false // Definitely not in the set
		}
	}
	return true // Might be in the set (possible false positive)
}
