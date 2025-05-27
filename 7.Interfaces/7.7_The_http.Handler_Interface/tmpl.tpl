
<!DOCTYPE html>
<html>
  <head>
    <title>Items Table</title>
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
	      <th>Item</th>
	      <th>Price</th>
      </tr>
      {{range $k, $v := .}}
        <tr>
            <td>{{$k}}</td>
            <td>{{$v}}</td>
        </tr>
      {{end}}
    </table>
  </body>
</html>