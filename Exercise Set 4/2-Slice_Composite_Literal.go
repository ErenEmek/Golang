package main

import (
	"fmt"
)

func main() {

	sli := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0} // slice composite literal
	for i, v := range sli {                    // range over the elements
		fmt.Println(i, v)
	}
	fmt.Printf("%T", sli) //print out the type of the variable

}
