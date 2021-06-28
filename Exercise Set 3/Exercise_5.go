package main

import "fmt"

func main () {
	//remainder for each number btw 10 and 100 divided by 4
	for i:= 10 ; i < 100 ; i++ {
		fmt.Println(i, " ", i%4)
	}
}