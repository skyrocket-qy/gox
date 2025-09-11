package countminsketch

import (
	"testing"
)

func TestCountMinSketch_AddEstimate(t *testing.T) {
	cms := New(0.001, 0.01) // epsilon=0.001, delta=0.01

	items := []string{
		"apple", "banana", "apple", "cherry", "banana",
		"apple", "fig", "banana", "grape", "apple",
		"cherry", "banana", "apple", "kiwi", "banana",
	}

	for _, item := range items {
		cms.Add([]byte(item))
	}

	// True frequencies
	trueFrequencies := make(map[string]int)
	for _, item := range items {
		trueFrequencies[item]++
	}

	// Check estimated frequencies against true frequencies
	for item, trueCount := range trueFrequencies {
		estimatedCount := cms.Estimate([]byte(item))
		// Count-Min Sketch can overestimate, but should not underestimate
		if estimatedCount < uint32(trueCount) {
			t.Errorf("For item %s: Estimated count %d is less than true count %d", item, estimatedCount, trueCount)
		}
		// We can also check for a reasonable upper bound, but that's more complex
	}

	// Check for an item not added
	itemNotAdded := "orange"
	estimatedCountNotAdded := cms.Estimate([]byte(itemNotAdded))
	if estimatedCountNotAdded > 0 {
		t.Errorf("For item %s: Estimated count %d should be 0, but it's not", itemNotAdded, estimatedCountNotAdded)
	}
}
