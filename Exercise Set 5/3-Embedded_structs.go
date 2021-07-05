package main

import (
	"fmt"
)

type vehicle struct { //create a type vehicle underlying type is struct.
	doors int
	color string
}

type truck struct {
	vehicle	// type vehicle is embedded to type truck.
	fourWheel bool
}

type sedan struct {
	vehicle // type vehicle is embedded to type sedan
	luxury bool
}

func main() {
	v1 := truck{
		vehicle: vehicle{ //assign values to fields using composite literal
			doors: 2,
			color: "red",
		},
		fourWheel: true,
	}

	v2 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "blue",
		},
		luxury: false,
	}

	fmt.Println(v1.doors, v1.color, v2.luxury, v2.doors) 
}
