package main

import (
	"fmt"
	"sync"
	"time"
)

type Revenue struct {
	Value uint
	sync.Mutex
}

func (r *Revenue) Add(value uint) {
	r.Lock()
	defer r.Unlock()
	r.Value += value
	fmt.Printf("Add value: %d\n", value)
}

func main() {
	rv := Revenue{}
	fmt.Printf("Revenue value: %d\n", rv.Value)

	for _, v := range []uint{3, 5, 7, 8} {
		go rv.Add(v)
	}

	// This cannot ensure all goroutines will finish.
	time.Sleep(1 * time.Second)

	fmt.Printf("Revenue value: %d\n", rv.Value)
}
