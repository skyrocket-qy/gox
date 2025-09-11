package countminsketch

import (
	"hash/fnv"
	"math"
)

// CountMinSketch represents a probabilistic data structure for frequency estimation.
type CountMinSketch struct {
	// d (depth): number of hash functions
	// w (width): number of buckets per hash function
	counts [][]uint32
	d, w   int
	seeds  []uint64 // Seeds for hash functions
}

// New creates a new CountMinSketch.
// epsilon: controls the error in frequency estimates (e.g., 0.001).
// delta: controls the probability of the error exceeding epsilon (e.g., 0.01).
func New(epsilon, delta float64) *CountMinSketch {
	d := int(math.Ceil(math.Log(1 / delta)))
	w := int(math.Ceil(math.E / epsilon))

	counts := make([][]uint32, d)
	for i := range counts {
		counts[i] = make([]uint32, w)
	}

	seeds := make([]uint64, d)
	// Using simple sequential seeds for demonstration.
	// In practice, cryptographically secure random seeds are better.
	for i := 0; i < d; i++ {
		seeds[i] = uint64(i + 1) // Simple seed
	}

	return &CountMinSketch{
		counts: counts,
		d:      d,
		w:      w,
		seeds:  seeds,
	}
}

// hash computes a hash for the given data and seed.
func (cms *CountMinSketch) hash(data []byte, seed uint64) uint32 {
	h := fnv.New64a()
	h.Write(data)
	h.Write([]byte{byte(seed)}) // Incorporate seed
	return uint32(h.Sum64())
}

// Add increments the count for an item.
func (cms *CountMinSketch) Add(data []byte) {
	for i := 0; i < cms.d; i++ {
		idx := cms.hash(data, cms.seeds[i]) % uint32(cms.w)
		cms.counts[i][idx]++
	}
}

// Estimate returns the estimated frequency of an item.
func (cms *CountMinSketch) Estimate(data []byte) uint32 {
	minCount := uint32(math.MaxUint32) // Initialize with max possible value

	for i := 0; i < cms.d; i++ {
		idx := cms.hash(data, cms.seeds[i]) % uint32(cms.w)
		if cms.counts[i][idx] < minCount {
			minCount = cms.counts[i][idx]
		}
	}
	return minCount
}