package main

import (
	"fmt"
	"sync"
	"time"
)

type Revenue struct {
	Value uint
	sync.RWMutex
}

func (r *Revenue) Add(value uint) {
	r.Lock()
	defer r.Unlock()

	r.Value += value
	fmt.Printf("Add value: %d\n", value)
}

func (r *Revenue) Read() {
	r.RLock()
	defer r.RUnlock()

	fmt.Printf("Read value: %d\n", r.Value)
}

func main() {
	rv := Revenue{}
	fmt.Printf("Revenue value: %d\n", rv.Value)

	for _, v := range []uint{3, 5, 7, 8} {
		go rv.Add(v)
	}

	for range 4 {
		go rv.Read()
	}

	// This cannot ensure all goroutines will finish.
	time.Sleep(1 * time.Second)

	fmt.Printf("Revenue value: %d\n", rv.Value)
}
