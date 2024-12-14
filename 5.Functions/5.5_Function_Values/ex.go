package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func ex() {
	// for _, url := range os.Args[1:] {

	doc, err := reqAndParse("http://gopl.io")
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
	}

	// file, err := os.Open("out.html")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()
	// doc, err := html.Parse(file)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// ex5_7(doc, prettyStart, prettyEnd)
	ex5_8(doc, os.Args[1], startOk, endOk)
	// }

}

func ex5_7(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for node := n.FirstChild; node != nil; node = node.NextSibling {
		ex5_7(node, pre, post)
	}
	if post != nil {
		post(n)
	}
}

var depth int

func prettyStart(n *html.Node) {
	switch n.Type {
	case html.CommentNode:
		fmt.Printf("<!--%s-->\n", n.Data)
	case html.TextNode:
		text := strings.TrimSpace(n.Data)
		if len(text) == 0 {
			return
		}
		fmt.Printf("%*s%s\n", depth*2, "", text)
	case html.ElementNode:
		end := ">"
		if n.FirstChild == nil {
			end = "/>"
		}
		fmt.Printf("%*s<%s%s%s\n", depth*2, "", n.Data, join(n.Attr), end)
		depth++
	}
}
func prettyEnd(n *html.Node) {
	switch n.Type {
	case html.CommentNode, html.TextNode:

	case html.ElementNode:
		depth--
		if n.FirstChild == nil {
			return
		}
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

func join(att []html.Attribute) string {
	if len(att) == 0 {
		return ""
	}
	var res string
	for _, a := range att {
		res += fmt.Sprintf(" %s='%s'", a.Key, a.Val)
	}
	return res
}

func ElementByID(doc *html.Node, id string) *html.Node {
	return ex5_8(doc, id, startOk, endOk)
}
func ex5_8(n *html.Node, id string, pre, post func(n *html.Node, id string) bool) *html.Node {
	if pre != nil {
		if pre(n, id) {
			return n
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if ex5_8(c, id, pre, post) != nil {
			return c
		}
	}
	if post != nil {
		if post(n, id) {
			return n
		}
	}
	return nil
}

// For  example,  the  functions  startElement  and endElement
// print the start and end tags of an HTML element like <b>...</b>:
func startOk(n *html.Node, id string) bool {
	for _, att := range n.Attr {
		if att.Key == id {
			fmt.Println("found", att.Val)
			return true
		}
	}
	return false
}

func endOk(n *html.Node, id string) bool {
	for _, att := range n.Attr {
		if att.Key == id {
			fmt.Println("found", att.Val)
			return true
		}
	}
	return false
}

// func ex5_9(expand(...)){}
func expand(s string, f func(string) string) string {
	strs := strings.Split(s, " ")
	for i, str := range strs {
		if strings.HasPrefix(str, "$") {
			strs[i] = f(strings.TrimPrefix(str, "$"))
		}
	}
	return strings.Join(strs, " ")
}

func replace(s string) string {
	switch s {
	case "fname":
		return "John"
	case "lname":
		return "Wick"
	}
	return ""
}
