package main

import (
	"log"
	"time"

	"github.com/skyrocket-qy/gox/bench/pkg"
)

func main() {
	err := pkg.Bench(func() {
		time.Sleep(10 * time.Millisecond)
	}, 100, "bench.png")
	if err != nil {
		log.Fatalf("Bench failed: %v", err)
	}
}
