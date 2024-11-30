package main

import (
	"GitHubCRUD/handlers"
	"GitHubCRUD/types"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
)

var (
	token *string
)

// Exercise 4.11: Build  a  tool  that  lets  users  create,  read,  update,  and  delete  GitHub
// issues from the command line, invoking their preferred text editor when substantial
// text input is required.

func main() {

	// Log file init
	var logFile *os.File
	logFile, err := os.OpenFile("./logs/.log", 2, 0666)
	if err != nil {
		logFile, err = os.Create("./logs/.log")
		if err != nil {
			fmt.Fprintf(os.Stdout, "Unable to init logger")
		}
	}
	defer logFile.Close()
	log.SetOutput(logFile)
	fmt.Fprintln(os.Stdout, "Logger init success")

	token = flag.String("token", "", "GitHub personal access token")
	flag.Parse()
	if *token == "" {
		fmt.Println("Token cannot be empty string")
		log.Fatalln("Token cannot be empty string")
	}

	var repo string
	fmt.Fprintf(os.Stdout, "Type repository to iterate with:\n")
	fmt.Scanf("%s\n", &repo)

	var user string
	fmt.Fprintf(os.Stdout, "And repo Owner:\n")
	fmt.Scanf("%s\n", &user)

	var (
		action string
		ans    rune
	)
Zalupa:
	{
		fmt.Fprintf(os.Stdout, "Type method to iterate with issues of %s repo:\n", path.Join(user, repo))
		fmt.Scanf("%s\n", &action)
		switch types.Action(action) {
		case types.Create:
			handlers.CreateIssue(*token, user, repo)
		case types.Read:
			handlers.ReadIssue(*token, user, repo)
		case types.Update:
			handlers.UpdateIssue(*token, user, repo)
		case types.Delete:
			handlers.DeleteIssue(*token, user, repo)
		default:
			fmt.Println("Unknown method, try one of those (create, read, update, delete(lock))")
			goto Zalupa
		}
		fmt.Println("Do you want more? y/n")
		fmt.Scanf("%c\n", &ans)
		if ans == 'y' {
			goto Zalupa
		}
	}

}
