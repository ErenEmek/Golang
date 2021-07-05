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

	m := map[string]pers{	//map to pers' created with last names as keys.
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range m {	//ranged over the map and the slices
		fmt.Println(v.first, v.last)
		for _, v2 := range v.fav_flavor {
			fmt.Println(v2)
		}
	}
}
