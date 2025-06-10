// The du1 command computes the disk usage of the files in a directory.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	start := time.Now()
	defer func() { fmt.Printf("Elapsed: %.2f ms\n", time.Since(start).Seconds()) }()

	// Determine the initial directories.
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	{
		// Traverse the file tree.
		// fileSizes := make(chan int64)
		// go func() {
		// 	for _, root := range roots {
		// 		walkDir(root, fileSizes)
		// 	}
		// 	close(fileSizes)
		// }()
	}

	// Traverse each root of the file tree in parallel.
	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir1(root, &n, fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()

	{
		// stop := make(chan struct{})
		// go spinner(stop)
		// Print the results.
		// var nfiles, nbytes int64
		// for size := range fileSizes {
		// 	nfiles++
		// 	nbytes += size
		// }
		// stop <- struct{}{}
	}

	// Print the results periodically.
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(5 * time.Second)
	}
	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was closed
			}
			nfiles++
			nbytes += size
		case <-tick: //* select statement whose channel is nil is never selected
			printDiskUsage(nfiles, nbytes)
		}
	}

	printDiskUsage(nfiles, nbytes) // final totals

}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.2f GB\n", nfiles, float64(nbytes)/1e9)
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		} else {
			info, err := entry.Info()
			if err != nil {
				log.Println(err)
				continue
			}
			fileSizes <- info.Size()
		}
	}
}

// dirents returns the entries of directory dir.
func dirents(dir string) []os.DirEntry {
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

func walkDir1(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents1(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			n.Add(1)
			go walkDir1(subdir, n, fileSizes)
		} else {
			info, err := entry.Info()
			if err != nil {
				log.Println(err)
				continue
			}
			fileSizes <- info.Size()
		}
	}
}

// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents1(dir string) []os.DirEntry {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token
	return dirents(dir)
}

func spinner(stopch <-chan struct{}) {
	t := time.Now()
	for {
		select {
		case <-stopch:
			fmt.Print("\r")
			return
		default:
			for _, r := range `-\|/` {
				fmt.Printf("\r%c Elapsed %.6sms", r, time.Since(t).Round(time.Microsecond*10))
				time.Sleep(time.Millisecond * 300)
			}
		}
	}
}
