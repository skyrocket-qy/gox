package main

import (
	"fmt"
	"github.com/skyrocket-qy/prob-filters/xor-filter/code/xorfilter"
)

func main() {
	// XOR filters are static and built from a set of items.
	// This example simulates a pre-built filter.
	// In a real scenario, you'd construct it from your dataset.

	// Simulate a set of items that would be used to build the filter
	items := []string{"apple", "banana", "cherry", "date", "elderberry"}
	_ = items // Suppress unused variable warning

	// Create a conceptual XOR filter (not actually building it here)
	filter := xorfilter.New(len(items))

	fmt.Println("\nChecking items:")
	checkItems := []string{"apple", "grape", "cherry", "kiwi", "date"}
	for _, item := range checkItems {
		if filter.Contains([]byte(item)) {
			fmt.Printf("'%s' might be in the set (conceptual check).\n", item)
		} else {
			fmt.Printf("'%s' is definitely NOT in the set (conceptual check).\n", item)
		}
	}

	fmt.Println("\nNote: XOR Filter is static. Add/Delete operations are not supported after construction.")
}
