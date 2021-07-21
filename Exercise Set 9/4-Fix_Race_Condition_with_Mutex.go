package main

import (
	"fmt"
	// "runtime"
	"sync"
)

var counter int
var wg sync.WaitGroup
var mt sync.Mutex

const gs = 100

func main() {
wg.Add(gs)
	for i := 0; i < 100; i++ {
		go func() {
			mt.Lock()
			variable := counter
			// runtime.Gosched()
			variable++
			counter = variable
			mt.Unlock()
			wg.Done()
		}()
	}
	
	wg.Wait()
	fmt.Println(counter)
}
