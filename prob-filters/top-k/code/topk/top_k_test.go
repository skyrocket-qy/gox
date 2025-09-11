package topk

import (
	"testing"
)

func TestTopK_AddGetTopK(t *testing.T) {
	k := 3 // Track top 3 items
	topKTracker := New(k, 0.01, 0.01) // epsilon=0.01, delta=0.01

	items := []string{
		"apple", "banana", "apple", "cherry", "banana",
		"apple", "fig", "banana", "grape", "apple",
		"cherry", "banana", "apple", "kiwi", "banana",
		"apple", "banana", "apple", "cherry", "banana",
		"apple", "fig", "banana", "grape", "apple",
	}

	for _, item := range items {
		topKTracker.Add(item)
	}

	currentTopK := topKTracker.GetTopK()

	// True frequencies for comparison
	trueFrequencies := make(map[string]int)
	for _, item := range items {
		trueFrequencies[item]++
	}

	// Check if the top K items are as expected
	// This test is simplified and assumes the top K items are exactly as expected
	// based on the input data. In a real scenario, due to the probabilistic nature
	// of Count-Min Sketch, there might be slight deviations.

	// Expected top 3 items based on true frequencies:
	// apple: 9
	// banana: 8
	// cherry: 3

	if len(currentTopK) != k {
		t.Errorf("Expected %d top items, got %d", k, len(currentTopK))
	}

	// Check specific items and their estimated counts
	// Note: The estimated counts from Count-Min Sketch can be overestimates.
	// We are primarily checking if the correct items are identified as top K.

	// Helper to find an item in the topK list
	findItem := func(val string) (Item, bool) {
		for _, item := range currentTopK {
			if item.Value == val {
				return item, true
			}
		}
		return Item{}, false
	}

	expectedTopItems := map[string]int{
		"apple":  9,
		"banana": 8,
		"cherry": 3,
	}

	for expectedItem, expectedCount := range expectedTopItems {
		if item, found := findItem(expectedItem); !found {
			t.Errorf("Expected top item %s not found in the result", expectedItem)
		} else {
			// Check if estimated count is at least the true count (due to overestimation)
			if item.Count < uint32(expectedCount) {
				t.Errorf("Estimated count for %s (%d) is less than true count (%d)", expectedItem, item.Count, expectedCount)
			}
		}
	}
}
