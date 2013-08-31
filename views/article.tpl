{{template "header.tpl"}}
<link href="/static/css/dialog.css" rel="stylesheet" type="text/css"></link>    
<script src="static/js/global.js"></script>
  		<div class="jt_content">
        <div class="jt_left">
          <div class="jt_lm">
            <a href="" class="jt_lm_selected">
              <span>精选</span>
            </a>
          </div>

          <div class="jt_lm">
            <a href="">
              <span>工作</span>
            </a>
          </div>

          <div class="jt_lm">
            <a href="">
              <span>产品</span>
            </a>
          </div>

          <div class="jt_lm">
            <a href="">
              <span>创业</span>
            </a>
          </div>

          <div class="jt_lm">
            <a href="">
              <span>技术</span>
            </a>
          </div>

          <div class="jt_lm">
            <a href="">
              <span>码农爱旅行</span>
            </a>
          </div>
        </div>
        <div class="jt_main clearfix">
          <div class="jt_mc clearfix">





            <div class="jt_posts clearfix">
              <div class="jt_cont_li_wrap1"></div>
              <div class="jt_cont_li_wrap2"></div>
              <div class="jt_cont_li_wrap3"></div>
              

              

             


            </div>
          </div>
        </div>
  		</div>
<script id="list_tpl" type="text/template">
<div id="content_html">
{@each Arts as v,k}
  <div id="list_item$${k%3}" class="list_item clearfix" style="padding:0 24px 0 0;">
  <div class="jt_cont_li clearfix">
    <div class="jt_cont_wrap">
      <a href="/art/$${v.Id}" target="_blank" title="$${v.Title}">

        <img src="$${v.Images}" width="100%"/>
      </a>
    </div>

    <div class="jt_cont_desc">
      <a href="/art/$${v.Id}"  target="_blank" title="$${v.Title}">
        $${v.Desc}
      </a>
    </div>

    <div class="jt_cont_op">
      <span class="jt_fdopt">
        <a rel="16352704" href="javascript:void(0)" class="jt_fav"></a>

        <span class="jt_split">-</span>

        <a class="jt_post_comm" href="javascript:void(0);">
          <span id="jt_comm">评论</span>
          <span class="jt_ml3 jt_reply_count" id="jt_comm_num">2</span>
        </a>

        <span class="jt_cont_span"></span>
        <a class="jt_cont_all" href="/art/$${v.Id}" target="_blank" title="$${v.Title}">查看全文</a>
      </span>
    </div>
    
  </div>
  <div style="height:35px;"></div>
  </div>
{@/each}
</div>
</script>
<script src="/static/js/juicer.js"></script>
{{template "dialog.tpl"}}
<script src="/static/js/dialog_comm.js"></script>
<script src="/static/js/article.js"></script>

{{template "footer.tpl"}}
