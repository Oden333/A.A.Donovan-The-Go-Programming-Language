package requester

import (
	"encoding/json"
	"ex4_14/types"
	"fmt"
	"net/http"
	"strings"
)

const pathURL = "/%s/%s/issues"

func ReadIssues(user, repo string) ([]*types.Issue, error) {
	reqUrl := strings.Join([]string{types.ApiRepoPath, fmt.Sprintf(pathURL, user, repo)}, "")
	resp, err := http.Get(reqUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("unable to proceed request", resp.StatusCode)
		return nil, err
	}
	var issues []*types.Issue
	err = json.NewDecoder(resp.Body).Decode(&issues)
	if err != nil {
		return nil, err
	}

	return issues, nil
}
