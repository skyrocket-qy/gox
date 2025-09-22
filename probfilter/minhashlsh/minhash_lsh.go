package minhashlsh

import (
	"errors"
	"fmt"
	"hash/fnv"
	"math"
	"sort"
	"strconv"
)

// MinHasher generates MinHash signatures for sets.
type MinHasher struct {
	numPermutations int
	seeds           []uint64
}

// NewMinHasher creates a new MinHasher with a specified number of permutations.
func NewMinHasher(numPermutations int) *MinHasher {
	seeds := make([]uint64, numPermutations)
	for i := range numPermutations {
		seeds[i] = uint64(i)*uint64(2) + 1 // Simple seeds
	}

	return &MinHasher{
		numPermutations: numPermutations,
		seeds:           seeds,
	}
}

// Signature computes the MinHash signature for a given set of elements.
func (mh *MinHasher) Signature(elements []string) []uint64 {
	signature := make([]uint64, mh.numPermutations)
	for i := range signature {
		signature[i] = math.MaxUint64 // Initialize with max value
	}

	for _, element := range elements {
		for i := range mh.numPermutations {
			h := fnv.New64a()
			_, _ = h.Write([]byte(element))
			fmt.Fprintf(h, "%d", mh.seeds[i]) // Incorporate seed

			hashValue := h.Sum64()
			if hashValue < signature[i] {
				signature[i] = hashValue
			}
		}
	}

	return signature
}

// JaccardSimilarity computes the Jaccard similarity between two MinHash signatures.
func JaccardSimilarity(sig1, sig2 []uint64) float64 {
	if len(sig1) != len(sig2) {
		return 0.0
	}

	intersection := 0

	for i := range sig1 {
		if sig1[i] == sig2[i] {
			intersection++
		}
	}

	return float64(intersection) / float64(len(sig1))
}

// LSH represents a basic Locality Sensitive Hashing component.
type LSH struct {
	numBands int
	numRows  int
	buckets  map[string][]string // Map hash of band to list of document IDs
}

// NewLSH creates a new LSH component.
// numBands: Number of bands for LSH.
// numRows: Number of rows per band.
func NewLSH(numBands, numRows int) (*LSH, error) {
	if numBands <= 0 || numRows <= 0 {
		return nil, errors.New("numBands and numRows must be greater than 0")
	}

	return &LSH{
		numBands: numBands,
		numRows:  numRows,
		buckets:  make(map[string][]string),
	}, nil
}

// Add adds a document's signature to the LSH buckets.
func (lsh *LSH) Add(docID string, signature []uint64) error {
	if len(signature) != lsh.numBands*lsh.numRows {
		return errors.New("signature length does not match LSH configuration")
	}

	for b := range lsh.numBands {
		band := signature[b*lsh.numRows : (b+1)*lsh.numRows]
		// Hash the band to get a bucket key
		h := fnv.New64a()
		for _, val := range band {
			fmt.Fprintf(h, "%d", val)
		}

		bucketKey := strconv.FormatUint(h.Sum64(), 16)

		lsh.buckets[bucketKey] = append(lsh.buckets[bucketKey], docID)
	}

	return nil
}

// Query returns candidate similar documents for a given document ID.
func (lsh *LSH) Query(docID string, signature []uint64) []string {
	candidates := make(map[string]struct{})

	for b := range lsh.numBands {
		band := signature[b*lsh.numRows : (b+1)*lsh.numRows]
		// Hash the band to get a bucket key
		h := fnv.New64a()
		for _, val := range band {
			fmt.Fprintf(h, "%d", val)
		}

		bucketKey := strconv.FormatUint(h.Sum64(), 16)

		if docs, ok := lsh.buckets[bucketKey]; ok {
			for _, d := range docs {
				if d != docID {
					candidates[d] = struct{}{} // Add to candidates set
				}
			}
		}
	}

	result := make([]string, 0, len(candidates))
	for c := range candidates {
		result = append(result, c)
	}

	sort.Strings(result)

	return result
}
