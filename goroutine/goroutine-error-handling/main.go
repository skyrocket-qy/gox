package main

import (
	"fmt"
	"time"
)

func main() {
	res := checkVariable()

	for {
		if res.Error() != "" {
			fmt.Println(res.Error())

			break
		}

		time.Sleep(1 * time.Second)
	}

	resCh := checkChannel()
	fmt.Println(<-resCh)
}

type Result struct {
	ErrorMessage string
}

func (r Result) Error() string {
	return r.ErrorMessage
}

// You can choose any return type, like struct, string, bool ...etc.
func checkVariable() error {
	fmt.Println("Variable method...")

	err := Result{}
	go func(err *Result) {
		time.Sleep(2 * time.Second)

		err.ErrorMessage = "check variable error"
	}(&err)

	return &err
}

func checkChannel() chan string {
	fmt.Println("channel method...")

	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)

		ch <- "check channel error"
	}()

	return ch
}
