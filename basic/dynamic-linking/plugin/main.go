package main

import (
	"fmt"
	"plugin"
)

func main() {
	// Open the plugin
	p, err := plugin.Open("lib/lib.so")
	if err != nil {
		fmt.Println("Error loading plugin:", err)

		return
	}

	// Look up the Add function
	addSymbol, err := p.Lookup("Add")
	if err != nil {
		fmt.Println("Error looking up Add:", err)

		return
	}

	// Assert that Add is a function with the expected signature
	addFunc, ok := addSymbol.(func(int, int) int)
	if !ok {
		fmt.Println("Invalid function signature for Add")

		return
	}

	// Call the Add function
	result := addFunc(2, 3)
	fmt.Println("Result from Add:", result)

	// Look up the Greet function
	greetSymbol, err := p.Lookup("Greet")
	if err != nil {
		fmt.Println("Error looking up Greet:", err)

		return
	}

	// Assert that Greet is a function with the expected signature
	greetFunc, ok := greetSymbol.(func(string))
	if !ok {
		fmt.Println("Invalid function signature for Greet")

		return
	}

	// Call the Greet function
	greetFunc("Alice")
}
