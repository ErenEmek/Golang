#Notes

##Function Syntax in Go

*func (r receiver) identifier (parameters) (return(s)) { code }*

Difference btw. parameters and arguments: parameters when you define func, arguments passed to the func.
**Everything in Go is Pass By Value**

foo() // arguments

func foo () { //parameters
	...
}

##Variadic Parameter

 ...

func foo (x ...int) int {	//type is now a slice of int. ([]int)
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	sum := 0
	for _, v in range x {
		sum += x
	}
return sum
}

##Unfurling a Slice

xi := []int{1,2,3,4,5,6}
foo(xi)	// will not run. foo is waiting for number of ints however gets a slice of int.
foo(xi...) // this will work. ... changes slice of int to number of ints.

##Defer

Defers the execution of a function to the exit of surrounding function.

func main() {
	defer foo()
	bar()
}			//foo will run after bar where main function is exiting.

func foo (){
	fmt.Println("Foo")
}
func bar (){
	fmt.Println("Bar)
}

##Methods
Methods are functions that attached to the structs.

type person struct {
	first string
	last string
}
type agent struct {
	person
	ltk bool
}

func (s agent) speak() { //s agent is receiver, which has access to this function
	fmt.Println("I am", s.first, s.last)
}			//receiver attaches the function to the value of that type

func main(){
	sa1 := agent{
		person: person{
				"James",
				"Bond",
			},
		ltk: true,
	}
	sa1.speak()
}

##Interfaces & Polymorphism

Interfaces allowing us to define behaviour. Allows us to do polymorphism.

type person struct {
	first string
	last string
}
type agent struct {
	person
	ltk bool
}

type human interface { //keyword: "type", identifier: "human", type: "interface"
	speak()
}

func bar(h human) {
	fmt.Prinln("I was passed into bar", h)
}

/* a value can be more than one type. sa1 is a type of agent, and able to speak(),
therefore it is also type human */


func (s agent) speak() { //any type that speak attached to it is human
	fmt.Println("I am", s.first, s.last, "-agent speak")
}			

func (p person) speak() { //any type that speak attached to it is human
	fmt.Println("I am", p.first, p.last, "-person speak")
}

func main(){
	sa1 := agent{
		person: person{
				"James",
				"Bond",
			},
		ltk: true,
	}
	sa1.speak()
	
	p1 := {
		first : "dr.",
		last : "yes",
	}
	bar(sa1) //func bar takes type agent
	bar(p1)	 //func bar takes type person
//function bar takes more than one type. this is called polymorphism.
//Interfaces allows a value to be of more than one type.
//First, we bond type human to method speak. And define a func with receiver human.
//We can call the func with any type that got this common "method".
//If something receives empty interface i.e. (a... interface{}) it can take any value
//of any type.
}

##Anonymous func

Is a function without identifier.

func(x int){	//define parameters as shown.
	fmt.Print("anonymous func run", x)
}(1280)		//pass value as shown.

func(){		//without parameters
	fmt.Print("anonymous func run", x)
}()

##A func expression

f := func(x int){	//f is defined this way.
	fmt.Println("func expression", x)
}
f(42)		//execution is standard.

##Return a function from a function

func bar() func() int { //bar func returns a function that returns an integer.
	return func() int { //anonymous function returns int
		return 42
	} //self execute () should not be executed.
}

x := bar()
i := x() //x executed and returns 42
fmt.Println(i) //result is 42

##Callback

Passing a function as an argument.

func sum(xi ...int) int{
	t := 0
	for _, v := range xi{
		t += v
	}
}

func even(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi{
		if v % 2 == 0 { 
			yi = append(yi, v)
		}
	}
return f(yi)
}

ii := []int{1,2,3,4,5}
s:= sum(ii...)

x := even(sum, ii...) //sum function is passed to even, variadic ii is also passed.

##Closure

Enclose scope of variable to an area.

When used above func main() => Global scope

When used in func main() => narrowed to func main

When used in code block { } => narrowed to that code block

func incrementor() func() int {
	var x int
	return func() int {
		x++
	}
}

a := incrementor()
b := incrementor()

fmt.Println(a()) //prints 1
fmt.Println(a()) //prints 2
fmt.Println(b()) //prints 1
fmt.Println(b()) //prints 2

scopes are different. values are different.

##Recursion

Anything can be done with recursion could be done with loops.

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

























