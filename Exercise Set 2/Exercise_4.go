package main

import "fmt"

func main() {

	v:= 42
	fmt.Printf("%b, %d, %#X\n", v, v, v)

	y:= v << 1

	fmt.Printf("%b, %d, %#X", y, y, y)

}
