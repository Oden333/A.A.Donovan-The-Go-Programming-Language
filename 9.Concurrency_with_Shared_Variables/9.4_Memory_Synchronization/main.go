package main

import (
	"fmt"
	"time"
)

// In a modern computer there may be dozens of processors, each with its own local cache of the main memory.
// For efficiency, writes to memory are buffered within each processor and flushed out to main memory only when necessary.
// They may even be committed to main memory in a different order than they were written by the writing goroutine

//* Synchronization primitives like channel communications and mutex operations
//* cause the processor to flush out and commit all its accumulated writes so
//* that the effects of goroutine execution up to that point are guaranteed to be visible to
//* goroutines running on other processors

func main() {
	var x, y int
	for {
		select {
		case <-time.Tick(time.Microsecond * 10):
			return
		default:
			go func() {
				x = 1                   // A1
				fmt.Print("y:", y, " ") // A2
			}()
			go func() {
				y = 1                   // B1
				fmt.Print("x:", x, " ") // B2
			}()
		}
	}
}

//? Within a single goroutine, the effects of each statement are guaranteed to occur in the
//? order of execution; goroutines are sequentially consistent.
// But in the absence of explicit synchronization using a channel or mutex,
// there is no guarantee that events are seen in the same order by all goroutines.

//* Although goroutine A must observe the effect of the write x = 1 before it reads the value of y,
//* it does not necessarily observe the write to y done by goroutine B, so A may print a stale value of y.

//* It  is  tempting  to  try  to  understand  concurrency  as  if  it  corresponds  to  some
//* interleaving of the statements of each goroutine, but as the example above shows, this
//* is  not  how  a  modern  compiler  or  CPU  works.  Because  the  assignment  and  the
//* Print refer to different variables, a compiler may conclude that the order of the two
//* statements cannot affect the result, and swap them. If the two goroutines execute on
//* different CPUs, each with its own cache, writes by one goroutine are not visible to
//* the other goroutine’s Print until the caches are synchronized with main memory.
