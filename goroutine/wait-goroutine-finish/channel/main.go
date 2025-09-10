package main

import (
	"log"
	"time"
)

func main() {
	ch := make(chan string)

	go say("world", ch)
	go say("hello", ch)

	<-ch
	<-ch
}

func say(s string, ch chan string) {
	for range 5 {
		time.Sleep(100 * time.Millisecond)
		log.Println(s)
	}

	ch <- "finish"
}
