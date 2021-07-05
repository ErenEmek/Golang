package main

import (
	"fmt"
)

func main() {

	sli := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51} // slice composite literal
	fmt.Println(sli[:5])                                 //from start to 4th element.
	fmt.Println(sli[5:])                                 //from index 5 element to the end.
	fmt.Println(sli[2:7])                                //from i:2 element (44) to i:6 element (48).
	fmt.Println(sli[1:6])                                //from i:1 element to i:5 element (47).
}
