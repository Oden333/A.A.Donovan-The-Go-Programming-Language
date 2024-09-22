// Server1 is a minimal "echo" server.
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

// Server2 is a minimal "echo" and counter server.
var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

// func main() {
// 	http.HandleFunc("/", handler) // each request calls handler
// 	log.Fatal(http.ListenAndServe("localhost:8080", nil))
// }

// // handler echoes the Path component of the requested URL.
// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
// }
