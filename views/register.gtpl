<!DOCTYPE html>
<html>
<head>
<title></title>
<script type="text/javascript" src="https://ajax.googleapis.com/ajax/libs/jquery/1.4.4/jquery.min.js"></script>
<script type="text/javascript" src="https://raw.github.com/GRINPublishing/GTPL/master/lib/gtpl.js"></script>
</head>
<body>

<form name="register" method="post">
  <label for="Name">Name:</label>
  <input type="text" id="Name" name="Name"><br><br>
  <label for="Password">Password:</label>
  <input type="text" id="Password" name="Password"><br><br>
  <label for="Age">Age:</label>
  <input type="text" id="Age" name="Age"><br><br>
  <label for="Email">Email:</label>
  <input type="text" id="Email" name="Email"><br><br>
  <input type="submit" id="submit-btn" value="Submit">
</form>
<label for="Status" id="Status"></label>
</body>
<script>
    $(document).ready(function(){
        $("#submit-btn").click(function(e)
        {
            $.ajax({
                url : "http://192.168.56.105:8080/api/user",
                type: "POST",
                data : JSON.stringify({ "Name": $('#Name').val(), "Age":  parseInt($('#Age').val()), "Email": $('#Email').val(), 
                      "Password":$('#Password').val(), Validation:false}),
                success: function(data){
                    console.log('AJAX SUCCESS, data : '+ data.status); 
                    if (data.status == 'Failed')
                        $("#Status").text('Failed');
                    else
                        window.location.href = "http://192.168.56.105:8080/index";
                },
                error: function(errMsg){ 
                    console.log('AJAX FAILED, message : '+errMsg);
                }
            });
            e.preventDefault();
        });
    });
</script>
</html>