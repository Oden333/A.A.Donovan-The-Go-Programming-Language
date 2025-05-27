package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// The  ListenAndServe  function  requires
// a  server  address,  such  as "localhost:8000",
// an  instance  of  the  Handler  interface  to  which  all requests should be dispatched.
// It runs forever, or until the server fails (or fails to start) with an error, always non-nil, which it returns
func maina() {
	db := database{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

// type database map[string]dollars

func (db database) ServeHTTP1(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) ServeHTTP(w http.ResponseWriter,
	req *http.Request) {
	switch req.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/price":
		item := req.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound) // 404
			fmt.Fprintf(w, "no such item: %q\n", item)
			return
		}
		fmt.Fprintf(w, "%s\n", price)
	default:
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such page: %s\n", req.URL)
	}
}

// Go doesn’t have a canonical web framework analogous  to  Ruby’s  Rails  or  Python’s  Django.
// This  is  not  to  say  that  such frameworks don’t exist,
// but the building blocks in Go’s standard library are flexible enough that frameworks are often unnecessary

func mains() {
	db := database{"shoes": 50, "socks": 5}
	// ? ServeMux - a  request  multiplexer,  to  simplify  the  association  between URLs and handlers.
	mux := http.NewServeMux()
	// mux.Handle("/list", http.HandlerFunc(db.list))
	mux.HandleFunc("/list", db.list)
	// mux.Handle("/price", http.HandlerFunc(db.price))
	mux.HandleFunc("/price", db.price)
	mux.HandleFunc("/update", db.update)
	mux.HandleFunc("/delete", db.delete)

	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

type database map[string]dollars

// ? db.list is a function that implements handler-like behavior,
// ? but since it has no methods, it doesn’t satisfy the http.Handler  interface  and  can’t  be  passed directly to mux.Handle.
func (db database) list0(w http.ResponseWriter, req *http.Request) {

	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	qs := req.URL.Query()
	if i := qs.Get("item"); i != "" {
		delete(db, i)
	}
	w.WriteHeader(http.StatusOK)
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	qs := req.URL.Query()
	if i, q := qs.Get("item"), qs.Get("price"); i != "" {
		if _, ok := db[i]; !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		qf, err := strconv.ParseFloat(q, 32)
		if err != nil {
			panic(err)
		}
		db[i] = dollars(qf)
	}
	w.WriteHeader(http.StatusOK)
}

// HandlerFunc demonstrates some unusual features of Go’s interface mechanism.
// It is a function type that has methods and satisfies an interface, http.Handler

//? HandlerFunc  is  thus  an  adapter  that  lets  a  function  value  satisfy  an  interface,
//? where the function and the interface’s sole method have the same signature.

//? So,  for  convenience,  net/http  provides
//? a  global  ServeMux  instance  called DefaultServeMux
//? and  package-level  functions  called  http.Handle
//? and http.HandleFunc.

// To use DefaultServeMux as the server’s main handler,
// we needn’t pass it to ListenAndServe; nil will do.
var tpl = template.Must(template.ParseFiles("./tmpl.tpl"))

func main() {
	db := database{"shoes": 50, "socks": 5}
	// ? ServeMux - a  request  multiplexer,  to  simplify  the  association  between URLs and handlers.
	mux := http.NewServeMux()
	// mux.Handle("/list", http.HandlerFunc(db.list))
	mux.HandleFunc("/list", db.list)
	// mux.Handle("/price", http.HandlerFunc(db.price))
	mux.HandleFunc("/price", db.price)
	mux.HandleFunc("/update", db.update)
	mux.HandleFunc("/delete", db.delete)

	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

//! The  web  server invokes each handler in a new goroutine, so handlers must take precautions such as
//! locking when accessing variables that other goroutines, including other requests to the same  handler,  may  be  accessing.

func (db database) list(w http.ResponseWriter, req *http.Request) {
	// tpl.Execute(os.Stdout, db)
	tpl.Execute(w, db)
}
