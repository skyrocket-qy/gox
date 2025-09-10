package main

import (
	"fmt"
	"time"
)

// This method not ensure the finish
func main() {
	go say("world")
	go say("hello")

	time.Sleep(5 * time.Second)
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
