package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)


type Jt_user struct{
	Id          int        `orm:"auto"`            		 // 设置为auto主键
	User_name   string     `orm:"size(50);unique"` // 设置字段为unique
	Email       string     `orm:"size(50);unique"`       // 设置string字段长度时,会使用varchar类型
	Password    string     `orm:"size(64)"`
	Mobile    	string     `orm:"size(12)"`
	Mblog    	string     `orm:"size(128)"`
	Add_time    time.Time  `orm:"auto_now_add;type(datetime)"`           // 创建时自动设置时间
	Last_login  time.Time  `orm:"auto_now;type(datetime)"`                          // 每次更新时自动设置时间
	//Jt_content       []*Jt_content `orm:"reverse(many)" json:"-"` // fk 的反向关系
	orm.Manager `json:"-"` // 每个model都需要定义orm.Manager
}


// 定义NewModel进行orm.Manager的初始化(必须)
func NewUser() *Jt_user {
	obj := new(Jt_user)
	obj.Manager.Init(obj)
	return obj
}


