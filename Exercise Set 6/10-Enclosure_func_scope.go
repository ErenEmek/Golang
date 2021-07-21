package main

import (
	"fmt"
)

var z int = 10 //global variable

func test(y int) int { //scope of y is in curlies
	return y * 2
}

func test2(y int) int{ //scope of y is in curlies
	return y*3
}

func main() {
	fmt.Println(z)

	a := 2
	
	test(a)
	test2(a)


}
