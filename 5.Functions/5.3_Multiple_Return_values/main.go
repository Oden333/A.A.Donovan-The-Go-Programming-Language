package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	// findAndVisit()
	w, i, err := CountWordsAndImages("https://golang.org")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Words: %d, Images: %d", w, i)
}
func findAndVisit() {
	for _, url := range os.Args[1:] {
		links, err := findLinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks2: %v\n",
				err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}
}

// findLinks performs an HTTP GET request for url, parses the
// response as HTML, and extracts and returns the links.
func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	//!	We must ensure that resp.Body is closed so that network resources are properly
	//! released even in case of error.

	//& Go’s garbage collector recycles unused memory, but do
	//& not  assume  it  will  release  unused  operating  system  resources  like  open  files  and
	//& network connections. They should be closed explicitly
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
	return visit(nil, doc), nil
}

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

//? Well-chosen names can document the significance of a function’s results.
// Names are particularly valuable when a function returns multiple results of the same type, like
//* func Size(rect image.Rectangle) (width, height int)
//* func Split(path string) (dir, file string)
//* func HourMinSec(t time.Time) (hour, minute, second int)

// In a function with named results, the operands of a return statement may be omitted.
// This is called a bare return.

// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

// func ex5_5(){}
func countWordsAndImages(n *html.Node) (words, images int) {
	if n == nil {
		return
	}
	switch n.Type {
	case html.TextNode:
		if n.Parent.Data != "script" && n.Parent.Data != "style" {
			scanner := bufio.NewScanner(strings.NewReader(n.Data))
			scanner.Split(bufio.ScanWords)
			for scanner.Scan() {
				if scanner.Text() == "end" {
					break
				} else {
					words++
				}
			}
		}
	case html.ElementNode:
		if n.Data == "img" {
			images++
		}
	}

	wrd, img := countWordsAndImages(n.FirstChild)
	words += wrd
	images += img
	wrd, img = countWordsAndImages(n.NextSibling)
	words += wrd
	images += img
	return
}
