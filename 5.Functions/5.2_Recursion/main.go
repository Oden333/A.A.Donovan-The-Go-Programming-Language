package main

//! Functions  may  be  recursive,  that  is,  they  may  call  themselves,
//! either  directly  or indirectly.

// The golang.org/x/... repositories hold packages designed and maintained by the Go team for applications
// These  packages are  not  in  the  standard  library  because  they’re  still  under  development  or  because
// they’re rarely needed by the majority of Go programmers.

// The  function  html.Parse  reads  a  sequence  of  bytes,  parses  them,  and
// returns the root of the HTML document tree, which is an html.Node.
// HTML has several kinds of nodes—text, comments, and so on—but here we are concerned
// only with element nodes of the form <name key='value'>.
import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func parseLinks() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

// The visit function traverses an HTML node tree, extracts the link from the href
// attribute of each anchor element <a href='...'>
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func main() {
	// doc, err := html.Parse(os.Stdin)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "outline: %v\n", err)
	// 	os.Exit(1)
	// }
	// outline(nil, doc)
	ex()
}

// When outline calls itself recursively, the callee receives a copy of  stack.
// Although  the  callee  may  append  elements  to  this  slice,  modifying  its
// underlying array and perhaps even allocating a new array, it doesn’t modify the initial
// elements  that  are  visible  to  the  caller,  so  when  the  function  returns,  the  caller’s
// stack is as it was before the call.
func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

// Many  programming  language  implementations  use  a  fixed-size  function  call  stack;
// sizes from 64KB to 2MB are typical.
// Fixed-size stacks impose a limit on the depth of
// recursion, so one must be careful to avoid a stack overflow when traversing large data
// structures  recursively;  fixed-size  stacks  may  even  pose  a  security  risk.

// In  contrast,
// typical  Go  implementations  use  variable-size  stacks  that  start  small  and  grow  as
// needed up to a limit on the order of a gigabyte.
