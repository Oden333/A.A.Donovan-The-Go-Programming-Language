package main

import (
	"fmt"
)

//
//! The Go scheduler uses a parameter called GOMAXPROCS to determine how many
//! OS threads may be actively executing Go code simultaneously
//
//? Its default value is the number of CPUs on the machine
// GOMAXPROCS is the n in m:n scheduling

//* Goroutines that are sleeping or blocked in a communication do not need a thread at all.
//* Goroutines that are blocked in I/O or other system calls or are calling non-Go functions,
// do need an OS thread, but GOMAXPROCS need not account for them

func main() {
	var counter int
	for {
		fmt.Print(1)

		counter++
		go fmt.Print(counter % 3)
	}
}

//? You can explicitly control this parameter using the GOMAXPROCS environment variable
//? or the runtime.GOMAXPROCS function.

// In the first run, at most one goroutine was executed at a time
//* GOMAXPROCS=1 go run main.go
//* 111111111111111111110000000000000000000011111...
// After a period of time, the Go scheduler put it to sleep
// and woke up the goroutine that prints zeros, giving it a turn to run on the OS thread
