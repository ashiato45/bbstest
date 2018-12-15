<html>
<head>
<title></title>
</head>
<body>
<form action="/" method="post">
    Username: <input type="text" name="username">
    Message: <input type="password" name="password">
    <input type="hidden" name="token" value="{{.}}">
    <input type="submit" value="Write">
</form>
</body>
</html>