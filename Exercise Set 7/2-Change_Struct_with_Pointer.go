package main

import (
	"fmt"
)

type person struct {
	name string
}

func changeMe(p *person) { //takes the adress of person
	(*p).name = "new name" //dereferenced the struct with *p
}

func main() {
	p1 := person{
		name: "old name",
	}

	fmt.Println(p1.name)

	changeMe(&p1)

	fmt.Println(p1.name)

}
