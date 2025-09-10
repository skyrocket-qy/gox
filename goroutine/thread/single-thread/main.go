package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for range 5 {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	say("hello")
	say("world")
}
