package main

import (
	"ex4_14"
	"flag"
)

var (
	user *string
	repo *string
)

func init() {

	user = flag.String("user", "", "GitHub user")
	repo = flag.String("repo", "", "GitHub repository")
	flag.Parse()

	// if *user == "" || *repo == "" {
	// 	log.Fatalf("You must pass user and repo")
	// 	return
	// }
}
func main() {
	ex4_14.Process(*user, *repo)
}
