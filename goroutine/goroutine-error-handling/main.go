package main

import (
	"log"
	"time"
)

func main() {
	res := checkVariable()

	for {
		if res.Error() != "" {
			log.Println(res.Error())

			break
		}

		time.Sleep(1 * time.Second)
	}

	resCh := checkChannel()
	log.Println(<-resCh)
}

type ResultError struct {
	ErrorMessage string
}

func (r ResultError) Error() string {
	return r.ErrorMessage
}

// You can choose any return type, like struct, string, bool ...etc.
func checkVariable() error {
	log.Println("Variable method...")

	err := ResultError{}
	go func(err *ResultError) {
		time.Sleep(2 * time.Second)

		err.ErrorMessage = "check variable error"
	}(&err)

	return &err
}

func checkChannel() chan string {
	log.Println("channel method...")

	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)

		ch <- "check channel error"
	}()

	return ch
}
