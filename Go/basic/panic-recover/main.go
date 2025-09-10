package main

import "fmt"

func main() {
	parent()
	fmt.Println("end")
}

func parent() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("recovered in parent func", r)
		}
	}()

	fmt.Println("Calling child func")
	child(0)
	fmt.Println("this cannot be execute")
}

func child(i int) {
	if i > 3 {
		panicNum := fmt.Sprintf("%v", i)
		fmt.Println("Panicking!", panicNum)
		panic(panicNum)
	}

	defer fmt.Println("Defer in child", i)
	fmt.Println("Printing in child", i)
	child(i + 1)
}
