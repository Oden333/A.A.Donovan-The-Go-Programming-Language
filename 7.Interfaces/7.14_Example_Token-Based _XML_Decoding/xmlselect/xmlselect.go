// Package Xmlselect prints the text of selected elements of an XML document.
package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

// The xmlselect program extracts and prints the text found beneath certain
// elements in an XML document tree. Using the API above, it can do its job in a single
// pass over the input without ever materializing the  tree
func ex7_17() {
	file, err := os.OpenFile("test", 2, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	dec := xml.NewDecoder(file)

	var stack []string        // stack of element names
	var els = []string{"div"} // = os.Args[1:]

	var tok xml.Token
	for {
		tok, err = dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok.Name.Local) // push onto a stack
			if m := attrsMapping(tok.Attr, els); len(m) > 0 {
				fmt.Printf("%s: %s\n----------------------\n", strings.Join(stack, " "), m)
			}
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop the name from the stack.
		case xml.CharData:
			if containsAll(stack, els) {
				fmt.Printf("%s: %s\n----------------------\n", strings.Join(stack, " "), tok)
			}

		}
	}
}

// The API  guarantees  that  the  sequence  of  StartElement and EndElement  tokens
// will  be  properly  matched,  even  in  ill-formed  documents.  Comments  are  ignored

// containsAll reports whether x contains the elements of y, in order.
func containsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}

func attrsMapping(attrs []xml.Attr, els []string) []string {
	res := make([]string, 0)
	for _, a := range attrs {
		if slices.Contains(els, a.Name.Local) {
			res = append(res, a.Value)
		}
	}
	return res
}

func parse(r io.Reader) (Node, error) {
	dec := xml.NewDecoder(r)
	// fmt.Println(io.ReadAll(r))
	var stack []*Element
	var root Node
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			el := &Element{tok.Name, tok.Attr, nil}
			if len(stack) == 0 {
				root = el
			} else {
				parent := stack[len(stack)-1]
				parent.Children = append(parent.Children, el)
			}
			stack = append(stack, el) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if len(bytes.TrimSpace(tok)) == 0 {
				continue
			}
			parent := stack[len(stack)-1]
			parent.Children = append(parent.Children, CharData(tok))
		}
	}
	return root, nil
}

type Node interface{} // CharData or *Element
type CharData string
type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func (n *Element) String() string {
	b := &bytes.Buffer{}
	visit(n, b, 0)
	return b.String()
}

func visit(n Node, w io.Writer, depth int) {
	switch n := n.(type) {
	case *Element:
		fmt.Fprintf(w, "%*s%s %s\n", depth*2, "", n.Type.Local, n.Attr)
		for _, c := range n.Children {
			visit(c, w, depth+1)
		}
	case CharData:
		fmt.Fprintf(w, "%*s%q\n", depth*2, "", n)
	default:
		panic(fmt.Sprintf("got %T", n))
	}
}

func main() {
	file, err := os.OpenFile("test", 2, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	node, err := parse(file)
	if err != nil {
		panic(err)
	}
	fmt.Println(node)
}
