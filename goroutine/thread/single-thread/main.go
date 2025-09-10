package main

import (
	"log"
	"time"
)

func say(s string) {
	for range 5 {
		time.Sleep(100 * time.Millisecond)
		log.Println(s)
	}
}

func main() {
	say("hello")
	say("world")
}
