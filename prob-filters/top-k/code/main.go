package main

import (
	"fmt"
	"github.com/skyrocket-qy/prob-filters/top-k/code/topk"
)

func main() {
	k := 3 // Track top 3 items
	topKTracker := topk.New(k, 0.01, 0.01) // epsilon=0.01, delta=0.01

	items := []string{
		"apple", "banana", "apple", "cherry", "banana",
		"apple", "fig", "banana", "grape", "apple",
		"cherry", "banana", "apple", "kiwi", "banana",
		"apple", "banana", "apple", "cherry", "banana",
		"apple", "fig", "banana", "grape", "apple",
	}

	fmt.Printf("Adding items to Top-%d tracker:\n", k)
	for _, item := range items {
		topKTracker.Add(item)
		fmt.Printf("Added: %s\n", item)
	}

	fmt.Printf("\nTop %d items:\n", k)
	currentTopK := topKTracker.GetTopK()
	for i, item := range currentTopK {
		fmt.Printf("%d. %s (Estimated Count: %d)\n", i+1, item.Value, item.Count)
	}

	// True frequencies for comparison
	trueFrequencies := make(map[string]int)
	for _, item := range items {
		trueFrequencies[item]++
	}
	fmt.Println("\nTrue frequencies (for comparison):")
	for item, count := range trueFrequencies {
		fmt.Printf("  %s: %d\n", item, count)
	}
}
