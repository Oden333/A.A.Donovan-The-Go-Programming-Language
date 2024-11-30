package handlers

import (
	"GitHubCRUD/types"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func UpdateIssue(token, user, repo string) {
	var (
		issue = &types.Issue{}
		ans   rune
		reqId string
		err   error
	)
Num:
	fmt.Println("Type Issue number")
	fmt.Scanf("%s\n", &reqId)
	if issue.Id, err = strconv.Atoi(reqId); err != nil {
		if reqId != "" {
			fmt.Println("Invalid number")
			goto Num
		}
	}

	fmt.Println("Do you want to load JSON struct from file? y/n")
	fmt.Scanf("%c\n", &ans)
	switch ans {
	case 'y':
		createFromJSON(issue)
	case 'n':
		createFromCMD(issue)
	}

	log.Printf("Got issue to create %+v", &issue)
	temp, _ := json.MarshalIndent(&issue, " ", "  ")
	fmt.Printf("Scanned issue: %s\n", string(temp))

	reqUrl := strings.Join([]string{types.ApiRepoPath, fmt.Sprintf(BaseURL, user, repo), "/", strconv.Itoa(issue.Id)}, "")
	err = Requester(reqUrl, http.MethodPatch, token, bytes.NewReader(temp))
	if err != nil {
		panic(err)
	}
}
