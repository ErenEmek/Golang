package main

import (
	"fmt"
)

func main() {
	v1 := struct {	//anonymous struct created.
		vehicle string
		doors   int
		color   string
		details map[string]int //map element added.
	}{vehicle: "sedan", //values assigned using composite literal.
		doors: 4,
		color: "black",
		details: map[string]int{
			"lenght": 4700,
			"weight": 1500,
		},
	}
	fmt.Println(v1)
}
