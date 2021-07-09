  
<html>
<head>
<title></title>
<script type="text/javascript" src="https://ajax.googleapis.com/ajax/libs/jquery/1.4.4/jquery.min.js"></script>
<script type="text/javascript" src="https://raw.github.com/GRINPublishing/GTPL/master/lib/gtpl.js"></script>
<link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/css/bootstrap.min.css"> <!-- load bulma css -->
<link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css"> <!-- load fontawesome -->
</head>
<body>
<form name="login" method="post" action="/login">
    使用者名稱:<input type="text" id="name" name="Username">
    密碼:<input type="password" id="password" name="Password">
    <input type="submit" id="submit-btn" name="submit" value="登入">
</form>
<form action="/register" method="get">
    <input type="submit" value="regist account" />
    <br></br>
    <a href="/auth/google" class="btn btn-danger"><span class="fa fa-google"></span> SignIn with Google</a>
</form>
</body>
</html>