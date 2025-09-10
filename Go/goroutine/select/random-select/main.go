package main

import "fmt"

func main() {
	ch := make(chan bool)
	n := 6
	for i := 0; i < n; i++ {
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
