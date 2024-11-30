package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func Ex4_10_SearchWithAges() {
	var terms []string

	createdFrom := fmt.Sprintf("created:>%s", time.Now().AddDate(-1, 0, 0).Format(time.DateOnly))
	terms = append(terms, createdFrom)
	terms = append(terms, os.Args[1:]...)
	result, err := SearchIssues(terms)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("There are %d issues since %s:\n", result.TotalCount, time.Now().AddDate(-1, 0, 0).Format(time.DateOnly))
	for _, item := range result.Items {
		fmt.Printf("#%-5d\t%s\t%9.9s\t%.55s\n",
			item.Number, item.CreatedAt.Format(time.DateOnly), item.User.Login, item.Title)
	}
}
