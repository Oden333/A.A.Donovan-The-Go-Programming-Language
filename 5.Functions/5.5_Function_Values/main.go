package main

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

//& Functions are first-class values in Go: like other values, function values have types,
//& and  they  may  be  assigned  to  variables  or  passed  to  or  returned  from  functions.
//& A function value may be called like any other function.

func square(n int) int     { return n * n }
func negative(n int) int   { return -n }
func product(m, n int) int { return m * n }

func basics() {
	f := square
	fmt.Println(f(3)) // "9"
	f = negative
	fmt.Println(f(3))     // "-3"
	fmt.Printf("%T\n", f) // "func(int) int"
	//* f = product           // compile error: can't assign f(int, int) int to f(int) int

	//& The zero value of a function type is nil
	var fn func(int) int
	//may cause panic: call of nil function

	//! Function values may be compared with nil:
	if fn != nil {
		f(3)
	}
	//! but they are not comparable, so they may not be compared against each other or used as keys in a map

}
func add1(r rune) rune { return r + 1 }

func example() {
	//& Function values let us parameterize our functions over not just data, but behavior too.
	// For example, strings.Map applies a function to each character of a string,
	// joining the results to make another string.

	fmt.Println(strings.Map(add1, "HAL-9000")) // "IBM.:111"
	fmt.Println(strings.Map(add1, "VMS"))      // "WNT"
	fmt.Println(strings.Map(add1, "Admix"))    // "Benjy"
}

func main() {
	// for _, url := range os.Args[1:] {
	// 	doc, err := reqAndParse(url)
	// 	if err != nil {
	// 		fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
	// 		continue
	// 	}
	// 	forEachNode(doc, startElement, endElement)
	// }
	// ex()
	fmt.Println(expand("His name is $fname $lname", replace))
}

func reqAndParse(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url,
			resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return doc, nil
}

// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
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

// For  example,  the  functions  startElement  and endElement
// print the start and end tags of an HTML element like <b>...</b>:
func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}
