package main

import (
	"log"
	"time"
)

func main() {
	// same channel
	ch := make(chan string)
	go A(ch)
	go B(ch)
	go C(ch)

	log.Println(<-ch)

	// separate channel
	aCh := make(chan string)
	bCh := make(chan string)
	cCh := make(chan string)

	go A(aCh)
	go B(bCh)
	go C(cCh)

	select {
	case a := <-aCh:
		log.Println(a)
	case b := <-bCh:
		log.Println(b)
	case c := <-cCh:
		log.Println(c)
	default:
		// time.Sleep(time.Second)
	}

	// another
	dCh := D()
	eCh := E()
	fCh := F()

	select {
	case d := <-dCh:
		log.Println(d)
	case e := <-eCh:
		log.Println(e)
	case f := <-fCh:
		log.Println(f)
	default:
		// time.Sleep(time.Second)
	}
}

func A(ch chan string) {
	time.Sleep(2 * time.Second)

	ch <- "A error"
}

func B(ch chan string) {
	time.Sleep(2 * time.Second)

	ch <- "B error"
}

func C(ch chan string) {
	time.Sleep(2 * time.Second)

	ch <- "C error"
}

func D() chan string {
	ch := make(chan string)
	go func(ch chan string) {
		time.Sleep(2 * time.Second)

		ch <- "D error"
	}(ch)

	return ch
}

func E() chan string {
	ch := make(chan string)
	go func(ch chan string) {
		time.Sleep(2 * time.Second)

		ch <- "E error"
	}(ch)

	return ch
}

func F() chan string {
	ch := make(chan string)
	go func(ch chan string) {
		time.Sleep(2 * time.Second)

		ch <- "F error"
	}(ch)

	return ch
}
