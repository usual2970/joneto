package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/usual2970/joneto/models"
	"strconv"
)

type Data struct{
	Cnt int
	Arts []*models.Jt_content
}

type MainController struct {
	beego.Controller
}


func (this *MainController) Get() {
	this.TplNames = "article.tpl"
}

type ListController struct{
	beego.Controller
}

func (this *ListController) Get(){
	st,_:=strconv.ParseInt(this.Ctx.Params[":st"],10,64)
	o := orm.NewOrm()
	var contents []*models.Jt_content
	cnt, _ := o.QueryTable("Jt_content").RelatedSel().Limit(9,st).OrderBy("-id").All(&contents)
	this.Data["json"]=&Data{int(cnt),contents}
	this.ServeJson()
}

type ArtController struct{
	beego.Controller
}
func (this *ArtController) Get(){
	this.Data["siteurl"]=beego.AppConfig.String("siteurl")
	id,_:=strconv.ParseInt(this.Ctx.Params[":id"],10,64)
	o := orm.NewOrm()
	prev:=models.NewContent()
	content:=models.NewContent()
	next:=models.NewContent()
	o.QueryTable("Jt_content").RelatedSel().Limit(1).Filter("id",id).One(content)
	o.QueryTable("Jt_content").RelatedSel().Limit(1).OrderBy("id").Filter("id__gt",id).One(next)
	o.QueryTable("Jt_content").RelatedSel().Limit(1).OrderBy("-id").Filter("id__lt",id).One(prev)
	this.Data["next"]=next
	this.Data["content"]=content
	this.Data["prev"]=prev
	this.TplNames = "content.tpl"
}

