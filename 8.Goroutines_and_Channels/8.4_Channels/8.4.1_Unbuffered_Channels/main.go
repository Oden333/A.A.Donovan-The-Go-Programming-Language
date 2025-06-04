package main

import (
	"io"
	"log"
	"net"
	"os"
	"time"
)

//! A send operation on an unbuffered channel blocks the sending goroutine until another
//! goroutine executes a corresponding receive on the same channel, at which point the
//! value  is  transmitted  and  both  goroutines  may  continue.
//? Conversely,  if  the  receive operation  was  attempted  first,  the  receiving  goroutine  is  blocked
//? until  another goroutine performs a send on the same channel.

// Communication  over  an  unbuffered  channel  causes  the  sending  and  receiving goroutines to synchronize

//* When a value is sent on an unbuffered channel, the receipt of the value happens before the reawakening of the sending goroutine

// In  discussions  of  concurrency,  when  we  say  x  happens  before  y,  we  don’t  mean merely that x occurs earlier in time than y;
// we mean that it is guaranteed to do so and that all its prior effects, such as updates to variables,
// are complete and that you may rely on them

// ? When x neither happens before y nor after y, we say that x is concurrent with y
// This doesn’t  mean  that  x  and  y  are  necessarily  simultaneous,  merely  that  we  cannot assume anything about their ordering

func netcat3() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	// we  use  a  channel  to  synchronize  the  two goroutines (main + below)
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring		errors
		log.Println("done")
		done <- struct{}{} // signal the main		goroutine
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done // wait for background goroutine to finish
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

// Messages sent over channels have two important aspects. Each message has a value,
// but sometimes the fact of communication and the moment at which it occurs are just
// as important

func ex8_3() {
	// conn, err := net.Dial("tcp", "localhost:8000")
	conn, err := net.DialTCP("tcp", nil, &net.TCPAddr{IP: []byte{127, 0, 0, 1}, Port: 8000})
	if err != nil {
		log.Fatal(err)
	}
	// we  use  a  channel  to  synchronize  the  two goroutines (main + below)
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring		errors
		log.Println("done")
		done <- struct{}{} // signal the main		goroutine
	}()
	go func() {
		time.Sleep(time.Second * 5)
		os.Stdin.Close()
	}()

	mustCopy(conn, os.Stdin)
	<-done // wait for background goroutine to finish
	conn.CloseWrite()
}

func main() {
	ex8_3()
}
