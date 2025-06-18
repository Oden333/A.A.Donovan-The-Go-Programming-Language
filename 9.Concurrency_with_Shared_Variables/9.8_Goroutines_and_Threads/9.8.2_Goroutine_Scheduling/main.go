package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

// OS threads are scheduled by the OS kernel. Every few milliseconds, a hardware timer
// interrupts the processor, which causes a kernel function called the scheduler to be invoked.
//
// This function suspends the currently executing thread and saves its registers
// in memory, looks over the list of threads and decides which one should run next,
// restores that thread’s registers from memory, then resumes the execution of that thread.
//
//? Because OS threads are scheduled by the kernel, passing control from one
//? thread to another requires a full context switch, that is, saving the state of one user
//? thread to memory, restoring the state of another, and updating the scheduler’s data structures.
//
//? This operation is slow, due to its poor locality and the number of memory
//? accesses required, and has historically only gotten worse as the number of CPU
//? cycles required to access memory has increased.

//
//! The Go runtime contains its own scheduler that uses a technique known as m:n
//! scheduling, because it multiplexes (or schedules) m goroutines on n OS threads. The
//! job of the Go scheduler is analogous to that of the kernel scheduler, but it is
//! concerned only with the goroutines of a single Go program.

//! The Go scheduler is not invoked periodically by a hardware timer, but implicitly by certain Go language constructs.

// For example, when a goroutine calls time.Sleep or blocks in a channel or mutex
// operation, the scheduler puts it to sleep and runs another goroutine until it is time to
// wake the first one up. Because it doesn’t need a switch to kernel context,
// rescheduling a goroutine is much cheaper than rescheduling a thread.

func main() {
	start := time.Now()
	var counter atomic.Int64
	ch := make(chan struct{})
	go func() {
		for {
			<-ch
			counter.Add(1)
			// fmt.Printf("\r%d")
			ch <- struct{}{}
		}
	}()

	go func() {
		for {
			ch <- struct{}{}
			counter.Add(1)
			// fmt.Printf("\r%d")
			<-ch
		}
	}()

	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Printf("Done %d operations %.2f\n", counter.Load(), time.Since(start).Seconds())
			counter.Swap(0)
		}
	}
}
