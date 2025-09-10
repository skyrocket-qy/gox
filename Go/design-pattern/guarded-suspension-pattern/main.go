package main

import (
	"fmt"
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
		for i := 0; i < 5; i++ {
			mq.Put(strconv.Itoa(i))
		}
	}()

	go func() {
		for {
			fmt.Println(mq.Pop())
		}
	}()

	time.Sleep(1 * time.Second)
}
