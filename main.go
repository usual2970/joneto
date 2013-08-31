package main

import (
	"github.com/lxy/web/controllers"
	"github.com/astaxie/beego"
)
func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/list/:st([0-9]+)", &controllers.ListController{})
	beego.Router("/art/:id([0-9]+)",&controllers.ArtController{})
	beego.Run()
}