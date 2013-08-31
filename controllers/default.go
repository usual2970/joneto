package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/lxy/web/models"
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

func init() {
	orm.RegisterModel(new(models.Jt_user))
	orm.RegisterModel(new(models.Jt_content))
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	// 参数1   自定义的数据库名称，用来在orm中切换数据库使用
	// 参数2   driverName
	// 参数3   对应的链接字符串
	// 参数4   设置最大的空闲连接数，使用 golang 自己的连接池
	orm.RegisterDataBase("default", "mysql", "root:123456!a@tcp(112.124.17.181:3306)/joneto?charset=utf8", 30)
	orm.BootStrap() // 强制在 main 函数里调用，检查 Model 关系，检测数据库参数，调用 orm 提供的 Command
}