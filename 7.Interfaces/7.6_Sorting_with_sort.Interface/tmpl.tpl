
<!DOCTYPE html>
<html>
  <head>
    <title>Music Table</title>
      <style>
        table {	      
          border-collapse: collapse;        
        }
        td, th {
	      border: solid 1px;
	      padding: 0.5em;
          text-align: right;
        }
      </style>
  </head>
  <body>
    <table>
      <tr>
	      <th><a href="./?by=title">Title</a></th>
	      <th><a href="./?by=artist">Artist</a></th>
	      <th><a href="./?by=album">Album</a></th>
	      <th><a href="./?by=year">Year</a></th>
	      <th><a href="./?by=length">Length</a></th>
	    </tr>
      {{range .}}
      <tr>
        <td>{{.Title}}</td>
        <td>{{.Artist}}</td>
        <td>{{.Album}}</td>
        <td>{{.Year}}</td>
        <td>{{.Length}}</td>
      </tr>
      {{end}}
    </table>
  </body>
</html>