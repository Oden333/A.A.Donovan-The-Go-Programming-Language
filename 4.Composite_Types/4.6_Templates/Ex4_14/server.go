package ex4_14

import (
	"ex4_14/requester"
	"ex4_14/types"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var issuesList = template.Must(template.New("issuesList").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Id}}</td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a>
</td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

func Process(user, repo string) {
	issues, err := requester.ReadIssues(user, repo)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%+d", len(issues))
	// var issuesList = template.Must(template.New("issuesList").ParseFiles(`../templates/issuesList.tpl`))
	var tmplIssues = types.IssuesSearchResult{
		TotalCount: len(issues),
		Items:      issues,
	}
	// err = issuesList.Execute(os.Stdout, tmplIssues)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	http.HandleFunc("/issues", func(w http.ResponseWriter, r *http.Request) {
		if err := issuesList.Execute(w, tmplIssues); err != nil {
			log.Fatal(err)
		}
	})
	fmt.Println("Server Started")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
