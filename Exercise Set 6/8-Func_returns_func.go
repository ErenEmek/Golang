package main

import (
	"fmt"
)

func returner() func() { //a func that returns func is created.
	return func() {
		fmt.Println("I'm a returned func")
	}
}

func main() {
	f := returner() //func returned and assigned to f identifier
	f()             //returned func is called
}
