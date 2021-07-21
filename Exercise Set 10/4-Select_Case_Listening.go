package main

import (
	"fmt"
)

func main() {
	q := make(chan int)	//defined channel q
	c := gen(q)		//defined channel c as gen(q) returns receive only channel. this will feed c with numbers.
	receive(c, q)		//pass c and q to receive function. q channel carries interrupt signal. Value of q is passed without return, etc.

	fmt.Println("about to exit")

}

func gen(q chan int) <-chan int {
	d := make(chan int)	//defined new channel d

	go func() {
		for i := 0; i < 100; i++ {
			d <- i	//filled d with values.
		}
		q <- 1		//after finished feeding d, we feed q channel to send finish message.
	}()
	return d		//return d as channel.
}

func receive(a <-chan int, b <-chan int) {
	for {
		select {
		case v := <-a:
			fmt.Println(v)
		case <-b:	//no need to define variable.
			return
		}
	}

}
