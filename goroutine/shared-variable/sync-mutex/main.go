package main

import (
	"log"
	"sync"
	"time"
)

type safeNumber struct {
	val int
	mux sync.Mutex
}

func main() {
	total := safeNumber{val: 0}

	for range 1000 {
		go func() {
			total.mux.Lock()
			total.val++
			total.mux.Unlock()
		}()
	}

	time.Sleep(time.Second)

	log.Println(total.val)
}
