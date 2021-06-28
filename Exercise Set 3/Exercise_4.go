package main

import "fmt"

func main () {
	// for { } loop
	x := 1985
	for {
		fmt.Println(x)
		if x == 2021 {
			break
		}
		x++
	}

}