package main

import (
	"fmt"
)

func main() {
	func() {
		fmt.Println("Anonymous func")
	}()

	func(s string) {
		fmt.Println(s)
	}("anonymous func")

	s := func(s string) string {
		return s
	}("anonymous")
	fmt.Println(s)
}
