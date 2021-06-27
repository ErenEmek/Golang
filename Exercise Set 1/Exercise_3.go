package main

import "fmt"
//Declaring package level variables

var x int
var y string
var z bool

func main() {
	//Assigning values to pre-declared variables.

	x = 42
	y = "James, Bond"
	z = true

	s:= fmt.Sprintf("%d, %s, %t", x, y, z) //Short declare a string and assign the value created with Sprintf function. "%v, %v, %v" could also be used.


	fmt.Print(s)

}
