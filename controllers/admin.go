package controllers

import (
	"github.com/astaxie/beego"
	"strings"
	"encoding/hex"
	"crypto/md5"
)

type AdmLogin struct{
	beego.Controller
}
//登录后台页面
func (this *AdmLogin) Get(){
	this.Data["siteurl"]=beego.AppConfig.String("siteurl")
	this.TplNames = "amd_login.tpl"
} 
//登录后台
func (this *AdmLogin) Post(){
	this.Data["siteurl"]=beego.AppConfig.String("siteurl")
	email:=this.GetString("email")
	password:=this.GetString("password")
	err:=""
	if strings.EqualFold(email,""){
		err="用户名不能为空"
	}else if strings.EqualFold(password,""){
		err="密码不能为空"
	}
	this.TplNames = "amd_login.tpl"
	if checkUser(email,password){
		this.SetSession("email",email)
		this.TplNames="amd_xxx.tpl"

	}else{
		err="用户名或密码不正确"
	}
	this.Data["err"]=err
	this.Data["email"]=email
	this.Data["password"]=password
}

func checkUser(email,password string) (bool){
	jone:=beego.AppConfig.String("email")
	to:=beego.AppConfig.String("password")
	enpass:=strings.Join([]string{"joneto",password},"")
	h := md5.New()
    h.Write([]byte(enpass)) 
    enpass=hex.EncodeToString(h.Sum(nil))
	return strings.EqualFold(jone,email) && strings.EqualFold(to,enpass)
}