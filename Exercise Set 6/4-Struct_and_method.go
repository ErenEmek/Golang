package main

import (
	"fmt"
)

type person struct { // struct with identifier person is defined.

	first string
	last  string
	age   int
}

func (p person) speak() { //method to type person is attached.
	fmt.Println("My words are: ", p.first, p.last, p.age)
}

func main() {
	p1 := person{ //a value of type person created.
		first: "James",
		last:  "Bond",
		age:   42,
	}
	p1.speak() //method is called from the value of type person
}
