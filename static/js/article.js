$(document).ready(function(){
	rend_list();
	$(window).unbind("scroll").bind("scroll",function(){
		scroll_func();
	});
	$(document).off("click","#jt_comm").on("click","#jt_comm",function(){
		$.dialog_comm();
	});

	
});

function scroll_func(){
	
	var dh=$(document).height();
	var wh=$(window).height();
	var sh=$(window).scrollTop();
	presh=!!window.presh?window.presh:0;
	if(dh-wh-sh<100){
		if (sh==presh){return false;}
		window.pre_presh=presh;
		window.presh=sh;
		$(window).unbind("scroll");
		rend_list();
	}
	
}


function rend_list(){
	var st=window.st;
	if(!st) st=0;
	$.ajax({
		type:"GET",
		url:"/list/"+st,
		dataType:"json",
		success:function(rs){
			
			if(rs.Cnt==0){}
			else{
				var tpl=$("#list_tpl").html();
				var html=juicer(tpl,rs);
				$(html).find("#list_item0").appendTo(".jt_cont_li_wrap1");
				$(html).find("#list_item1").appendTo(".jt_cont_li_wrap2");
				$(html).find("#list_item2").appendTo(".jt_cont_li_wrap3");

				if(rs.Cnt==9) {
					$(window).bind("scroll",function(){
						scroll_func();
					});
				}
				window.st=st+9;
				
			}
			return false;
		}
	});
}
