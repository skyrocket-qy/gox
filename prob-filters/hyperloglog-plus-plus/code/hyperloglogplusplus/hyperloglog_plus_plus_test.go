package hyperloglogplusplus

import (
	"testing"
	"math"
)

func TestHyperLogLogPlusPlus_EstimateMerge(t *testing.T) {
	hllpp1 := New(14) // 14-bit precision

	items1 := []string{
		"apple", "banana", "cherry", "date", "elderberry",
		"fig", "grape", "honeydew", "ice cream", "jicama",
	}

	for _, item := range items1 {
		hllpp1.Add([]byte(item))
	}

	hllpp2 := New(14)
	items2 := []string{
		"kiwi", "lemon", "mango", "nectarine", "orange",
		"apple", "banana", "pear", "quince", "raspberry",
	}

	for _, item := range items2 {
		hllpp2.Add([]byte(item))
	}

	// Calculate true unique count for merged set
	trueUniqueMerged := make(map[string]struct{})
	for _, item := range items1 {
		trueUniqueMerged[item] = struct{}{}
	}
	for _, item := range items2 {
		trueUniqueMerged[item] = struct{}{}
	}
	actualCount := float64(len(trueUniqueMerged))

	// Merge HLLPP2 into HLLPP1
	err := hllpp1.Merge(hllpp2)
	if err != nil {
		t.Fatalf("Merge failed: %v", err)
	}

	estimatedCount := hllpp1.Estimate()

	// HyperLogLog++ provides an estimate, so we check if it's within a reasonable range.
	allowedError := actualCount * 0.10 // 10% error margin for this basic test

	if math.Abs(estimatedCount-actualCount) > allowedError {
		t.Errorf("Estimated merged count %.2f is outside allowed error margin for true count %.2f (allowed error: %.2f)", estimatedCount, actualCount, allowedError)
	}
}
