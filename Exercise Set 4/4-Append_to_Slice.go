package main

import (
	"fmt"
)

func main() {

	sli := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51} // slice composite literal
	sli = append(sli, 52)                                //append single value
	fmt.Println(sli)
	sli = append(sli, 53, 54, 55) //append multiple values to a slice
	fmt.Println(sli)
	x := []int{56, 57, 58, 59, 60}
	sli = append(sli, x...) //append another slice to a slice
	fmt.Println(sli)
}
