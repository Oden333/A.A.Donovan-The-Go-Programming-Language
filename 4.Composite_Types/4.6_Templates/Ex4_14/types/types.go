package types

import "time"

const (
	ApiRepoPath = "https://api.github.com/repos"
)

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Id        int       `json:"id,omitempty"`
	Title     string    `json:"title"`
	Body      string    `json:"body,omitempty"`
	Milestone int       `json:"milestone,omitempty"`
	Labels    []string  `json:"labels,omitempty"`
	Assignees []string  `json:"assignees,omitempty"`
	HTMLURL   string    `json:"html_url,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	State     string    `json:"state,omitempty"`
	User      *User
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
