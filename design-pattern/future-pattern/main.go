package main

import (
	"log"
	"sync"
)

type Revenue struct {
	sync.Mutex

	Value uint
}

func (r *Revenue) Add(value uint, wg *sync.WaitGroup) {
	r.Lock()
	defer r.Unlock()

	r.Value += value
	log.Printf("Add value: %d\n", value)
	wg.Done()
}

func main() {
	rv := Revenue{}
	log.Printf("Revenue value: %d\n", rv.Value)

	wg := sync.WaitGroup{}
	wg.Add(4)

	for _, v := range []uint{3, 5, 7, 8} {
		go rv.Add(v, &wg)
	}

	// This cannot ensure all goroutines will finish.
	wg.Wait()

	log.Printf("Revenue value: %d\n", rv.Value)
}
