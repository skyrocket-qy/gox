package main

import (
	"fmt"
	"github.com/skyrocket-qy/prob-filters/counting-bloom-filter/code/countingbloomfilter"
)

func main() {
	filter := countingbloomfilter.New(1000, 0.01) // Capacity for 1000 items, 1% false positive rate

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

	fmt.Println("\nRemoving 'apple' and 'grape':")
	filter.Remove([]byte("apple"))
	filter.Remove([]byte("grape"))

	fmt.Println("\nChecking items after removal:")
	for _, item := range checkItems {
		if filter.Contains([]byte(item)) {
			fmt.Printf("'%s' might be in the set.\n", item)
		} else {
			fmt.Printf("'%s' is definitely NOT in the set.\n", item)
		}
	}
}
