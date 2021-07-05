package main

import (
	"fmt"
)

func main() {

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51} // slice composite literal
	x = append(x[0:3], x[6:]...) //deleted i:3, 4, 5 from slice
	fmt.Println(x)
}
