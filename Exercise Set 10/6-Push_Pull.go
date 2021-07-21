package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {

		for i := 0; i < 100; i++ { //pushes 100 values to channel c
			c <- i
		}
		close(c)

	}()

	for v := range c { //pulls the numbers from channel c and prints
		fmt.Println(v)
	}

}
