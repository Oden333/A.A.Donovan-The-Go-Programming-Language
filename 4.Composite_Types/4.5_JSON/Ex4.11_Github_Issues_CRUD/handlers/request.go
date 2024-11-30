package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const BaseURL = "/%s/%s/issues"

func Requester(url, method, token string, body io.Reader) error {

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	cli := http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		return err
	}
	fmt.Println("Status code:", resp.StatusCode)
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNoContent {
		return nil
	}

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var respText interface{}
	if err := json.Unmarshal(respBytes, &respText); err != nil {
		return fmt.Errorf("failed to parse response body: %w", err)
	}
	printResp, _ := json.MarshalIndent(respText, "", " ")

	fmt.Printf("Response body:\n%s\n", printResp)
	return nil
}
