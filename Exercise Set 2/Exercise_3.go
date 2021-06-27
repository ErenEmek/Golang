package main

import "fmt"

// Create typed and untyped constants, assign values to those constants
const(
	a = 42		// "a constant is just a simple, unchanging value" Rob Pike
	b int = 42
)

func main() {

	fmt.Println(a, b)

}
