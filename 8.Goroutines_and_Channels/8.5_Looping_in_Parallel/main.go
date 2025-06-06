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

	"gopl.io/ch8/thumbnail"
)

// makeThumbnails makes thumbnails of the specified files.
func makeThumbnails(filenames []string) {
	for _, f := range filenames {
		if _, err := thumbnail.ImageFile(f); err !=
			nil {
			log.Println(err)
		}
	}
}

// makeThumbnails3 makes thumbnails of the specified files in parallel.
func makeThumbnails3(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		// We passed the value of f as an explicit argument to the literal function
		// instead of using the declaration of f from the enclosing for loop:

		//! The single variable f is shared by all the anonymous function
		//! values and updated by successive loop iterations.
		go func(f string) {
			thumbnail.ImageFile(f) // NOTE: ignoring errors
			ch <- struct{}{}
		}(f)
	}
	// Wait for goroutines to complete.
	for range filenames {
		<-ch
	}
}

// makeThumbnails4 makes thumbnails for the specified files in parallel.
// It returns an error if any step failed.
func makeThumbnails4(filenames []string) error {
	errors := make(chan error)
	for _, f := range filenames {
		go func(f string) {
			_, err := thumbnail.ImageFile(f)
			errors <- err
		}(f)
	}
	for range filenames {
		if err := <-errors; err != nil {
			// This function has a subtle bug. When it encounters the first non-nil error, it returns  the  error  to  the  caller,
			//! leaving  no  goroutine  draining  the  errors  channel.
			return err // NOTE: incorrect: goroutine leak!
		}
	}
	return nil
}

// The  simplest  solution  is  to  use  a  buffered  channel  with  sufficient  capacity  that  no
// worker goroutine will block when it sends a message.

// makeThumbnails5 makes thumbnails for the specified files in parallel.
// It returns the generated file names in an arbitrary order, or an error if any step failed.
func makeThumbnails5(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbfile string
		err       error
	}
	ch := make(chan item, len(filenames))
	for _, f := range filenames {
		go func(f string) {
			var it item
			it.thumbfile, it.err = thumbnail.ImageFile(f)
			ch <- it
		}(f)
	}
	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}
	return thumbfiles, nil
}

//* To  know  when  the  last  goroutine  has  finished  (which  may  not  be  the  last  one  to
//* start), we need to increment a counter before each goroutine starts and decrement it
//* as each goroutine finishes.

// makeThumbnails6 makes thumbnails for each file received from the channel.
// It returns the number of bytes occupied by the files it creates.
func makeThumbnails6(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup // number of working	goroutines
	for f := range filenames {
		wg.Add(1) // must be called before the worker goroutine starts, not within it;
		// worker
		go func(f string) {
			defer wg.Done()
			thumb, err := thumbnail.ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb) // OK to ignore error
			sizes <- info.Size()
		}(f)
	}

	// closer
	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += size
	}

	return total
}

func handleConn(c net.Conn) {
	wg := sync.WaitGroup{}
	input := bufio.NewScanner(c)
	for input.Scan() {
		wg.Add(1)
		go func() {
			echo(c, input.Text(), 1*time.Second)
			wg.Done()
		}()
	}
	wg.Wait()
	// NOTE: ignoring potential errors from input.Err()
	c.Close()
}

// !-
func main() {
	ex8_5()
}
func ex8_4() {
	l, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func ex8_5() {

}
