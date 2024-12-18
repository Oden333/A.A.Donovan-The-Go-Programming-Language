package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func ex() {
	// ex5_15()
	// ex5_16()
	ex5_17()
}

func ex5_15() {
	a, _ := max([]int{1, 21, 35, 523, 45}...)
	fmt.Println(a)
	a, _ = min(1, 21, 35, 523, 45)
	fmt.Println(a)
}

func max(nums ...int) (int, error) {
	if len(nums) == 0 {
		return 0, fmt.Errorf("empty")
	}
	var max = nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	return max, nil
}
func min(nums ...int) (int, error) {
	if len(nums) == 0 {
		return 0, fmt.Errorf("empty")
	}
	var max = nums[0]
	for _, num := range nums {
		if num < max {
			max = num
		}
	}

	return max, nil
}

func ex5_16() {
	a := join([]string{"adsda", " ", "dasda"}...)
	fmt.Println(a)

}

func join(strs ...string) string {
	// return strings.Join(strs, "")
	res := new(strings.Builder)
	for _, str := range strs {
		res.WriteString(str)
	}
	return res.String()
}

func ex5_17() {
	for _, url := range []string{"http://golang.org"} {
		doc, err := get(url)
		if err != nil {
			log.Fatal(err)
		}
		images := ElementsByTagName(doc, "img")
		for i, img := range images {
			fmt.Printf("\nImage #%d ", i)
			for _, att := range img.Attr {
				fmt.Printf("\n%s\n", att.Val)
			}
		}
		headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
		for i, hdg := range headings {
			fmt.Printf("\nHeading #%d ", i)
			for _, att := range hdg.Attr {
				fmt.Printf("\n%s\n", att.Val)
			}
		}
	}

}

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	if doc == nil {
		return nil
	}
	nodes := make([]*html.Node, 0)
	for _, tag := range name {
		if strings.Contains(doc.Data, tag) {
			nodes = append(nodes, doc)
			break
		}
	}
	nodes = append(nodes, ElementsByTagName(doc.FirstChild, name...)...)
	return append(nodes, ElementsByTagName(doc.NextSibling, name...)...)
}

func get(url string) (*html.Node, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}
	return doc, nil
}
