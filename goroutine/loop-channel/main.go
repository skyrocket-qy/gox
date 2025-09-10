package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func() {
		for i := range 10 {
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
