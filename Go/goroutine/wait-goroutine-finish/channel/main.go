package main

import (
	"fmt"
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
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
	ch <- "finish"
}
