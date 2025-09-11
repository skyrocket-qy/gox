package main

import (
	"fmt"
	"github.com/skyrocket-qy/prob-filters/quotient-filter/code/quotientfilter"
)

func main() {
	filter := quotientfilter.New(1000) // Capacity for 1000 items

	itemsToAdd := []string{"apple", "banana", "cherry", "date", "elderberry"}
	for _, item := range itemsToAdd {
		if filter.Add([]byte(item)) {
			fmt.Printf("Added: %s\n", item)
		} else {
			fmt.Printf("Failed to add: %s (filter might be full)\n", item)
		}
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

	// Note: Deletion is not implemented in the basic example.
	fmt.Println("\nDeletion is not implemented in the basic Quotient Filter example.")
}
