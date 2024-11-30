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

func DeleteIssue(token, user, repo string) {
	var (
		reqUrl string
		reqId  string
		err    error
		reason string
	)

Num:
	fmt.Println("Type specific Issue number")
	fmt.Scanf("%s\n", &reqId)
	if _, err = strconv.Atoi(reqId); err != nil {
		fmt.Println("Invalid number")
		goto Num
	}
Res:
	fmt.Println("Type closing reason")
	fmt.Scanf("%s\n", &reason)
	switch reason {
	default:
		fmt.Println(
			`The reason for locking the issue or pull request conversation.	
Lock will fail if you don't use one of these reasons:
			off-topic
			too heated
			resolved
			spam`,
		)
		goto Res
	case "off-topic", "too heated", "resolved", "spam":
	}

	reqUrl = strings.Join([]string{types.ApiRepoPath, fmt.Sprintf(BaseURL, user, repo), "/", reqId, "/lock"}, "")
	body, err := json.Marshal(map[string]string{"lock_reason": reason})
	if err != nil {
		fmt.Println("unable to marshall reason")
		return
	}
	fmt.Printf("%s\n", body)
	err = Requester(reqUrl, http.MethodPut, token, bytes.NewReader(body))
	if err != nil {
		fmt.Println(err)
		log.Println(err)
	}
}
