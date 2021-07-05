package main

import (
	"fmt"
)

func main() {
	x := []string{"james", "bond", "shaken, not stirred"}
	y := []string{"miss", "moneypenny", "hellooo james"}

	z := [][]string{x, y} //slice of a slice of string. multi dimensional slice

	for i, v := range z { //ranging x, y
		for j, y := range v { //ranging elements 0, 1, 2
			fmt.Println(i, j, y)
		}
	}

}
