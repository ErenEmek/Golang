package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	c := gen()
	receive(c)
	wg.Wait()

}

func gen() <-chan int {
	c := make(chan int)
	go func() {		//we use func literal to encapsulate for loop from close function.
		for i := 0; i < 100; i++ {
		c <- i
		}
		close(c)	//deadlock happens if channel remains open.
	}()
wg.Done()
return c
}

func receive(c <-chan int) {

	for v := range c {
		fmt.Println(v)
	}
wg.Done()
}
