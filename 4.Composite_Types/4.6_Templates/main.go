package main

import (
	"github_searcher"
	html_template "html/template"
	"log"
	"os"
	"text/template"
	"time"
)

//! A  template  is  a  string  or  file  containing  one  or  more  portions
//! enclosed  in  double braces,  {{...}},  called  actions.

// Each  action  contains  an  expression  in  the  template language,
// a  simple  but  powerful  notation  for  printing  values,  selecting  struct  fields,
// calling functions and methods, expressing control flow such as if-else statements
// and  range  loops,  and  instantiating  other  templates.

// Within an action, there is a notion of the current value, referred to as “dot” and written as “.”, a period.
// & The dot initially refers to the template’s  parameter,
// which  will  be  a  github.IssuesSearchResult  in  this example.

//? The  {{.TotalCount}}  action  expands  to  the  value  of  the TotalCount  field,  printed  in  the  usual  way
//? The  {{range .Items}}  and {{end}} actions create a loop, text between  expanded multiple times
//? The | notation makes the result of one operation the argument of another
//? the printf function, is a built-in synonym for fmt.Sprintf in all templates

const templ = `{{.TotalCount}} issues:
{{range .Items}}--------------------------------------
--
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

// In the same way that a type may control its string formatting by defining certain methods,
// a type may also define methods to control its JSON marshaling and unmarshaling behavior.

func textTemplate() {
	issues, _ := github_searcher.SearchIssues(os.Args[1:])
	// Producing  output  with  a  template  is  a  two-step  process:
	// First  we  must  parse  the template into a suitable internal representation,
	report, err := template.New("report").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(templ)
	if err != nil {
		log.Fatal(err)
	}
	// then execute it on specific inputs.
	report.Execute(os.Stdout, issues)
}

//& Because  templates  are  usually  fixed  at  compile  time,
//& failure  to  parse  a  template indicates a fatal bug in the program.

// ? The template.Must helper function makes error handling more convenient:
// ? it accepts a template and an error, checks that the error is nil (and panics otherwise),
// ? and then returns the template.
var report = template.Must(template.New("issuelist").Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(templ))

func htmlTemplate() {
	// It  uses  the  same  API  and expression  language  as  text/template
	// but  adds  features  for  automatic  and context-appropriate
	// escaping of strings appearing within HTML, JavaScript, CSS, or URLs.

	// These  features  can  help  avoid  a  perennial  security  problem  of  HTML
	// generation, an injection  attack, in which an adversary crafts a string value like the
	// title  of  an  issue  to  include  malicious  code  that,  when  improperly  escaped  by  a
	// template, gives them control over the page
	var issueList = html_template.Must(html_template.New("issuelist").Parse(`
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
	  <td><a href='{{.HTMLURL}}'>{{.Number}}</td>
	  <td>{{.State}}</td>
	  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a>
	</td>
	  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
	</tr>
	{{end}}
	</table>
	`))
	issues, err := github_searcher.SearchIssues(os.Args[1:])
	if err != nil {
		panic(err)
	}
	issueList.Execute(os.Stdout, issues)

	//& The html/template package automatically HTML-escapes the titles so that they appear literally.

	// For example, using wrong package can lead to
	//
	//? text/template -> "&lt;"  would have been rendered as a less-than character '<',
	//? text/template -> "<link>" would have become a link element,
	// changing the structure of the HTML document and perhaps compromising its security
}

// We  can  suppress  this  auto-escaping  behavior  for  fields  that  contain  trusted  HTML
// data by using the named string type template.HTML instead of string.
func differentTmplPkg() {
	const templ = `<p>A: {{.A}}</p><p>B: {{.B}}</p>`
	t :=
		html_template.Must(html_template.New("escape").Parse(templ))
	var data struct {
		A string             // untrusted plain text
		B html_template.HTML // trusted HTML
	}
	data.A = "<b>Hello!</b>"
	data.B = "<b>Hello!</b>"
	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// textTemplate()
	// htmlTemplate()
	// differentTmplPkg()
}
