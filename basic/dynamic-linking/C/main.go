package main

// #cgo LDFLAGS: -L./lib -l:lib.so
// #include "lib/lib.h"
import "C"
import "fmt"

func main() {
	result := C.Add(2, 3)
	fmt.Println("Result:", result)
}
