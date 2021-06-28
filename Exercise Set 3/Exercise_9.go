package main

import "fmt"

func main () {
//Switch with a variable expression
	favSport := "none"	

	switch favSport {
	case "boxing":
		fmt.Println("Box")
	case "tennis":
		fmt.Println("Tennis")
	case "none":
		fmt.Println("none")
	}

}