package main

import (
	"fmt"
	"math/rand"
	"time"
	"github.com/skyrocket-qy/prob-filters/t-digest/code/tdigest"
)

func main() {
	digest := tdigest.New(100) // Compression factor of 100

	// Add some random data points
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Adding random data points to t-digest:")
	for i := 0; i < 1000; i++ {
		val := rand.NormFloat64()*10 + 50 // Normal distribution around 50
		digest.Add(val)
		if i % 100 == 0 {
			fmt.Printf("Added %.2f (total %d points)\n", val, i+1)
		}
	}

	fmt.Println("\nEstimating percentiles:")
	fmt.Printf("  5th percentile: %.2f\n", digest.Quantile(0.05))
	fmt.Printf(" 25th percentile: %.2f\n", digest.Quantile(0.25))
	fmt.Printf(" 50th percentile: %.2f\n", digest.Quantile(0.50))
	fmt.Printf(" 75th percentile: %.2f\n", digest.Quantile(0.75))
	fmt.Printf(" 95th percentile: %.2f\n", digest.Quantile(0.95))
	fmt.Printf(" 99th percentile: %.2f\n", digest.Quantile(0.99))
}
