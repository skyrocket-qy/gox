package main

import (
	"log"
	"time"
)

func main() {
	total := 0

	for range 1000 {
		go func() {
			total++
		}()
	}

	time.Sleep(time.Second)
	log.Println(total)
}
