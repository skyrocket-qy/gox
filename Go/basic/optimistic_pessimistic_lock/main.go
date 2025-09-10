package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Locker struct {
	M sync.Mutex
}

func main() {
	fmt.Println("Increase")
	testIncrease(increase)

	fmt.Println("Increase with lock")
	l := &Locker{
		M: sync.Mutex{},
	}
	testIncrease(l.increaseWithLock)

	fmt.Println("Increase with atomic")
	testIncrease(increaseWithAtomic)

	fmt.Println("Increase with CAS")
	testIncrease(increaseWithCAS)
}

func testIncrease(increaseFunc func(wg *sync.WaitGroup, num *int32, old int32)) {
	var num int32 = 0

	count := 1000
	wg := sync.WaitGroup{}
	wg.Add(count)

	curTime := time.Now()
	for index := 0; index < int(count); index++ {
		go increaseFunc(&wg, &num, int32(index))
	}

	wg.Wait()
	fmt.Printf("execute time = %s\n", time.Since(curTime).String())
	fmt.Printf("count is %d\n", num)
}

func increase(wg *sync.WaitGroup, num *int32, old int32) {
	defer wg.Done()
	*num++
}

func (l *Locker) increaseWithLock(wg *sync.WaitGroup, num *int32, old int32) {
	defer wg.Done()
	l.M.Lock()
	defer l.M.Unlock()
	*num++
}

func increaseWithAtomic(wg *sync.WaitGroup, num *int32, old int32) {
	defer wg.Done()
	atomic.AddInt32(num, 1)
}

func increaseWithCAS(wg *sync.WaitGroup, num *int32, old int32) {
	defer wg.Done()
	var retryCount int
	for {
		if atomic.CompareAndSwapInt32(num, old, old+1) {
			break
		}
		retryCount++
		time.Sleep(time.Millisecond) // avoid more self cycle
	}

	// fmt.Printf("interval(%d) retry(%d) times \n", old, retryCount)
}
