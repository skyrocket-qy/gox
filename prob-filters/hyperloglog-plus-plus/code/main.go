package main

import (
	"fmt"
	"github.com/skyrocket-qy/prob-filters/hyperloglog-plus-plus/code/hyperloglogplusplus"
)

func main() {
	hllpp := hyperloglogplusplus.New(14) // 14-bit precision

	items1 := []string{
		"apple", "banana", "cherry", "date", "elderberry",
		"fig", "grape", "honeydew", "ice cream", "jicama",
	}

	fmt.Println("Adding items to HyperLogLog++ (HLLPP1):")
	for _, item := range items1 {
		hllpp.Add([]byte(item))
		fmt.Printf("Added: %s\n", item)
	}

	estimate1 := hllpp.Estimate()
	fmt.Printf("\nHLLPP1 Estimated unique items: %.2f\n", estimate1)

	// Create another HLL++ for merging
	hllpp2 := hyperloglogplusplus.New(14)
	items2 := []string{
		"kiwi", "lemon", "mango", "nectarine", "orange",
		"apple", "banana", "pear", "quince", "raspberry",
	}

	fmt.Println("\nAdding items to HyperLogLog++ (HLLPP2):")
	for _, item := range items2 {
		hllpp2.Add([]byte(item))
		fmt.Printf("Added: %s\n", item)
	}

	estimate2 := hllpp2.Estimate()
	fmt.Printf("\nHLLPP2 Estimated unique items: %.2f\n", estimate2)

	// Merge HLLPP2 into HLLPP1
	fmt.Println("\nMerging HLLPP2 into HLLPP1...")
	err := hllpp.Merge(hllpp2)
	if err != nil {
		fmt.Printf("Merge error: %v\n", err)
		return
	}

	estimateMerged := hllpp.Estimate()
	fmt.Printf("Merged HLLPP Estimated unique items: %.2f\n", estimateMerged)

	// Calculate true unique count for merged set
	trueUniqueMerged := make(map[string]struct{})
	for _, item := range items1 {
		trueUniqueMerged[item] = struct{}{}
	}
	for _, item := range items2 {
		trueUniqueMerged[item] = struct{}{}
	}
	fmt.Printf("True unique items (merged): %d\n", len(trueUniqueMerged))
}
