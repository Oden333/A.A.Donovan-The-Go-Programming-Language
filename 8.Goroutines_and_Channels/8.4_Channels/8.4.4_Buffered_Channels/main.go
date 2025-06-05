package main

import "fmt"

// A  buffered  channel  has  a  queue  of  elements.  The  queue’s  maximum  size  is
// determined  when  it  is  created,  by  the  capacity  argument  to  make()
var ch = make(chan string, 3)

//? A send operation on a buffered channel inserts an element at the back of the queue,
//? and a receive operation removes an element from the front.

func test() {
	ch <- "A"
	ch <- "B"
	ch <- "C"

	// If we receive one value,
	fmt.Println(<-ch) // "A"
	//? the  channel  is  neither  full  nor  empty ,  so  either  a  send  operation  or  a
	//? receive operation could proceed without blocking.
	//? In this way, the channel’s buffer decouples the sending and receiving goroutines.

	//* The channel’s buffer capacity can be obtained by calling the built-in cap function
	fmt.Println(cap(ch)) // "3"

	//* len function returns the number of elements currently buffered
	fmt.Println(len(ch)) // "2"

}

// Incidentally,  it’s  quite
// normal for several goroutines to send values to the same channel concurrently, as in
// this example, or to receive from the same channel.
func mirroredQuery() string {
	responses := make(chan string, 3)
	go func() { responses <- request("asia.gopl.io") }()
	go func() { responses <- request("europe.gopl.io") }()
	go func() { responses <- request("americas.gopl.io") }()
	//? Had we used an unbuffered channel, the two slower goroutines would have gotten
	//? stuck trying to send their responses on a channel from which no goroutine will ever receive.
	//! This  situation,  called  a  goroutine  leak,  would  be  a  bug.
	return <-responses // return the quickest response

	//! Unlike  garbage variables, leaked goroutines are not automatically collected
}
func request(hostname string) (response string) { /*
		... */
	return
}

//? Unbuffered channels  give  stronger  synchronization  guarantees
// because  every  send  operation  is synchronized with its corresponding receive
