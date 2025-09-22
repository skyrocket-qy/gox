package minhashlsh_test

import (
	"log"
	"testing"

	"github.com/skyrocket-qy/gox/probfilter/minhashlsh"
	"github.com/stretchr/testify/require"
)

func TestMinHashLSHBasic(t *testing.T) {
	// Create a MinHasher
	mh := minhashlsh.NewMinHasher(128) // 128 permutations

	// Define some sets
	set1 := []string{"apple", "banana", "cherry", "date", "elderberry"}
	set2 := []string{"apple", "banana", "cherry", "fig", "grape"}
	set3 := []string{"kiwi", "lemon", "mango"}

	// Compute signatures
	sig1 := mh.Signature(set1)
	sig2 := mh.Signature(set2)
	sig3 := mh.Signature(set3)

	// Test Jaccard Similarity (approximate)
	sim12 := minhashlsh.JaccardSimilarity(sig1, sig2)
	sim13 := minhashlsh.JaccardSimilarity(sig1, sig3)

	log.Printf("Jaccard Similarity (set1, set2): %.2f\n", sim12)
	log.Printf("Jaccard Similarity (set1, set3): %.2f\n", sim13)

	// Expect set1 and set2 to be more similar than set1 and set3
	if sim12 < sim13 {
		t.Errorf("Expected sim12 (%.2f) to be greater than sim13 (%.2f)", sim12, sim13)
	}

	// Create an LSH component
	lsh, err := minhashlsh.NewLSH(16, 8) // 16 bands, 8 rows per band (16*8 = 128 permutations)
	if err != nil {
		t.Fatalf("Failed to create LSH: %v", err)
	}

	// Add documents to LSH
	err = lsh.Add("doc1", sig1)
	require.NoError(t, err)
	err = lsh.Add("doc2", sig2)
	require.NoError(t, err)
	err = lsh.Add("doc3", sig3)
	require.NoError(t, err)

	// Query for similar documents
	candidates1 := lsh.Query("doc1", sig1)
	candidates2 := lsh.Query("doc2", sig2)
	candidates3 := lsh.Query("doc3", sig3)

	log.Printf("Candidates for doc1: %v\n", candidates1)
	log.Printf("Candidates for doc2: %v\n", candidates2)
	log.Printf("Candidates for doc3: %v\n", candidates3)

	// Expect doc1 and doc2 to be candidates for each other
	if !contains(candidates1, "doc2") {
		t.Errorf("Expected doc1 candidates to include doc2, got %v", candidates1)
	}

	if !contains(candidates2, "doc1") {
		t.Errorf("Expected doc2 candidates to include doc1, got %v", candidates2)
	}

	// Expect doc3 to have fewer or no candidates from doc1/doc2
	if contains(candidates3, "doc1") || contains(candidates3, "doc2") {
		t.Logf("Unexpected candidates for doc3: %v", candidates3)
	}
}

func TestNewLSH_Error(t *testing.T) {
	_, err := minhashlsh.NewLSH(0, 8)
	if err == nil {
		t.Error("Expected error for numBands <= 0, but got nil")
	}

	_, err = minhashlsh.NewLSH(16, 0)
	if err == nil {
		t.Error("Expected error for numRows <= 0, but got nil")
	}
}

func TestLSH_Add_Error(t *testing.T) {
	lsh, _ := minhashlsh.NewLSH(16, 8)

	err := lsh.Add("doc1", make([]uint64, 127))
	if err == nil {
		t.Error("Expected error for signature length mismatch, but got nil")
	}
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}

	return false
}
