package main

import (
	"fmt"
	"github.com/skyrocket-qy/prob-filters/skip-list/code/skiplist"
)

func main() {
	sl := skiplist.New()

	itemsToAdd := []int{10, 5, 20, 15, 7, 25, 12}
	fmt.Println("Inserting items into Skip List:")
	for _, item := range itemsToAdd {
		sl.Insert(item)
		fmt.Printf("Inserted: %d\n", item)
	}

	fmt.Println("\nSearching for items:")
	checkItems := []int{7, 1, 20, 18, 25}
	for _, item := range checkItems {
		if sl.Search(item) {
			fmt.Printf("Found: %d\n", item)
		} else {
			fmt.Printf("Not found: %d\n", item)
		}
	}

	fmt.Println("\nDeleting items:")
	deleteItems := []int{5, 15, 1}
	for _, item := range deleteItems {
		if sl.Delete(item) {
			fmt.Printf("Deleted: %d\n", item)
		} else {
			fmt.Printf("Failed to delete: %d (not found)\n", item)
		}
	}

	fmt.Println("\nSearching for items after deletion:")
	for _, item := range checkItems {
		if sl.Search(item) {
			fmt.Printf("Found: %d\n", item)
		} else {
			fmt.Printf("Not found: %d\n", item)
		}
	}
}
