<!DOCTYPE html>
<html>
  	<head>
    	<title>囧途网-囧途中的囧图集，囧文集，旅途百科，驴友集散地</title>
    	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    	<link rel="Shortcut Icon" href="/static/img/logo_black.png" type="image/x-icon">
    	<link href="{{.siteurl}}/static/css/global.css" rel="stylesheet" type="text/css"></link>
      <link href="{{.siteurl}}/static/css/art.css" rel="stylesheet" type="text/css"></link>
      <link href="{{.siteurl}}/static/css/dialog.css" rel="stylesheet" type="text/css"></link>
      <script src="{{.siteurl}}/static/js/jquery-1.8.3.min.js"></script>
      <script src="{{.siteurl}}/static/js/jquery.lazyload.min.js"></script>
      <script src="{{.siteurl}}/static/js/global.js"></script>
  	</head>
  	
  	<body>
  	   <div class="jt_c_wrap">
          <div class="jt_c">
            <div class="jt_c_title">
              {{.content.Title}}
            </div>
            <div class="jt_c_body">
              <span class="jt_cp_info">
                <span class="jt_cp_wen">文</span>
                <a href="#" target="_blank">iphone.myzaker.com</a>
              </span>
              <pre class="prebody">
                {{str2html .content.Content}}
              </pre>
            </div>
            <div class="jt_cont_op">       
              <span class="jt_fdopt">         
                <a rel="16352704" href="javascript:void(0)" class="jt_fav"></a>
                <span class="jt_split">-</span>         
                <a class="jt_post_comm" href="javascript:void(0);">
                  <span>评论</span>           
                  <span class="jt_ml3 jt_reply_count">2</span>
                </a> 
                <span class="jt_cont_span"></span>
                <span class="jt_c_time">8月2日 下午2点33分</span>
              </span>     
            </div>
          </div>
       </div>

       <div class="jt_c_ext">
        {{if .next.Id}}
        <a class="next_c cnext" href="{{.siteurl}}/art/{{.next.Id}}" title="{{.next.Title}}">
          <span style="vertical-align:15px;">{{substr .next.Title 0 15}}...</span>
          <span class="arrow"></span>
        </a>
        {{end}}
        {{if .prev.Id}}
        <a class="pre_c cprev" href="{{.siteurl}}/art/{{.prev.Id}}" title="{{.prev.Title}}">
          <span class="arrow"></span>
         {{substr .prev.Title 0 15}}...
        </a>
        {{end}}
      </div>

      <script src="/static/js/juicer.js"></script>
      {{template "dialog.tpl"}}
      <script src="/static/js/dialog_comm.js"></script>
      <script>
        $(document).ready(function(){
          $(document).off("click",".jt_post_comm").on("click",".jt_post_comm",function(){
            $.dialog_comm();
          });
        });
      </script>
      </body>
</html>