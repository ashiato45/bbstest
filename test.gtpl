<html>
<head>
<title></title>
</head>
<body>
<form action="/" method="post">
    Username: <input type="text" name="username">
    Message: <input type="text" name="message">
    <input type="hidden" name="token" value="{{.token}}">
    <input type="submit" value="Write">
</form>
<table border="1">
  <tr>
    <th>Username</th>
    <th>Message</th>
    <th>Date</th>
  </tr>
  {{range .rows}}
  <tr>
    <td>{{.username}}</td>
    <td>{{.message}}</td>
    <td>{{.date}}</td>    
  </tr>
  {{end}}

</table>
</body>
</html>
