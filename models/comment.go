package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)


type Jt_comment struct{
	Id          	int        	`orm:"auto"`            		
	User_id   		int     	``
	Content    		string     	``
	Content_id    	int     	``
	Add_time		time.Time 	`orm:"auto_now_add;type(datetime)"`
	Parent_id		*Jt_comment	`orm:"rel(fk)"`
	Contents 		[]*Jt_content `orm:"reverse(many)" json:"-"`
	orm.Manager `json:"-"`
}


// 定义NewModel进行orm.Manager的初始化(必须)
func NewComment() *Jt_comment {
	obj := new(Jt_comment)
	obj.Manager.Init(obj)
	return obj
}


