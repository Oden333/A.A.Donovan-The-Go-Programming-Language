package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

func count() {
	fmt.Println("Commencing countdown.  Press return to abort.")

	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()

	// The  time.Tick function  returns  a  channel  on  which  it  sends  events  periodically,  acting  like  a metronome.
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
			// Do nothing.
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	}

	fmt.Println("launch()")
}

//? Like a switch statement, "select" has  a  number  of  cases  and  an  optional  default.
//? Each  case  specifies  a communication  (a  send  or  receive  operation  on  some  channel)  and  an  associated block of statements.

//*   A select  with  no  cases,  select{},  waits forever.

func selects() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		switch { //! Switch cases picked in order
		case i > 0:
			fmt.Println(0)
		case i > 1:
			fmt.Println(1)
		}

		//! If multiple cases are ready, select picks one at random,
		//! which ensures that every channel  has  an  equal  chance  of  being  selected.

		select { //! Switch cases picked randomly
		case x := <-ch:
			fmt.Println("---", x) // "0" "2" "4" "6" "8"
		case ch <- i:
		}
	}
}

// ? When  the countdown  function  above  returns,  it  stops  receiving  events  from  tick,
// ? but  the ticker  goroutine  is  still  there,  trying  in  vain  to  send  on  a  channel
// ? from  which  no goroutine is receiving — a goroutine leak
func cl() {
	ticker := time.NewTicker(1 * time.Second)
	fmt.Println(<-ticker.C) // receive from the ticker's channel
	ticker.Stop()           // cause the ticker's goroutine to	terminate
}

// Sometimes we want to try to send or receive on a channel but avoid blocking
// if the channel is not ready—a non-blocking communication

//? The zero value for a channel is nil. Perhaps surprisingly, nil channels are sometimes useful.
//? Because send and receive operations on a nil channel block forever, a case in
//? a  select  statement  whose  channel  is  nil  is  never  selected.  This  lets  us  use  nil  to
//? enable  or  disable  cases  that  correspond  to  features  like  handling  timeouts  or
//? cancellation,  responding  to  other  input  events,  or  emitting  output.

func ex8_8() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		log.Printf("Connection %s accepted", conn.RemoteAddr())
		go handleConn(conn) // handle connections 		concurrently
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	scanner := bufio.NewScanner(c)
	timeOut := time.NewTicker(time.Second * 5)

	input := make(chan string)
	wg := sync.WaitGroup{}
	defer func() {
		c.Close()
		wg.Wait()
		log.Printf("Connection %s closed", c.RemoteAddr())
	}()

	go func() {
		for scanner.Scan() {
			input <- scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			log.Println(err)
		}
	}()
	for {
		select {
		case <-timeOut.C:
			timeOut.Stop()
			return
		case msg := <-input:
			wg.Add(1)
			go echo(c, msg, 1*time.Second)
			timeOut.Reset(time.Second * 5)
		}
	}
}

func main() {
	ex8_8()
}
