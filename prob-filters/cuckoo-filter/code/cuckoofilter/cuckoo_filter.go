package cuckoofilter

import (
	"hash/fnv"
	"math/rand"
	"time"
)

const (
	// MaxDisplacements is the maximum number of displacements before giving up on insertion.
	MaxDisplacements = 500 // A common heuristic value
	// FingerprintSize is the size of the fingerprint in bytes.
	FingerprintSize = 1 // Using 1 byte for simplicity, typically 2-4 bytes
)

// CuckooFilter represents a probabilistic data structure for set membership testing with deletions.
type CuckooFilter struct {
	buckets []bucket
	numBuckets int
	entriesPerBucket int
	numItems int
	seed uint64
}

// bucket is a slice of fingerprints.
type bucket []byte

// New creates a new CuckooFilter.
// capacity: approximate number of items to store.
// entriesPerBucket: number of entries per bucket (e.g., 4).
func New(capacity, entriesPerBucket int) *CuckooFilter {
	numBuckets := nextPowerOf2(capacity / entriesPerBucket)
	if numBuckets == 0 {
		numBuckets = 1 // Ensure at least one bucket
	}

	buckets := make([]bucket, numBuckets)
	for i := range buckets {
		buckets[i] = make(bucket, entriesPerBucket * FingerprintSize)
	}

	return &CuckooFilter{
		buckets: buckets,
		numBuckets: numBuckets,
		entriesPerBucket: entriesPerBucket,
		numItems: 0,
		seed: uint64(time.Now().UnixNano()), // Simple seed for hash functions
	}
}

// nextPowerOf2 returns the smallest power of 2 greater than or equal to n.
func nextPowerOf2(n int) int {
	if n == 0 {
		return 1
	}
	n--
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	n++
	return n
}

// generateFingerprint generates a fingerprint for the given data.
func (cf *CuckooFilter) generateFingerprint(data []byte) byte {
	h := fnv.New32a()
	h.Write(data)
	return byte(h.Sum32() & 0xFF) // Take lower 8 bits for 1-byte fingerprint
}

// getAltIndex calculates the alternate bucket index.
func (cf *CuckooFilter) getAltIndex(idx int, fp byte) int {
	h := fnv.New32a()
	h.Write([]byte{fp})
	h.Write([]byte{byte(idx)})
	return (idx ^ int(h.Sum32())) % cf.numBuckets
}

// Add inserts an item into the filter. Returns true on success, false if filter is full.
func (cf *CuckooFilter) Add(data []byte) bool {
	fp := cf.generateFingerprint(data)
	h := fnv.New32a()
h.Write(data)
	idx1 := int(h.Sum32() % uint32(cf.numBuckets))
	idx2 := cf.getAltIndex(idx1, fp)

	// Try to insert into either bucket
	if cf.insertToBucket(idx1, fp) || cf.insertToBucket(idx2, fp) {
		cf.numItems++
		return true
	}

	// Both buckets are full, start kicking out
	currentIdx := idx1
	currentFp := fp
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < MaxDisplacements; i++ {
		// Randomly choose an entry to kick out
		entryIndex := r.Intn(cf.entriesPerBucket) * FingerprintSize
		
		// Swap currentFp with a random fingerprint from the bucket
		tempFp := cf.buckets[currentIdx][entryIndex]
		cf.buckets[currentIdx][entryIndex] = currentFp
		currentFp = tempFp

		// Calculate alternate index for the kicked-out fingerprint
		currentIdx = cf.getAltIndex(currentIdx, currentFp)

		// Try to insert the kicked-out fingerprint into its alternate bucket
		if cf.insertToBucket(currentIdx, currentFp) {
			cf.numItems++
			return true
		}
	}
	return false // Failed to add after max displacements
}

// insertToBucket attempts to insert a fingerprint into a bucket.
func (cf *CuckooFilter) insertToBucket(idx int, fp byte) bool {
	for i := 0; i < cf.entriesPerBucket; i++ {
		if cf.buckets[idx][i*FingerprintSize] == 0 { // Assuming 0 means empty
			cf.buckets[idx][i*FingerprintSize] = fp
			return true
		}
	}
	return false
}

// Contains checks if an item might be in the filter.
func (cf *CuckooFilter) Contains(data []byte) bool {
	fp := cf.generateFingerprint(data)
	h := fnv.New32a()
h.Write(data)
	idx1 := int(h.Sum32() % uint32(cf.numBuckets))
	idx2 := cf.getAltIndex(idx1, fp)

	return cf.lookupInBucket(idx1, fp) || cf.lookupInBucket(idx2, fp)
}

// lookupInBucket checks if a fingerprint exists in a bucket.
func (cf *CuckooFilter) lookupInBucket(idx int, fp byte) bool {
	for i := 0; i < cf.entriesPerBucket; i++ {
		if cf.buckets[idx][i*FingerprintSize] == fp {
			return true
		}
	}
	return false
}

// Delete removes an item from the filter. Returns true on success, false if not found.
func (cf *CuckooFilter) Delete(data []byte) bool {
	fp := cf.generateFingerprint(data)
	h := fnv.New32a()
h.Write(data)
	idx1 := int(h.Sum32() % uint32(cf.numBuckets))
	idx2 := cf.getAltIndex(idx1, fp)

	if cf.deleteFromBucket(idx1, fp) || cf.deleteFromBucket(idx2, fp) {
		cf.numItems--
		return true
	}
	return false
}

// deleteFromBucket attempts to delete a fingerprint from a bucket.
func (cf *CuckooFilter) deleteFromBucket(idx int, fp byte) bool {
	for i := 0; i < cf.entriesPerBucket; i++ {
		if cf.buckets[idx][i*FingerprintSize] == fp {
			cf.buckets[idx][i*FingerprintSize] = 0 // Mark as empty
			return true
		}
	}
	return false
}