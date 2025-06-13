package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

//* For multiple cancellation we need a reliable mechanism to broadcast an event over a channel
//* so that many goroutines can see it as it occurs and can later see that it has occurred

// ! Recall that after a channel has been closed and drained of all sent values, subsequent
// ! receive operations proceed immediately, yielding zero values.
// ? We can exploit this to create a broadcast mechanism: don’t send a value on the channel, close it.

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	start := time.Now()
	defer func() { fmt.Printf("Elapsed: %d ms\n", time.Since(start).Milliseconds()) }()

	// Determine the initial directories.
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	/// Traverse each root of the file tree in
	// 	parallel.
	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes)
	}

	go func() {
		n.Wait()
		close(fileSizes)
	}()

	// Cancel traversal when input is detected.
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single	byte
		close(done)
	}()

	// Print the results periodically.
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles, nbytes int64
loop:
	for {
		select {
		case <-done:
			// Drain fileSizes to allow existing goroutines to finish.

			//? The function returns if this case is ever selected, but before it returns
			//! it must first drain the fileSizes channel, discarding all values until the channel is closed.
			//! It does this to ere thatnsu any active calls to walkDir can run to completion
			//! without getting stuck sending to fileSizes.
			for range fileSizes {
				// Do nothing.
			}
			return
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was closed
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes) // final totals
}
func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles,
		float64(nbytes)/1e9)
}

func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	// The  walkDir  goroutine  polls  the  cancellation  status  when  it  begins,  and  returns
	// without  doing  anything  if  the  status  is  set.
	if cancelled() {
		return
	}
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			//* Cancellation involves a trade-off;a quicker response often requires more intrusive changesto program logic
			if !cancelled() {
				n.Add(1)
				go walkDir(subdir, n, fileSizes)
			}

		} else {
			fileSizes <- entry.Size()
		}
	}
}

var sema = make(chan struct{}, 20) // concurrency-limiting counting semaphore

// dirents returns the entries of directory dir.
// !+5
func dirents(dir string) []os.FileInfo {
	select {
	case sema <- struct{}{}: // acquire token
	case <-done:
		return nil // cancelled
	}
	defer func() { <-sema }() // release token

	// ...read directory...
	//!-5

	f, err := os.Open(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	defer f.Close()

	entries, err := f.Readdir(0) // 0 => no limit; read all entries
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		// Don't return: Readdir may return partial results.
	}
	return entries
}
