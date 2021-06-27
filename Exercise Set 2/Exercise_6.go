package main

import "fmt"
// print last 4 years using iota
const (
	y1 = iota + 2018
	y2
	y3
	y4
)
func main() {
	fmt.Println(y1, y2, y3, y4)

}
