package main

import (
	"fmt"
)

func main() {

	arr := [5]int{1, 2, 3, 4, 5} // array composite literal
	for i, v := range arr {      // range over the elements
		fmt.Println(i, v)
	}
	fmt.Printf("%T", arr) //print out the type of the variable

}
