package main

import (
	"fmt"
)

func passed(x int) int { //passed is a func takes an int as argument returns an int
	return x * 10
}

func passed2(x int) int {
	return x * 2
}

func funk(p func(a int) int, x int) int { //funk takes a func and an integer as argument returns an in
	return p(x)			//funk takes value x, p(x) it with passed func, returns resulting value
}

func main() {
	x := 5
	fmt.Println(funk(passed, x))
	fmt.Println(funk(passed2, x)) 
}

