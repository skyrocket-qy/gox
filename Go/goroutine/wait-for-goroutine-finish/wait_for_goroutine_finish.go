package wait_for_goroutine_finish

import "sync"

func WaitForGoroutineFinishUseWaitGroup(funcs ...func()) {
	count := len(funcs)
	wg := sync.WaitGroup{}
	wg.Add(count)

	for _, f := range funcs {
		fc := f
		go func() {
			fc()
			wg.Done()
		}()
	}

	wg.Wait()
}

func WaitForGoroutineFinishUseChannel(funcs ...func()) {
	ch := make(chan bool)

	for _, f := range funcs {
		fc := f
		go func() {
			fc()
			ch <- true
		}()
	}

	for i := 0; i < len(funcs); i++ {
		<-ch
	}
}
