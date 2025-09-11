package main

import (
	"fmt"
	"github.com/skyrocket-qy/prob-filters/minhash-lsh/code/minhashlsh"
)

func main() {
	// Example 1: Jaccard Similarity Estimation
	fmt.Println("---", "Jaccard Similarity Estimation", "---")

	setA := []string{"apple", "banana", "cherry", "date", "elderberry"}
	setB := []string{"apple", "banana", "fig", "grape", "elderberry"}

	mh1 := minhashlsh.NewMinHash(128) // 128 permutations
	mh1.GenerateSignature(setA)

	mh2 := minhashlsh.NewMinHash(128)
	mh2.GenerateSignature(setB)

	estimatedJaccard := mh1.JaccardSimilarity(mh2)

	// Calculate true Jaccard Similarity
	intersection := 0
	union := make(map[string]struct{})
	for _, item := range setA {
		union[item] = struct{}{}
	}
	for _, item := range setB {
		union[item] = struct{}{}
		found := false
		for _, aItem := range setA {
			if aItem == item {
				found = true
				break
			}
		}
		if found {
			intersection++
		}
	}

trueJaccard := float64(intersection) / float64(len(union))

	fmt.Printf("Set A: %v\n", setA)
	fmt.Printf("Set B: %v\n", setB)
	fmt.Printf("True Jaccard Similarity: %.2f\n", trueJaccard)
	fmt.Printf("Estimated Jaccard Similarity (MinHash): %.2f\n", estimatedJaccard)

	// Example 2: LSH Bucketing (Conceptual)
	fmt.Println("\n---", "LSH Bucketing (Conceptual)", "---")
	lsh := minhashlsh.NewLSH(4, 32) // 4 bands, 32 rows per band (total 128 permutations)

	// Generate a third signature for LSH demonstration
	setC := []string{"apple", "banana", "cherry", "date", "fig"}
	mh3 := minhashlsh.NewMinHash(128)
	mh3.GenerateSignature(setC)

	buckets1 := lsh.GetLSHBuckets(mh1)
	buckets2 := lsh.GetLSHBuckets(mh2)
	buckets3 := lsh.GetLSHBuckets(mh3)

	fmt.Println("LSH Buckets for Signature 1 (Set A):")
	for i, b := range buckets1 {
		fmt.Printf("  Band %d: %v\n", i+1, b)
	}
	fmt.Println("LSH Buckets for Signature 2 (Set B):")
	for i, b := range buckets2 {
		fmt.Printf("  Band %d: %v\n", i+1, b)
	}
	fmt.Println("LSH Buckets for Signature 3 (Set C):")
	for i, b := range buckets3 {
		fmt.Printf("  Band %d: %v\n", i+1, b)
	}

	fmt.Println("\nNote: In a real LSH system, signatures with matching bands would be considered candidate pairs for more precise similarity checks.")
}
