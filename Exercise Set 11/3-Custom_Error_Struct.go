package main

import (
	"fmt"
)

type customErr struct {
	i string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("information about error: %v", ce.i)
}



func main() {
	fmt.Println("Hello, playground")
	c1 := customErr{
		i: "test",
	}
	
	foo(c1)
}

func foo(e error){
	fmt.Println("foo ran -", e)
}