package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func ex() {
	// ex5_1()
	// ex5_2()
	// ex5_3()
	ex5_4()
}

func ex5_1() {
	in, err := os.Open("gorg.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer in.Close()

	// q(in)

	in = os.Stdin
	doc, err := html.Parse(in)
	// doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatalf("findlinks1: %v\n", err)
	}
	for _, link := range findLinks(doc, nil) {
		fmt.Println(link)
	}
}

func q(in *os.File) {
	content, err := io.ReadAll(in)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("File content:")
	fmt.Println(string(content[:1000]))
	in.Seek(0, 0)
}

func findLinks(n *html.Node, links []string) []string {
	if n == nil {
		return links
	}
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	links = findLinks(n.FirstChild, links)
	return findLinks(n.NextSibling, links)
}

//
//
//
//
//
//
//

func ex5_2() {
	in, err := os.Open("gorg.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer in.Close()

	// q(in)
	doc, err := html.Parse(in)
	// doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatalf("findlinks1: %v\n", err)
	}
	elems := make(map[string]map[string]int)

	var populate func(n *html.Node)
	populate = func(n *html.Node) {
		if n == nil {
			return
		}
		if n.Type == html.ElementNode {
			for _, a := range n.Attr {
				if elems[n.Data] == nil {
					elems[n.Data] = make(map[string]int)
					elems[n.Data][a.Key]++
					continue
				}
				elems[n.Data][a.Namespace]++
			}
		}
		populate(n.FirstChild)
		populate(n.NextSibling)
		return
	}
	populate(doc)
	for k, v := range elems {
		fmt.Printf("%s:\n", k) // Выводим ключ верхнего уровня
		for key, value := range v {
			fmt.Printf("%s: %d\n", key, value) // Выводим вложенные пары ключ-значение
		}
		fmt.Println()
	}
}

func ex5_3() {
	in, err := os.Open("gorg.html")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer in.Close()

	// q(in)

	// in = os.Stdin
	doc, err := html.Parse(in)
	if err != nil {
		log.Fatalf("findlinks1: %v\n", err)
	}
	for _, link := range findText(doc, nil) {
		if len(strings.TrimSpace(link)) != 0 {
			fmt.Fprintln(os.Stdout, link)
		}
	}
}

func findText(node *html.Node, text []string) []string {
	if node == nil {
		return text
	}
	if node.Type == html.TextNode {
		if node.Parent.Data != "script" && node.Parent.Data != "style" {
			for _, line := range strings.Split(node.Data, "\n") {
				if len(line) != 0 {
					text = append(text, line)
				}
			}
		}
	}
	text = findText(node.FirstChild, text)
	return findText(node.NextSibling, text)
}

func ex5_4() {
	in, err := os.Open("gorg.html")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer in.Close()

	// q(in)

	// in = os.Stdin
	doc, err := html.Parse(in)
	if err != nil {
		log.Fatalf("findlinks1: %v\n", err)
	}
	for _, link := range findLinksExtended(doc, nil) {
		if len(strings.TrimSpace(link)) != 0 {
			fmt.Fprintln(os.Stdout, link)
		}
	}
}

func findLinksExtended(node *html.Node, text []string) []string {
	if node == nil {
		return text
	}
	if node.Type == html.ElementNode {
		for _, att := range node.Attr {
			if strings.Contains(att.Val, "http") {
				text = append(text, att.Val)
			}
		}
		// Или так, дешевле по ресурсам
		//
		// switch n.Data {
		// case "a", "link":
		// 	for _, a := range n.Attr {
		// 		if a.Key == "href" {
		// 			links = append(links, a.Val)
		// 		}
		// 	}
		// case "img", "script":
		// 	for _, a := range n.Attr {
		// 		if a.Key == "src" {
		// 			links = append(links, a.Val)
		// 		}
		// 	}
		// }
	}
	text = findLinksExtended(node.FirstChild, text)
	return findLinksExtended(node.NextSibling, text)
}
