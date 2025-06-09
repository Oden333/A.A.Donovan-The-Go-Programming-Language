package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"sync"
)

// Write a concurrent program that creates a local mirror of a web site,
// fetching each reachable page and writing it to a directory on the local disk.

// Only pages within  the  original  domain  (for  instance,  golang.org)  should  be  fetched.

// URLs within mirrored pages should be altered as needed so that they refer to the mirrored page, not the original.

// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)
var seen = make(map[string]bool)
var seenLock = sync.Mutex{}
var start *url.URL

var maxDepth = flag.Int("d", 3, "max crawl depth")

func main() {
	flag.Parse()
	wg := &sync.WaitGroup{}
	if len(flag.Args()) == 0 {
		fmt.Fprintln(os.Stderr, "usage: mirror URL ...")
		os.Exit(1)
	}

	u, err := url.Parse(flag.Arg(0))
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid url: %s\n", err)
	}
	start = u

	for _, link := range flag.Args() {
		wg.Add(1)
		go crawl(link, 1, wg)
	}

	wg.Wait()
}
