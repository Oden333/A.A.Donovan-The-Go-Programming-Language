package main

import (
	"fmt"

	"golang.org/x/net/html"
)

//! If the built-in recover function is called within a deferred function and the function
//! containing the defer  statement  is  panicking,  recover  ends  the  current  state  of
//! panic and returns the panic value.
//! The function that was panicking does not continue where it left off but returns normally.
//? If recover is called at any other time, it has no effect and returns nil.

//? Recovering indiscriminately from panics is a dubious practice because the state of a
//? package’s  variables  after  a  panic  is  rarely  well  defined  or  documented.

//& Recovering from a panic within the same package can help simplify the handling of
//& complex  or  unexpected  errors,  but  as  a  general  rule,
//! you  should  not  attempt  to recover from another package’s panic.

//& Similarly, you should not recover from a panic that may pass through a function you do not maintain,
//& such as a caller-provided callback, since you cannot reason about its safety

// For  example,  the  net/http  package  provides  a  web  server  that  dispatches
// incoming requests to user-provided handler functions. Rather than let a panic in one
// of these handlers kill the process, the server calls recover, prints a stack trace, and
// continues serving.
// This is convenient in practice, but it does risk leaking resources or
// leaving the failed handler in an unspecified state that could lead to other problems.

// It’s safest to recover selectively:
// !Recover only from panics that were intended to be recovered from, which should be rare

// This intention can be encoded by using a distinct, unexported type for the panic
// value and testing whether the value returned by recover has that type.

// soleTitle returns the text of the first non-empty title element in doc,
// and an error if there was not exactly one.
func soleTitle(doc *html.Node) (title string, err error) {
	type bailout struct{}
	defer func() {
		switch p := recover(); p {
		case nil:
			// no panic
		case bailout{}:
			// "expected" panic
			err = fmt.Errorf("multiple title elements")
		default:
			panic(p) // unexpected panic; carry on panicking
		}
	}()
	// Bail out of recursion if we find more than one non-empty title.
	forEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			if title != "" {
				panic(bailout{}) // multiple title elements
			}
			title = n.FirstChild.Data
		}
	}, nil)
	if title == "" {
		return "", fmt.Errorf("no title element")
	}
	return title, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

//! From some conditions there is no recovery. Running  out  of  memory,  for  example,
//! causes the Go runtime to terminate the program with a fatal error.

func ex5_19(num *int) {
	defer func() {
		switch r := recover(); r {
		case nil:

		case 8, 5:
			n := r.(int)
			*num = n
		default:
			panic(r)
		}
	}()

	for i := range 10 {
		if i == 5 {
			panic(i)
		}
		if i == 8 {
			panic(i)
		}
	}

}
func main() {
	num := 0
	ex5_19(&num)
	fmt.Println(num)
}
