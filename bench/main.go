package main

import (
	"fmt"
	"log"

	"github.com/skyrocket-qy/gox/bench/pkg"
)

func main() {
	err := pkg.Bench(bNumbers, 1, "bench.png")
	if err != nil {
		log.Fatalf("Bench failed: %v", err)
	}

	err = pkg.Bench(bt, 100, "bt.png")
	if err != nil {
		log.Fatalf("Bench failed: %v", err)
	}
}

const t = 100000

func numbers() <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch) // close when done
		for i := 0; i <= t; i++ {
			ch <- i // "yield return i"
		}
	}()
	return ch
}

func bNumbers() {
	sum := 0
	for v := range numbers() {
		sum += v
	}
	fmt.Println(sum)
}

func tradit() []int {
	res := []int{}
	for i := 0; i < t; i++ {
		res = append(res, i)
	}
	return res
}

func bt() {
	sum := 0
	for v := range tradit() {
		sum += v
	}
	fmt.Println(sum)
}
