package handlers

import (
	"GitHubCRUD/types"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func ReadIssue(token, user, repo string) {
	var (
		reqId  string
		reqUrl string
		err    error
	)
Num:
	fmt.Println("Type specific Issue issue_number, or just enter to see all issues:")
	fmt.Scanf("%s\n", &reqId)
	if _, err = strconv.Atoi(reqId); err != nil {
		if reqId != "" {
			fmt.Println("Invalid number, it should be a string")
			goto Num
		}
	}
	reqUrl = strings.Join([]string{types.ApiRepoPath, fmt.Sprintf(BaseURL, user, repo)}, "")
	if reqId != "" {
		reqUrl = strings.Join([]string{reqUrl, "/", reqId}, "")
	}
	err = Requester(reqUrl, http.MethodGet, token, nil)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
	}
}
