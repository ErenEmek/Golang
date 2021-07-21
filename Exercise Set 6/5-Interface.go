package main

import (
	"fmt"
)

type square struct { //square type is defined.
	leng int
}

type circle struct { //circle type is defined.
	rad int
}

func (s square) calc_area() int { //calculate method for square type
	return (s.leng * s.leng)
}

func (c circle) calc_area() int { //calculate method for circle type
	return (c.rad * c.rad * 3)
}

type shape interface { //shape is defined as an interface to anything that has calc_area method
	calc_area() int //method returns an int
}

func info(s shape) { //func info takes the shape and prints the area
	fmt.Println("Area is: ", s.calc_area()) //s.calc_area uses correct func to underlying type 
}

func main() {
	sq1 := square{ 
		leng: 5,
	}
	cr1 := circle{
		dia: 10,
	}

	info(sq1)
	info(cr1)

}
