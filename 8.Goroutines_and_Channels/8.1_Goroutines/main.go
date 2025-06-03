package main

import (
	"fmt"
	"time"
)

// Go enables two styles of concurrent programming.
// This chapter presents goroutines and channels, which support communicating  sequential  processes or CSP,
// a model of  concurrency  in  which  values  are  passed  between  independent  activities
// (goroutines) but variables are for the most part confined to a single activity

//! When a program starts, its only goroutine is the one that calls the main function, so we  call  it  the  main  goroutine.

// ! The go statement itself completes immediately
func main() {
	go spinner(200 * time.Millisecond)
	const n = 50
	fibN := fib(n) // slow
	// The  main  function  then  returns.
	// When  this  happens,  all  goroutines  are  abruptly terminated and the program exits.

	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	t := time.Now()
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c %.6sms", r, time.Since(t).Round(time.Microsecond*10))
			t = time.Now()
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
