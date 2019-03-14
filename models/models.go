package models

import (
	_"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego/orm"
)

//表的设计
type User struct {
	Id int
	Username string
	Password string
}

func init()  {
	//设置数据库基本信息
	orm.RegisterDataBase("default","mysql","root:12345678@tcp(127.0.0.1:3306)/test?charset=utf8",30)
	orm.RegisterModel(new(User))
	orm.RunSyncdb("default",false,true)
}
