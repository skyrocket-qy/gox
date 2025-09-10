package main

import (
	"fmt"
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
	fmt.Println(total)
}
