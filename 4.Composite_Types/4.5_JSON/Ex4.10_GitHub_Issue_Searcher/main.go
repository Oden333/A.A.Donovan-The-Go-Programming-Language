// Package github provides a Go API for the GitHub issue tracker.
// See https://developer.github.com/v3/search/#search-issues.
package main

import (
	"time"
)

// Exercise 4.10: Modify issues to report the results in age categories, say less than a
// month old, less than a year old, and more than a year old

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func main() {
	// result, err := SearchIssues(os.Args[1:])
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%d issues:\n", result.TotalCount)
	// for _, item := range result.Items {
	// 	fmt.Printf("#%-5d %9.9s %.55s\n",
	// 		item.Number, item.User.Login, item.Title)
	// }

	Ex4_10_SearchWithAges()

}
