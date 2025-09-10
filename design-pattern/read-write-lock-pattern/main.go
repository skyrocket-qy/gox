package main

import (
	"log"
	"sync"
	"time"
)

type Revenue struct {
	sync.RWMutex

	Value uint
}

func (r *Revenue) Add(value uint) {
	r.Lock()
	defer r.Unlock()

	r.Value += value
	log.Printf("Add value: %d\n", value)
}

func (r *Revenue) Read() {
	r.RLock()
	defer r.RUnlock()

	log.Printf("Read value: %d\n", r.Value)
}

func main() {
	rv := Revenue{}
	log.Printf("Revenue value: %d\n", rv.Value)

	for _, v := range []uint{3, 5, 7, 8} {
		go rv.Add(v)
	}

	for range 4 {
		go rv.Read()
	}

	// This cannot ensure all goroutines will finish.
	time.Sleep(1 * time.Second)

	log.Printf("Revenue value: %d\n", rv.Value)
}
