package main

import "fmt"

func main() {
//Short declaration + assigned values to variables x, y ,z
	x:= 42
	y:= "James, Bond"
	z:= true
//Print seperately
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
//Print multiply
	fmt.Printf("%d, %s, %t", x, y, z)
}
