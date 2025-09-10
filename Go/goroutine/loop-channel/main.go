package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	for {
		if i, ok := <-ch; !ok {
			break
		} else {
			fmt.Println(i)
		}
	}
}
