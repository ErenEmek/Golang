package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	c := make (chan int, 1)
	
	c <- 42
	
	go func(){ 
		fmt.Println(<-c)
		wg.Done()
	}()
	
	wg.Wait()
}
