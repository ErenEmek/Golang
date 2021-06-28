package main

import "fmt"

func main () {
//Use switch without expression
	switch {
	case false:
		fmt.Print("This is false")
	case true:
		fmt.Println("True")
		fmt.Println("even in a line below")
	case true:
		fmt.Println("Should not be seen")
	}

}