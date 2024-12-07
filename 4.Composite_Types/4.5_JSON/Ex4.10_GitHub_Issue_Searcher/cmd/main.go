// Package github provides a Go API for the GitHub issue tracker.
// See https://developer.github.com/v3/search/#search-issues.
package main

import (
	"github_searcher"
)

// Exercise 4.10: Modify issues to report the results in age categories, say less than a
// month old, less than a year old, and more than a year old

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

	github_searcher.Ex4_10_SearchWithAges()

}
