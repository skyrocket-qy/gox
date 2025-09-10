package main

import (
	"fmt"
	"sync"
	"time"
)

type safeNumber struct {
	val int
	mux sync.Mutex
}

func main() {
	total := safeNumber{val: 0}
	for i := 0; i < 1000; i++ {
		go func() {
			total.mux.Lock()
			total.val++
			total.mux.Unlock()
		}()
	}
	time.Sleep(time.Second)

	fmt.Println(total.val)
}
