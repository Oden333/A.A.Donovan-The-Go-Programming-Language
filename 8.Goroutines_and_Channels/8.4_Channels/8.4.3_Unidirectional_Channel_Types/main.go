package main

import (
	"fmt"
	"time"
)

//* The  type  chan<- int, a send-only  channel  of  int, allows sends but not receives.
//* The type   <-chan int, a receive-only channel of int, allows receives but not sends.

// Since the close operation asserts that no more sends will occur on a channel, only
// the sending goroutine is in a position to call it, and for this reason it is a compile-time
// error to attempt to close a receive-only channel.

func counter(out chan<- int) {
	for x := 0; x < 10; x++ {
		out <- x
		time.Sleep(time.Second / 3)
	}
	close(out)
}
func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}
func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}
func main() {
	naturals := make(chan int)
	squares := make(chan int)

	//? The call counter(naturals) implicitly converts naturals,
	//? a  value  of  type chan  int,  to  the  type  of  the  parameter,  chan<-  int.
	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}
