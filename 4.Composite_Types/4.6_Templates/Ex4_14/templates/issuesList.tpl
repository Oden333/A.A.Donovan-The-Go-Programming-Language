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