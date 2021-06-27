package main

import "fmt"

type myType int	// new type created as myType with underlying type int
var x myType	// new variable with the identifier "x" created by using keyword "var"

func main() {

	fmt.Println(x) 				 // default value is 0.
	fmt.Printf("%T\n", x) // print type with %T parameter
	x = 42
	fmt.Println(x)

}
