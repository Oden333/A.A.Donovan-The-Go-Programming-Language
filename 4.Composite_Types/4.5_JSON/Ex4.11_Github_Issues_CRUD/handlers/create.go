package handlers

import (
	"GitHubCRUD/types"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func CreateIssue(token, user, repo string) {
	var (
		issue = &types.Issue{}
		ans   rune
		err   error
	)

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

	reqUrl := strings.Join([]string{types.ApiRepoPath, fmt.Sprintf(BaseURL, user, repo)}, "")
	err = Requester(reqUrl, http.MethodPost, token, bytes.NewReader(temp))
	if err != nil {
		panic(err)
	}
}

func createFromJSON(issue *types.Issue) {
	var (
		ans  rune
		path string
		err  error
		file *os.File
	)
	fmt.Println("Type in path to JSON file:")
	fmt.Scanf("%s\n", &path)

LOOP:
	file, err = os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			log.Println(err)
			fmt.Printf("The file with the following path does not exist:%s\nWanna try another y/n?", path)
		}
		log.Println("Error opening file:", err)
		fmt.Scanf("%c\n", &ans)
		if ans == 'y' {
			goto LOOP
		}
		return
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
	fileBytes := make([]byte, fileInfo.Size())
	_, err = file.Read(fileBytes)
	if err != nil {
		fmt.Println("Unable to read the file")
		log.Println(err)
		return
	}

	err = json.Unmarshal(fileBytes, issue)
	if err != nil {
		fmt.Println("unable to unmarshall the file", err)
		log.Println(err)
		return
	}
}

func createFromCMD(issue *types.Issue) {
	fmt.Println("not imptemented")
}
