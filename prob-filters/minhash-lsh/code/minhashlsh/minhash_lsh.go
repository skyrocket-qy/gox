package minhashlsh

import (
	"encoding/binary"
	"hash/fnv"
	"math"
	"sort"
	"strconv"
)

// MinHash represents a MinHash signature for a set.
type MinHash struct {
	numPermutations int
	// Stores the minimum hash values for each permutation
	signature []uint64
}

// NewMinHash creates a new MinHash instance.
// numPermutations: The number of hash functions/permutations to use.
// A higher number leads to better accuracy but larger signatures.
func NewMinHash(numPermutations int) *MinHash {
	return &MinHash{
		numPermutations: numPermutations,
		signature:       make([]uint64, numPermutations),
	}
}

// hashValue computes a hash for the given data and seed.
func hashValue(data []byte, seed uint64) uint64 {
	h := fnv.New64a()
	h.Write(data)
	seedBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(seedBytes, seed)
	h.Write(seedBytes) // Incorporate seed
	return h.Sum64()
}

// GenerateSignature generates the MinHash signature for a given set of items.
func (mh *MinHash) GenerateSignature(items []string) {
	// Initialize signature with max possible hash values
	for i := range mh.signature {
		mh.signature[i] = math.MaxUint64
	}

	// For each item in the set
	for _, item := range items {
		itemBytes := []byte(item)
		// For each permutation (hash function)
		for i := 0; i < mh.numPermutations; i++ {
			// Compute the hash for the item with a unique seed for each permutation
			h := hashValue(itemBytes, uint64(i))
			// Update the signature if this hash is smaller
			if h < mh.signature[i] {
				mh.signature[i] = h
			}
		}
	}
}

// JaccardSimilarity estimates the Jaccard similarity between two MinHash signatures.
func (mh *MinHash) JaccardSimilarity(other *MinHash) float64 {
	if mh.numPermutations != other.numPermutations {
		return 0.0 // Signatures must have the same number of permutations
	}

	matches := 0
	for i := 0; i < mh.numPermutations; i++ {
		if mh.signature[i] == other.signature[i] {
			matches++
		}
	}
	return float64(matches) / float64(mh.numPermutations)
}

// LSH (Locality Sensitive Hashing) for MinHash signatures.
// This is a conceptual representation of LSH banding.
type LSH struct {
	bands int
	rows  int
	// In a real LSH, this would be a collection of hash tables, one for each band.
	// For simplicity, we'll just store the parameters.
}

// NewLSH creates a new LSH instance.
// bands: Number of bands.
// rows: Number of rows per band.
func NewLSH(bands, rows int) *LSH {
	return &LSH{
		bands: bands,
		rows:  rows,
	}
}

// GetLSHBuckets generates LSH buckets for a MinHash signature.
// In a real LSH, these buckets would be used to store and retrieve candidate pairs.
func (lsh *LSH) GetLSHBuckets(mh *MinHash) [][]byte {
	if len(mh.signature) != lsh.bands * lsh.rows {
		// Signature length must match bands * rows
		return nil
	}

	buckets := make([][]byte, lsh.bands)
	for b := 0; b < lsh.bands; b++ {
		bandSignature := make([]byte, 0)
		for r := 0; r < lsh.rows; r++ {
			idx := b*lsh.rows + r
			// Hash the signature part for this row to get a bucket ID
			h := fnv.New64a()
			h.Write([]byte(strconv.FormatUint(mh.signature[idx], 10)))
			bandSignature = append(bandSignature, byte(h.Sum64() & 0xFF)) // Take lower 8 bits
		}
		// Sort the band signature to make it canonical for hashing into a bucket
		sort.Slice(bandSignature, func(i, j int) bool {
			return bandSignature[i] < bandSignature[j]
		})
		buckets[b] = bandSignature
	}
	return buckets
}