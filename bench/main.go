package main

import (
	"fmt"
	"log"

	"github.com/skyrocket-qy/gox/bench/pkg"
)

func main() {
	err := pkg.Bench(benchNumbersWithChannel, 1, "bench.png")
	if err != nil {
		log.Fatalf("Bench failed: %v", err)
	}

	err = pkg.Bench(benchSliceAppend, 1, "bt.png")
	if err != nil {
		log.Fatalf("Bench failed: %v", err)
	}

	err = pkg.Bench(benchCallbackLoop, 1, "cb.png")
	if err != nil {
		log.Fatalf("Bench failed: %v", err)
	}
}

const t = 100000000

func numbers() <-chan int {
	ch := make(chan int, 1024)

	go func() {
		defer close(ch) // close when done

		for i := range t {
			ch <- i // "yield return i"
		}
	}()

	return ch
}

func benchNumbersWithChannel() {
	sum := 0
	for v := range numbers() {
		sum += v
	}

	fmt.Println(sum)
}

func tradit() []int {
	res := []int{}
	for i := range t {
		res = append(res, i)
	}

	return res
}

func benchSliceAppend() {
	sum := 0
	for v := range tradit() {
		sum += v
	}

	fmt.Println(sum)
}

func CallbackLoop(f func(v int)) []int {
	res := []int{}

	for i := range t {
		f(i)
	}

	return res
}

func benchCallbackLoop() {
	sum := 0

	CallbackLoop(func(v int) {
		sum += v
	})
	fmt.Println(sum)
}
