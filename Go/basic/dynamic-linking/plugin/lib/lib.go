package main

import "fmt"

// Add is a simple function that will be exposed as a plugin.
func Add(a, b int) int {
	return a + b
}

// Greet prints a greeting message.
func Greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}
