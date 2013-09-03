package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)


type Jt_content struct{
	Id          	int        	`orm:"auto"`            		
	Title   		string     	`orm:"size(255);unique"`
	Desc       		string     	``       
	Content    		string     	``
	Hits    		int16     	``
	Has_img    		int16     	``
	Images    		string  	`orm:"size(255)"`
	Content_type  	int16  		`` 
	Add_time		time.Time 	`orm:"auto_now_add;type(datetime)"`
	Last_edit		time.Time 	`orm:"auto_now;type(datetime)"`
	Sort_order		int16	   	``
	Category_id 	int    		``
	Album_id		int    		``
	Comments		[]*Jt_comment `orm:"rel(m2m)"`
	orm.Manager `json:"-"`
}


// 定义NewModel进行orm.Manager的初始化(必须)
func NewContent() *Jt_content {
	obj := new(Jt_content)
	obj.Manager.Init(obj)
	return obj
}


