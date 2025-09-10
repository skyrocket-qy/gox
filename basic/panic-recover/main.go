package main

import (
	"log"
	"strconv"
)

func main() {
	parent()
	log.Println("end")
}

func parent() {
	defer func() {
		r := recover()
		if r != nil {
			log.Println("recovered in parent func", r)
		}
	}()

	log.Println("Calling child func")
	child(0)
	log.Println("this cannot be execute")
}

func child(i int) {
	if i > 3 {
		panicNum := strconv.Itoa(i)
		log.Println("Panicking!", panicNum)
		panic(panicNum)
	}

	defer log.Println("Defer in child", i)

	log.Println("Printing in child", i)
	child(i + 1)
}
