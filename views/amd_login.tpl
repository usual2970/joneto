<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="description" content="">
    <meta name="author" content="">
    <link rel="shortcut icon" href="/static/img/logo_black.png">
    <title>囧途网——后台管理</title>
    <link href="{{.siteurl}}/static/css/bootstrap.min.css" rel="stylesheet">
    <script src="{{.siteurl}}/static/js/jquery-1.8.3.min.js"></script>
    <style>
      body{
        padding-top: 40px;
        padding-bottom: 40px;
        background-color: #f5f5f5;
      }
      .form-signin{max-width: 300px;
padding: 19px 29px 29px;
margin: 0 auto 20px;
background-color: #fff;
border: 1px solid #e5e5e5;
-webkit-border-radius: 5px;
-moz-border-radius: 5px;
border-radius: 5px;
-webkit-box-shadow: 0 1px 2px rgba(0,0,0,.05);
-moz-box-shadow: 0 1px 2px rgba(0,0,0,.05);
box-shadow: 0 1px 2px rgba(0,0,0,.05);}

    </style>
  </head>

  <body>

    <div class="container">

      <form class="form-signin" method="post" action="{{.siteurl}}/amd/login">
        <h3 class="form-signin-heading">管理员登录</h2>
        <input type="text" class="form-control" placeholder="Email address" autofocus name="email" value="{{.email}}">
        <input type="password" class="form-control" placeholder="Password" name="password" value="{{.password}}">
        {{if .err}}
        <p class="text-error">{{.err}}</p>
        {{end}}
        <label class="checkbox">
          <input type="checkbox" value="remember-me"> 记住我
        </label>
        <button class="btn btn-lg btn-primary" type="submit">登&nbsp;&nbsp;&nbsp;录</button>
      </form>

    </div> 
    <script src="{{.siteurl}}/static/js/bootstrap.min.js"></script>
  </body>
</html>
