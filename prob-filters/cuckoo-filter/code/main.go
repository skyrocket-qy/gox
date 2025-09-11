package main

import (
	"fmt"
	"github.com/skyrocket-qy/prob-filters/cuckoo-filter/code/cuckoofilter"
)

func main() {
	filter := cuckoofilter.New(1000, 4) // Capacity for 1000 items, 4 entries per bucket

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

	fmt.Println("\nRemoving 'apple' and 'grape':")
	if filter.Delete([]byte("apple")) {
		fmt.Println("Removed: apple")
	} else {
		fmt.Println("Failed to remove: apple")
	}
	if filter.Delete([]byte("grape")) {
		fmt.Println("Removed: grape")
	} else {
		fmt.Println("Failed to remove: grape")
	}

	fmt.Println("\nChecking items after removal:")
	for _, item := range checkItems {
		if filter.Contains([]byte(item)) {
			fmt.Printf("'%s' might be in the set.\n", item)
		} else {
			fmt.Printf("'%s' is definitely NOT in the set.\n", item)
		}
	}
}
