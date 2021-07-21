package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	c := make(chan int) //we cannot send this as argument if we defined it as send or receive only channel.

	go func(c chan<- int) { //send only channel 
		c <- 42
		wg.Done()
	}(c)

	go func(c <-chan int) { //receive only channel
		fmt.Println(<-c)
		wg.Done()
	}(c)

	wg.Wait()
}
