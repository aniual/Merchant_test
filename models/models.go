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
	Money float64
}


type AutoDeposit struct {
	MerchantId string
	CoUserName string
	ReqAmount float64
	MiniAmount float64
	GameId string
	Type int
	Serialnum string
	Sign string
}



type AutoWithdraw struct {
	MerchantId string
	CoUserName string
	ReqAmount float64
	MiniAmount float64
	GameId string
	Type int
	Serialnum string
	Sign string
}


func init()  {
	//设置数据库基本信息
	orm.RegisterDataBase("default","mysql","liuy:123456@tcp(192.168.2.120:3306)/test1?charset=utf8",30)
	orm.RegisterModel(new(User))
	orm.RunSyncdb("default",false,true)
}
