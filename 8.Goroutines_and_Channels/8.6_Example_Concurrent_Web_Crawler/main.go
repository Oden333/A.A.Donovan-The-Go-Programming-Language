package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main1() {

	worklist := make(chan []string) // a worklist records the queue of items that need processin
	// Start with the command-line arguments.
	go func() { worklist <- os.Args[1:] }()
	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- crawl0(link)
				}(link)
			}
		}
	}
}

// The initial error message is a surprising report of a DNS lookup failure for a reliable
// domain.  The  subsequent  error  message  reveals  the  cause:  the  program  created  so
// many  network  connections  at  once  that  it  exceeded  the  per-process  limit  on  the
// number  of  open  files,  causing  operations  such  as  DNS  lookups  and  calls  to net.Dial to start failing

//? The program is too parallel.
//! Unbounded parallelism is rarely a good idea since there is always a limiting factor in the system, such as
//! - the number of CPU cores for compute- bound workloads,
//! - the number of spindles and heads for local disk I/O operations,
//! - thebandwidth of the network for streaming downloads, or
//! - the serving capacity of a web service

//* We  can  limit  parallelism  using  a  buffered  channel  of  capacity  n
//* to  model  a concurrency  primitive  called  a  counting  semaphore.

// Conceptually, each of the n vacant slots in the channel buffer represents a token entitling the holder to proceed.
// tokens is a counting semaphore used to

// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)

func crawl2(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := extract(url)
	<-tokens // release the token
	if err != nil {
		log.Print(err)
	}
	return list
}

func counting_semaphore() {
	worklist := make(chan []string)
	var n int // number of pending sends to worklist
	// Start with the command-line arguments
	n++
	go func() { worklist <- os.Args[1:] }()
	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl2(link)
				}(link)
			}
		}
	}
}

func worker_pool() {
	worklist := make(chan []string)  // lists of URLs, may have duplicates
	unseenLinks := make(chan string) // de-duplicated URLs
	// Add command-line arguments to worklist.
	go func() { worklist <- os.Args[1:] }()
	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		//* The crawler goroutines are all fed by the same channel, unseenLinks
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl2(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	//* The main goroutine de-duplicates worklist items and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] { //The seen map can be accessed onlyby  that  goroutine
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}

func main() {
	// worker_pool()
	// solution_8_6()
	ex8_6()
}

var depth = flag.Int("depth", 5, "limits the depth of crawler's search operation")

func ex8_6() {
	flag.Parse()
	worklist := make(chan []dep)  // lists of URLs, may have duplicates
	unseenLinks := make(chan dep) // de-duplicated URLs
	// Add command-line arguments to worklist.
	args := flag.Args()
	startSlice := make([]dep, len(args))
	for _, s := range args {
		startSlice = append(startSlice, dep{s, 0})
	}
	go func() { worklist <- startSlice }()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		//* The crawler goroutines are all fed by the same channel, unseenLinks
		go func() {
			for d := range unseenLinks {
				go func() { worklist <- crawl4(d) }()
			}
		}()
	}

	//* The main goroutine de-duplicates worklist items and sends the unseen ones to the crawlers.

	seen := make(map[string]bool)

	for list := range worklist {
		for _, link := range list {
			if !seen[link.link] {
				seen[link.link] = true
				unseenLinks <- dep{link: link.link, dep: link.dep}
			}
		}
	}

	close(unseenLinks)
}

type dep struct {
	link string
	dep  int
}

func crawl4(l dep) []dep {
	fmt.Println(l.dep, l.link)
	var links []dep
	if l.dep < *depth {
		depth := l.dep + 1
		list, err := extract(l.link)
		if err != nil {
			log.Print(err)
		}

		for _, url := range list {
			links = append(links, dep{url, depth})
		}
	}
	return links
}
