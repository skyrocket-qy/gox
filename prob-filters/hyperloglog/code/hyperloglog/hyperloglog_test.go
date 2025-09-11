package hyperloglog

import (
	"testing"
	"math"
)

func TestHyperLogLog_Estimate(t *testing.T) {
	hll := New(14) // 14-bit precision

	items := []string{
		"apple", "banana", "cherry", "date", "elderberry",
		"apple", "fig", "grape", "honeydew", "ice cream",
		"banana", "jicama", "kiwi", "lemon", "mango",
		"cherry", "nectarine", "orange", "pear", "quince",
	}

	for _, item := range items {
		hll.Add([]byte(item))
	}

	// Calculate true unique count
	trueUnique := make(map[string]struct{})
	for _, item := range items {
		trueUnique[item] = struct{}{}
	}
	actualCount := float64(len(trueUnique))

	estimatedCount := hll.Estimate()

	// HyperLogLog provides an estimate, so we check if it's within a reasonable range.
	// The error rate depends on the precision. For precision 14, it's roughly 0.67%.
	// We'll allow a larger margin for a basic test.
	allowedError := actualCount * 0.10 // 10% error margin for this basic test

	if math.Abs(estimatedCount-actualCount) > allowedError {
		t.Errorf("Estimated count %.2f is outside allowed error margin for true count %.2f (allowed error: %.2f)", estimatedCount, actualCount, allowedError)
	}
}
