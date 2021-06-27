package main

import "fmt"

func main() {
// Prints number in binary, decimal, hex

	number := 10	// keep scope narrow as possible
	fmt.Printf("%b, %d, %#X", number, number, number)

}
