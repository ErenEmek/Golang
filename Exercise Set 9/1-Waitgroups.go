package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup //waitgroup is defined with variable name wg.

func main() {

	wg.Add(2) //2 waitgroups added.

	go func() {
		fmt.Println("something is")
		wg.Done() //our routine is finished.
	}()

	go func() {
		fmt.Println("wrong")
		wg.Done() //our routine is finished.
	}()

	wg.Wait() //main routine waits until wg's done.

}
