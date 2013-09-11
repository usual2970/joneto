package main

import (
	"github.com/usual2970/joneto/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/usual2970/joneto/models"
)
func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/list/:st([0-9]+)", &controllers.ListController{})
	beego.Router("/art/:id([0-9]+)",&controllers.ArtController{})
	beego.Router("/comm/:id([0-9]+)",&controllers.CommentController{})
	beego.Router("/amd/login",&controllers.AdmLogin{})
	beego.Router("/amd",&controllers.AmdArtlist{})
	beego.Run()
}

func init() {
	orm.RegisterModel(new(models.Jt_user))
	orm.RegisterModel(new(models.Jt_content))
	orm.RegisterModel(new(models.Jt_comment))
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	// 参数1   自定义的数据库名称，用来在orm中切换数据库使用
	// 参数2   driverName
	// 参数3   对应的链接字符串
	// 参数4   设置最大的空闲连接数，使用 golang 自己的连接池
	orm.RegisterDataBase("default", "mysql", "root:123456!a@tcp(112.124.17.181:3306)/joneto?charset=utf8", 30)
	orm.BootStrap() // 强制在 main 函数里调用，检查 Model 关系，检测数据库参数，调用 orm 提供的 Command
}