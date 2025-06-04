package main

import "fmt"

// ! Channels  are  the connections between goroutines
//
// ! A channel is a communication mechanism that lets one goroutine send values to another goroutine.
// ? Each channel is a conduit for values of a particular  type,  called  the  channel’s  element  type.
// ? The  type  of  a  channel  whose elements have type int is written chan int
var ch = make(chan int)

//! A channel is a reference to the data structure created by make

//? When we  copy  a  channel  or  pass  one  as  an  argument  to  a  function,  we  are  copying  a
//? reference,  so  caller  and  callee  refer  to  the  same  data  structure.

var a = ch

func c() {
	//? Two channels of the same type may be compared using ==.
	//? The comparison is true if both  are  references  to  the  same  channel  data  structure
	fmt.Println(ch, a, ch == a)

	x := 3

	ch <- x  // a send statement
	x = <-ch // a receive expression in an assignment	statement
	<-ch     // a receive statement; result is discarded

	//? Channels support a third operation, close, which sets a flag indicating that no more values  will  ever  be  sent  on  this  channel;
	close(ch)
	//! subsequent  attempts  to  Send  will  panic.
	ch <- x

	//! Receive operations on a closed channel yield the values that have been sent until no more values are left
	//! any receive operations thereafter complete immediately and yield the zero value of the channel’s element type
	<-ch
}

// A channel created with a simple call to make is called an unbuffered channel, but
// make accepts an optional second argument, an integer called the channel’s capacity.
//? If the capacity is non-zero, make creates a buffered channel
