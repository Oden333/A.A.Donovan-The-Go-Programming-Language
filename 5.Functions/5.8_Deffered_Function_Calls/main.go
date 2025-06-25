package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"golang.org/x/net/html"
)

// ! The function and argument expressions are evaluated when the statement is executed,
// & but the actual call is deferred until the function that contains
// & the defer statement has finished, whether normally, by executing a return statement
// & or falling off the end, or abnormally, by panicking.
func evaluatingDefer() {
	defer printElapsedTime()()
	// some tough logic here ...
	time.Sleep(300 * time.Millisecond)
}

func printElapsedTime() func() {
	start := time.Now() // Будет определено в момент регистрации defer
	return func() {
		fmt.Printf("Time elapsed: %s\n", time.Since(start))
	}
}

// Time elapsed: 301.044459ms
func Test_evaluatingDefer(t *testing.T) {
	evaluatingDefer()
}

//? Any number of calls may be deferred;
//! Deferred functions are invoked immediately before the surrounding function returns,
//! in the reverse order they were deferred

// A defer statement is often used with paired operations like open and close, connect
// and disconnect, or lock and unlock to ensure that resources are released in all cases,
// no matter how complex the control flow
func t() {
	// title("http://golang.org")
	// bigSlowOperation()
	// ClosurePassing()
	// UnnamedReturn()
	fmt.Println(NamedReturn())
}
func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	//? The right place for a defer statement that
	//? releases a resource is immediately after the resource has been successfully acquired
	defer resp.Body.Close()
	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		return fmt.Errorf("%s has type %s, not text/html", url, ct)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			fmt.Println(n.FirstChild.Data)
		}
	}
	forEachNode(doc, visitNode, nil)
	return nil
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

// & The defer statement can also be used to pair “on entry” and “on exit”
// & actions when debugging a complex function
func bigSlowOperation() {
	// Here we calls trace immediately, which does the “on entry” action then returns
	// a function value that, when called, does the corresponding “on exit” action
	defer trace("bigSlowOperation")()

	//&	But don’t forget the final parentheses in the defer statement, or the
	//& “on entry” action will happen on exit and the on-exit action won’t happen at all!
	//* defer trace("bigSlowOperation")

	// ...lots of work...
	time.Sleep(10 * time.Second) // simulate slow operation by sleeping
}
func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg,
			time.Since(start))
	}
}

//! Deferred functions run after return statements have updated the function’s result variables.

// Because an anonymous function can access its enclosing function’s
// variables, including named results, a deferred anonymous function can observe the function’s results.

func double(x int) (result int) {
	defer func() {
		fmt.Printf("double(%d) = %d\n", x, result)
	}()
	return x + x
}

// _ = double(4)
// Output:
// "double(4) = 8"

// & Because deferred functions aren’t executed until the very end of a function’s
// & execution, a defer statement in a loop deserves extra scrutiny.
func tests() error {
	filenames := os.Args[1:]
	for _, filename := range filenames {
		f, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer f.Close() // NOTE: risky; could run out of file descriptors
		// ...process f...
	}

	for _, filename := range filenames {
		if err := doFile(filename); err != nil {
			return err
		}
	}

	return nil
}

func doFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	// ...process f...
	return err
}
