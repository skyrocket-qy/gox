package main

import "fmt"

func main() {
	ch := make(chan bool)

	n := 6
	for range n {
		go func() {
			ch <- true
		}()
	}

	for n > 0 {
		select {
		case <-ch:
			fmt.Println("select case 1")
		case <-ch:
			fmt.Println("select case 2")
		}

		n--
	}
}
