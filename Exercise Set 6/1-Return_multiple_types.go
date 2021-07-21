package main

import (
	"fmt"
)

func foo() int {
	return 42
}

func bar() (int, string) {
	return 36, "hello"
}

func main() {
	
	x := foo()
	n, l := bar()
	fmt.Println(x, n, l)
}
