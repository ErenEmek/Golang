package main

import "fmt"

func main() {

	// Prints every rune of the uppercase letters 3 times.

	for j := 65; j < 91; j++ {
		fmt.Println(j)
		for i := 0 ; i < 3 ; i++ {
	
			fmt.Printf("%#U\n", j)
		}
	}
}
