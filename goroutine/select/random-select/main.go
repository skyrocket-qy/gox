package main

import "log"

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
			log.Println("selected case")
		}

		n--
	}
}
