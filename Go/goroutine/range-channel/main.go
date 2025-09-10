package main

import "fmt"

func main() {
	c := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()
	for i := range c {
		fmt.Println(i)
	}
}
