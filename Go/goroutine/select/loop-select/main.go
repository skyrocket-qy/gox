package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	defer close(ch)

	go func() {
		time.Sleep(time.Second * 2)
		ch <- 1
	}()

LOOP:
	for {
		select {
		case val := <-ch:
			fmt.Println(val)
			break LOOP
		default:
			fmt.Println("watching...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
