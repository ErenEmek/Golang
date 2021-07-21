package main

import (
	"fmt"
)

func foo(x ...int) int { //func takes in a variadic parameter of type int
	s := 0
	for _, v := range x {
		s += v
	}
	return s
}

func bar(y []int) int { //func takes in slice of ints
	s := 0
	for _, v := range y {
		s += v
	}

	return s
}

func main() {

	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	z := foo(x...)
	fmt.Println(z)

	y := bar(x)
	fmt.Println(y)

}
