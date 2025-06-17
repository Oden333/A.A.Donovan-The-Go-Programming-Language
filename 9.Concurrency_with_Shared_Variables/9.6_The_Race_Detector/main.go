package main

import (
	"fmt"
	"time"
)

// Go runtime and toolchain are equipped with a sophisticated and easy-to-use dynamic analysis tool, the race detector
// Just add the -race flag to your go build, go run, or go test command
// ? This causes the compiler to build a modified version of your application or test with
// ? additional instrumentation that effectively records all accesses to shared variables that
// ? occurred during execution, along with the identity of the goroutine that read or wrote the variable.
//
// ? In addition, the modified program records all synchronization events,
// ? such as go statements, channel operations, and calls to (*sync.Mutex).Lock, (*sync.WaitGroup).Wait, and so on.
func main() {
	var x, y map[int]int = make(map[int]int), map[int]int{}
	for {
		select {
		case <-time.Tick(time.Microsecond * 100):
			return
		default:
			go func() {
				x[1] = 1                // A1
				fmt.Print("y:", y, " ") // A2
			}()
			go func() {
				y[1] = 1                // B1
				fmt.Print("x:", x, " ") // B2
			}()
		}
	}
}

// The race detector reports all data races that were actually executed. However, it can
// only detect race conditions that occur during a run; it cannot prove that none will ever occur.

// Due to extra bookkeeping, a program built with race detection needs more time and
// memory to run, but the overhead is tolerable even for many production jobs
