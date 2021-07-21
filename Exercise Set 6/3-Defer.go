package main

import (
	"fmt"
)

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}

func main() {

	defer foo() //foo is executed when the main func is about to exit.
	bar()

}