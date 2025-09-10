package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	"time"

	cts "pqueue/custom_noupdate"
	lib "pqueue/library_noupdate"
)

var (
	SEED  int64 = 5487319684720
	LEN   int   = 500000000
	MAX   int   = 256
	SCOPE int   = 100000000
)

func main() {
	rand.Seed(SEED)
	eles := make([]int, LEN)
	for i := 0; i < LEN; i++ {
		eles[i] = rand.Intn(MAX)
	}

	ctsq := cts.NewPriorityQueue()
	for i := 0; i < LEN-SCOPE; i++ {
		ctsq.Push(cts.Node{
			Priority: uint(eles[i]),
		})
	}
	now := time.Now().Unix()
	for i := LEN - SCOPE; i < LEN; i++ {
		ctsq.Push(cts.Node{
			Priority: uint(eles[i]),
		})
	}
	fmt.Printf("push: %d\n", time.Now().Unix()-now)
	ctsres := make([]uint, SCOPE)
	now = time.Now().Unix()
	for i := 0; i < SCOPE; i++ {
		ctsres[i] = ctsq.Pop().Priority
	}
	fmt.Printf("pull: %d\n", time.Now().Unix()-now)

	// libq
	libq := lib.NewPriorityQueue()
	for i := 0; i < LEN-SCOPE; i++ {
		libq.Push(lib.Node{
			Priority: uint(eles[i]),
		})
	}
	heap.Init(&libq)
	now = time.Now().Unix()
	for i := LEN - SCOPE; i < LEN; i++ {
		heap.Push(&libq, lib.Node{
			Priority: uint(eles[i]),
		})
	}
	fmt.Printf("push: %d\n", time.Now().Unix()-now)
	libres := make([]uint, SCOPE)
	now = time.Now().Unix()
	for i := 0; i < SCOPE; i++ {
		libres[i] = heap.Pop(&libq).(lib.Node).Priority
	}
	fmt.Printf("pull: %d\n", time.Now().Unix()-now)

	for i := 0; i < SCOPE; i++ {
		if ctsres[i] != libres[i] {
			fmt.Println(i)
			fmt.Println(ctsres[i : i+10])
			fmt.Println(libres[i : i+10])
			break
		}
	}
}
