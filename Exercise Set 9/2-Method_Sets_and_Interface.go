package main

import (
	"fmt"
)

type person struct { //defined a struct with variable name person.
	name string
}

func (p *person) speak() { //attached a method speak to type person using a pointer receiver.
	p.name = "rename" //because of passed pointer, we can change the data.
	fmt.Println(p.name, "speaks")
}

type human interface { //human interface with speak method is created.
	speak()
}

func saySomething(h human) { //func saysomething takes human as parameter, calls the speak method.
	h.speak()
}

func main() {
	p1 := person{
		name: "Person 1",
	}
	saySomething(&p1) //If we pass pointer to the function, it will work. If we pass person to the function, it won't work.
	
	p1.speak() //also works
}
