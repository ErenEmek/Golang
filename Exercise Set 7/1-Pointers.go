package main

import (
	"fmt"
)

func main() {
	a := 10
	fmt.Println(&a)//prints adress of variable
	fmt.Printf("%T\n", &a) //prints type
	
	var b *int = &a 
	c := &a
	*b = 42 //dereferencing b from adress to value.
	fmt.Println(b, c, *c)
}
