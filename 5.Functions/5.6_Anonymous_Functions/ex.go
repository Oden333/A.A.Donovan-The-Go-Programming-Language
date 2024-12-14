package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func ex() {
	// ex5_10()
	// ex5_13()
	ex5_14()
}

func ex5_10() {
	for i, course := range topoSort1(prereqss) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

var prereqss = map[string][]string{
	"algorithms":            {"data structures"},
	"calculus":              {"linear algebra"},
	"compilers":             {"data structures", "formal languages", "computer organization"},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
	// ex5_11
	"linear algebra": {"calculus"},
}

func topoSort1(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	tempMark := make(map[string]bool)

	var visit func(items string)

	visit = func(item string) {
		if !seen[item] {
			if tempMark[item] {
				fmt.Println("Cycle warn. Key-", item)
				return
			}
			tempMark[item] = true
			for _, str := range m[item] {
				visit(str)
			}
			tempMark[item] = false
			seen[item] = true
			order = append(order, item)
		}

	}

	for key := range m {
		visit(key)
	}
	return order
}

func ex5_13() {
	breadthFirst(crawl1, os.Args[1:])
	// breadthFirst(crawl1, []string{"https://golang.org"})
}

func crawl1(url string) []string {
	fmt.Println(url)
	links, err := Extract(url)
	if err != nil {
		log.Print(err)
	}
	// if saved < 2 { // Debug
	if err = savePages(links); err != nil {
		log.Println(err)
	}
	// 	saved++
	// }
	return links
}

func savePages(urls []string) error {
	if len(urls) == 0 {
		return fmt.Errorf("empty urls to save")
	}
	domain, err := url.Parse(urls[0])
	if err != nil {
		return err
	}
	err = os.Mkdir(domain.Host, os.ModeDir)
	if err != nil {
		if !os.IsExist(err) {
			return err
		}
	}
	visits := make(map[string]struct{})
	for _, link := range urls[1:] {
		url, err := url.Parse(link)
		if err != nil {
			continue
		}
		if url.Host != domain.Host {
			continue
		}
		if _, ok := visits[link]; ok {
			log.Println("duplicate link", link)
			continue
		}
		visits[link] = struct{}{} // Additional deduplication
		err = savePage(url)
		if err != nil {
			log.Println(err)
		}
	}

	return nil
}

func savePage(urla *url.URL) error {
	path := strings.Join([]string{urla.Host, urla.Path}, "")
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, urla.String(), nil)
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	fmt.Println("Saving path:", path)
	err = os.MkdirAll(path, os.ModeDir)
	if err != nil {
		if !os.IsExist(err) {
			return err
		}
	}
	page, err := os.Create(strings.Join([]string{path, "/", "page.html"}, ""))
	if err != nil {
		return err
	}
	defer page.Close()

	if _, err = page.ReadFrom(resp.Body); err != nil {
		return err
	}

	return nil
}

func ex5_14() {
	for corse, prereq := range prereqs {
		fmt.Fprintf(os.Stdout, "==========\n%q depends on:\n", corse)
		breadthFirst(deps, prereq)
		fmt.Fprintf(os.Stdout, "==========\n\n")
	}
}

func deps(corse string) (order []string) {
	fmt.Fprintf(os.Stdout, " - %s\n", corse)
	return append(order, prereqs[corse]...)
}
