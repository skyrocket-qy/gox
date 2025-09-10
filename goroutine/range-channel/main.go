package main

import "log"

func main() {
	c := make(chan int, 10)

	go func() {
		for i := range 10 {
			c <- i
		}

		close(c)
	}()

	for i := range c {
		log.Println(i)
	}
}
