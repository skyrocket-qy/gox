package main

import "C"

// Exported function to add two numbers
//
//export Add
func Add(a, b int) int {
	return a + b
}

func main() {} // Necessary for a shared library
