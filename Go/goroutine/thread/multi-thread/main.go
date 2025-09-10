package main

import (
	"fmt"
	"time"
)

func main() {
	go say("world")

	say("hello")
}

func say(s string) {
	for range 5 {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
