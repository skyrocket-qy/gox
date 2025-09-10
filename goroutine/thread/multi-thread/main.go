package main

import (
	"log"
	"time"
)

func main() {
	go say("world")

	say("hello")
}

func say(s string) {
	for range 5 {
		time.Sleep(100 * time.Millisecond)
		log.Println(s)
	}
}
