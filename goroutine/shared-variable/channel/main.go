package main

import (
	"log"
	"time"
)

func main() {
	total := 0

	ch := make(chan int, 1)
	ch <- total

	for range 1000 {
		go func() {
			ch <- <-ch + 1
		}()
	}

	time.Sleep(time.Second)
	log.Println(<-ch)
}
