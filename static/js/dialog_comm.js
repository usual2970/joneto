/*
此插件基于Jquery
插件名：jquery.dialog_comm
开发者 usual2970
版本 1.0  
Blog：http://t.qq.com/usual2970
*/
(function($){
	$.dialog_comm=function(option){
		var opts=$.extend({},$.dialog_comm.defaults,option);
		var pos=comp_position(opts.width);
		var tpl=$("#dialog_tpl").html();
		var data={
			pos:pos,
			siteurl:site_url
		};
		var html=juicer(tpl,data);
		$(html).appendTo("body");
		var bdh=pos[1]-$(".jt_dl_hd").height()-$(".jt_dl_ft").height()-40;
		$(".jt_dl_bd").height(bdh);
		$("body #jt_shade").css("opacity","0.7");

		$(document).off("click",".jt_dl_close").on("click",".jt_dl_close",function(){
			dl_close();
		})
	}

	function comp_position(w){
		var wh=$(window).height();
		var h=wh*95/100;
		var ww=$(window).width();
		return [w,h,(ww-w)/2,(wh-h)/2];
	}
	function dl_close(){
		$("body #jt_shade").remove();
		$(".jt_dialog").remove();
	}

	$.dialog_comm.defaults={
		width:720
	};

})(jQuery);