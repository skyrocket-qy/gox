package main

import (
	"fmt"
	"sync"
	"time"
)

type taskqueue struct {
	Jobs []bool
	sync.Mutex
}

func (q *taskqueue) Pop() bool {
	q.Lock()
	defer q.Unlock()
	if len(q.Jobs) == 0 {
		time.Sleep(1 * time.Hour)
	}
	res := q.Jobs[0]
	q.Jobs = q.Jobs[1:]
	return res
}

func main() {
	workers := 10000
	jobs := 2000000
	g := sync.WaitGroup{}
	g.Add(jobs)

	tq := taskqueue{
		Jobs: make([]bool, jobs),
	}
	for i := 0; i < jobs; i++ {
		tq.Jobs[i] = true
	}

	now := time.Now().Unix()
	for i := 0; i < workers; i++ {
		go worker(&tq, &g)
	}

	g.Wait()
	fmt.Println(time.Now().Unix() - now)

	tq2 := taskqueue2{
		Jobs: make([]bool, jobs),
	}
	for i := 0; i < jobs; i++ {
		tq2.Jobs[i] = true
	}
	qs := make([]chan bool, workers)
	for i := 0; i < workers; i++ {
		qs[i] = make(chan bool, 10)
	}
	g.Add(jobs)
	now = time.Now().Unix()
	go func(t *taskqueue2, qs []chan bool) {
		for {
			for i, q := range qs {
				if len(q) < 10 {
					qs[i] <- t.Pop()
				}
			}
		}
	}(&tq2, qs)

	for i := 0; i < workers; i++ {
		go worker2(qs[i], &g)
	}

	g.Wait()
	fmt.Println(time.Now().Unix() - now)
}

func worker(queue *taskqueue, g *sync.WaitGroup) {
	for {
		queue.Pop()
		time.Sleep(50 * time.Millisecond)
		g.Done()
	}
}

type taskqueue2 struct {
	Jobs []bool
}

func (q *taskqueue2) Pop() bool {
	if len(q.Jobs) == 0 {
		time.Sleep(1 * time.Hour)
	}
	res := q.Jobs[0]
	q.Jobs = q.Jobs[1:]

	return res
}

type WorkerQueue struct {
	Jobs chan bool
}

func worker2(jobs chan bool, g *sync.WaitGroup) {
	for {
		<-jobs
		time.Sleep(50 * time.Millisecond)
		g.Done()
	}
}
