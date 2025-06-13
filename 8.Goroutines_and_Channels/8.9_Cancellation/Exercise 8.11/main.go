package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

// Following  the  approach  of  mirroredQuery  in  Section  8.4.4,
// implement a variant of fetch that requests several URLs concurrently.
// As soon as the first response arrives, cancel the other requests

func main() {
	cancel := make(chan struct{})
	done := make(chan string, len(os.Args[1:]))
	for _, url := range os.Args[1:] {
		go func() {
			done <- fetch(url, cancel)
		}()
	}

	timeout := time.Tick(time.Second * 5)
	for {
		select {
		case host := <-done:
			if host != "" {
				log.Printf("Done; Created page of %q", host)
				close(cancel)
				return
			}
		case <-timeout:
			log.Println("Timeout")
			close(cancel)
			return
		}
	}
}

func fetch(url string, cancel <-chan struct{}) (filename string) {
	if !strings.HasPrefix(url, "https") {
		url = "https://" + url
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
		return
	}
	req.Cancel = cancel
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		log.Printf("%s responsed with failure status %d", req.Host, resp.StatusCode)
		return
	}
	defer resp.Body.Close()
	file, err := os.Create(fmt.Sprintf("%s.html", req.Host))
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	io.Copy(file, resp.Body)
	filename = req.Host
	return
}
