package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	v, ok := <-c 
	fmt.Println(v, ok)	//v holds value, ok holds the status of the channel.

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
}
