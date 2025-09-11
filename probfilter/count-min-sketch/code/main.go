package main

import (
	"fmt"

	"github.com/skyrocket-qy/prob-filters/count-min-sketch/code/countminsketch"
)

func main() {
	cms := countminsketch.New(0.001, 0.01) // epsilon=0.001, delta=0.01

	items := []string{
		"apple", "banana", "apple", "cherry", "banana",
		"apple", "fig", "banana", "grape", "apple",
		"cherry", "banana", "apple", "kiwi", "banana",
	}

	fmt.Println("Adding items to Count-Min Sketch:")
	for _, item := range items {
		cms.Add([]byte(item))
		fmt.Printf("Added: %s\n", item)
	}

	fmt.Println("\nEstimated frequencies:")
	checkItems := []string{"apple", "banana", "cherry", "grape", "orange"}
	for _, item := range checkItems {
		fmt.Printf("Estimated frequency of '%s': %d\n", item, cms.Estimate([]byte(item)))
	}

	// True frequencies for comparison
	trueFrequencies := make(map[string]int)
	for _, item := range items {
		trueFrequencies[item]++
	}
	fmt.Println("\nTrue frequencies:")
	for item, count := range trueFrequencies {
		fmt.Printf("True frequency of '%s': %d\n", item, count)
	}
}
