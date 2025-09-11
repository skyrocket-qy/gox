package main

import (
	"fmt"
	"github.com/skyrocket-qy/prob-filters/hyperloglog/code/hyperloglog"
)

func main() {
	hll := hyperloglog.New(14) // 14-bit precision

	items := []string{
		"apple", "banana", "cherry", "date", "elderberry",
		"apple", "fig", "grape", "honeydew", "ice cream",
		"banana", "jicama", "kiwi", "lemon", "mango",
		"cherry", "nectarine", "orange", "pear", "quince",
	}

	fmt.Println("Adding items to HyperLogLog:")
	for _, item := range items {
		hll.Add([]byte(item))
		fmt.Printf("Added: %s\n", item)
	}

	estimate := hll.Estimate()
	fmt.Printf("\nEstimated unique items: %.2f\n", estimate)

	// Calculate true unique count
	trueUnique := make(map[string]struct{})
	for _, item := range items {
		trueUnique[item] = struct{}{}
	}
	fmt.Printf("True unique items: %d\n", len(trueUnique))
}
