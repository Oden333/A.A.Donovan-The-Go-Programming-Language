package main

import (
	"fmt"
	"time"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	// Counter
	go func() {
		// If the sender knows that no further values will ever be sent on a channel, it is useful
		// to communicate this fact to the receiver goroutines so that they can stop waiting
		for x := 0; ; x++ {
			time.Sleep(time.Second / 2)
			if x == 10 {
				close(naturals)
				break
			}
			naturals <- x
		}
	}()

	// Squarer
	squarer := func() {
		for {
			x, ok := <-naturals
			if !ok {
				break // channel was closed and drained
			}
			squares <- x * x
		}
		close(squares)
		//! After a channel has been closed, any further send operations on it will panic
	}
	go squarer()

	// Printer (in main goroutine)

	for x := range squares {
		fmt.Println(x)
	}
	//! After the losed  channel  has  been  drained, (after  the  last  sent  element  has  been  received)
	//! all subsequent receive operations will proceed without blocking but will yield
	//! a zero value

}

//? You needn’t close every channel when you’ve finished with it. It’s only necessary to
//? close a channel when it is important to tell the receiving goroutines that all data have
//? been sent.
//! A channel that the garbage collector determines to be unreachable will have
//! its resources reclaimed whether or not it is closed
