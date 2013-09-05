package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/usual2970/joneto/models"
	"strconv"
)

type Cdata struct{
	Cnt int
	Comments []*models.Jt_comment
}

type CommentController struct{
	beego.Controller
}

func (this *CommentController) Get(){
	this.Data["siteurl"]=beego.AppConfig.String("siteurl")
	id,_:=strconv.ParseInt(this.Ctx.Params[":id"],10,64)
	o := orm.NewOrm()
	var comments []*models.Jt_comment
	cnt, _ := o.QueryTable("Jt_comment").RelatedSel().Filter("id",id).OrderBy("-id").All(&comments)
	this.Data["json"]=&Cdata{int(cnt),comments}
	this.ServeJson()
}
