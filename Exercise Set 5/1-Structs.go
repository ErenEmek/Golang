package main

import (
	"fmt"
)

type pers struct { // type pers created, underlying type is struct
	first      string
	last       string
	fav_flavor []string
}

func main() {

	p1 := pers{ //a value of type pers created with composite literal.
		"james",
		"bond",
		[]string{"vanilla", "chocolate"},
	}
	p2 := pers{
		"miss",
		"moneypenny",
		[]string{"hazelnut", "melon"},
	}
	fmt.Println(p1.first, p1.last)
	for i, v := range p1.fav_flavor {
		fmt.Println(i, v)
	}
	fmt.Println(p2.first, p2.last)
	for i, v := range p2.fav_flavor {
		fmt.Println(i, v)
	}
}
