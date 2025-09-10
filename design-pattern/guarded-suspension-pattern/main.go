package main

import (
	"log"
	"strconv"
	"time"
)

type MessageQueue struct {
	Queue chan string
}

func (m *MessageQueue) Put(message string) {
	m.Queue <- message
}

func (m *MessageQueue) Pop() string {
	return <-m.Queue
}

func main() {
	mq := MessageQueue{
		Queue: make(chan string, 2),
	}

	go func() {
		for i := range 5 {
			mq.Put(strconv.Itoa(i))
		}
	}()

	go func() {
		for {
			log.Println(mq.Pop())
		}
	}()

	time.Sleep(1 * time.Second)
}
