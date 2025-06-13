// The du4 command computes the disk usage of the files in a directory.
package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"
)

var depth = flag.Int("depth", 3, "limits the depth of crawler's search operation")

func main() {
	flag.Parse()
	worklist := make(chan []depL)  // lists of URLs, may have duplicates
	unseenLinks := make(chan depL) // de-duplicated URLs
	// Add command-line arguments to worklist.
	args := flag.Args()
	startSlice := make([]depL, len(args))
	for _, s := range args {
		if _, err := url.Parse(s); err != nil {
			log.Fatal("Invalid link", err)
		}
		startSlice = append(startSlice, depL{"https://gopl.io", 0})
	}
	go func() {
		worklist <- startSlice
	}()

	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single	byte
		close(done)
	}()

	for i := 0; i < 20; i++ {
		go func(i int) {
			for {
				select {
				case d, ok := <-unseenLinks:
					if !ok {
						return
					}
					if !cancelled() {
						go func() {
							worklist <- crawl(d)
						}()
					}
				case <-done:
					log.Println("Interrupted gorutine", i)
					return
				}
			}
		}(i)
	}

	//* The main goroutine de-duplicates worklist items and sends the unseen ones to the crawlers.

	seen := make(map[string]bool)

loop:
	for {
		select {
		case list := <-worklist:
			for _, link := range list {
				if !seen[link.link] {
					seen[link.link] = true
					if !cancelled() {
						unseenLinks <- depL{link: link.link, dep: link.dep}
					}
				}
			}
		case <-done:
			go func() {
				fmt.Println("\n\nShutting down")
				time.Sleep(time.Second * 5) // Timeout
				close(worklist)
			}()
			for range worklist {
			}
			break loop
		}

	}
	close(unseenLinks)
}

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

type depL struct {
	link string
	dep  int
}

func crawl(l depL) []depL {
	fmt.Println(l.dep, l.link)
	var links []depL
	if l.dep < *depth {
		depth := l.dep + 1
		list, err := extract(l.link)
		if err != nil {
			log.Print(err)
		}

		for _, url := range list {
			links = append(links, depL{url, depth})
		}
	}
	return links
}
