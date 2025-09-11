package minhashlsh

import (
	"testing"
	"math"
)

func TestMinHashLSH_JaccardSimilarity(t *testing.T) {
	setA := []string{"apple", "banana", "cherry", "date", "elderberry"}
	setB := []string{"apple", "banana", "fig", "grape", "elderberry"}

	mh1 := NewMinHash(128) // 128 permutations
	mh1.GenerateSignature(setA)

	mh2 := NewMinHash(128)
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

	// MinHash provides an estimate, so we check if it's within a reasonable range.
	allowedError := 0.20 // 20% error margin for this basic test

	if math.Abs(estimatedJaccard-trueJaccard) > allowedError {
		t.Errorf("Estimated Jaccard %.2f is outside allowed error margin for true Jaccard %.2f (allowed error: %.2f)", estimatedJaccard, trueJaccard, allowedError)
	}
}
