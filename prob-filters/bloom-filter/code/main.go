package main

import (
	"fmt"
	"github.com/skyrocket-qy/prob-filters/bloom-filter/code/bloomfilter"
)

func main() {
	filter := bloomfilter.New(1000, 0.01) // Capacity for 1000 items, 1% false positive rate

	itemsToAdd := []string{"apple", "banana", "cherry", "date", "elderberry"}
	for _, item := range itemsToAdd {
		filter.Add([]byte(item))
		fmt.Printf("Added: %s\n", item)
	}

	fmt.Println("\nChecking items:")
	checkItems := []string{"apple", "grape", "cherry", "kiwi", "date"}
	for _, item := range checkItems {
		if filter.Contains([]byte(item)) {
			fmt.Printf("'%s' might be in the set.\n", item)
		} else {
			fmt.Printf("'%s' is definitely NOT in the set.\n", item)
		}
	}

	// Demonstrate a false positive (if it occurs)
	fmt.Println("\nDemonstrating potential false positive:")
	falsePositiveCandidate := "zucchini" // A word not added
	if filter.Contains([]byte(falsePositiveCandidate)) {
		fmt.Printf("'%s' might be in the set (false positive).\n", falsePositiveCandidate)
	} else {
		fmt.Printf("'%s' is definitely NOT in the set.\n", falsePositiveCandidate)
	}
}
