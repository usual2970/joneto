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
    <style>
      body{padding-top:100px;}
    </style>
    <script src="{{.siteurl}}/static/js/jquery-1.8.3.min.js"></script>
   </head>

  <body>
    <div class="navbar navbar-inverse navbar-fixed-top">
      <div class="navbar-inner">
        <div class="container-fluid">
          <a href="{{.siteurl}}/amd/artlist" class="brand">囧途网--后台</a>

        </div>

        <div class="nav-collapse collapse">
          <p class="navbar-text pull-right">
            欢迎您： <a href="#" class="navbar-link">{{.email}}</a>
          </p>
          <ul class="nav">
            <li class="active"><a href="#">内容管理</a></li>
            <li><a href="#about">About</a></li>
            <li><a href="#contact">Contact</a></li>
          </ul>
        </div>
      </div>

    </div><!--头部-->
    <div class="container-fluid">
      <div class="row-fluid">
        <div class="span3">
          <div class="well sidebar-nav">
            <ul class="nav nav-list">
              <li class="nav-header">内容管理</li>
              <li class="active"><a href="#">测试</a></li>
              <li><a href="#">测试</a></li>
              <li><a href="#">测试</a></li>
            </ul>
          </div><!--well-->
        </div><!--/span-->
        <div class="span9">
          <div class="well">
            
          </div><!--well-->
        </div><!--/span-->
      </div>
    </div>

    <script src="{{.siteurl}}/static/js/bootstrap.min.js"></script>
  </body>
</html>
