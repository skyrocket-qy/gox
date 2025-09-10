package main

import (
	"fmt"
	"time"
)

// This method not ensure the finish.
func main() {
	go say("world")
	go say("hello")

	time.Sleep(5 * time.Second)
}

func say(s string) {
	for range 5 {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
