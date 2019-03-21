package models

import (
	_"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"net/url"
	"github.com/astaxie/beego"
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
	//获取配置信息
	dbhost := beego.AppConfig.String("db.host")
	dbuser := beego.AppConfig.String("db.user")
	dbpassword := beego.AppConfig.String("db.password")
	dbport := beego.AppConfig.String("db.port")
	dbname := beego.AppConfig.String("db.name")
	dbtimezone := beego.AppConfig.String("db.timezone")

	if dbport == ""{
		dbport = "3306"
	}
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	if dbtimezone != ""{
		dsn = dsn + "&loc=" + url.QueryEscape(dbtimezone)
	}
	//fmt.Println("dns:",dsn)
	//设置数据库基本信息
	//orm.RegisterDataBase("default","mysql","liuy:123456@tcp(192.168.2.120:3306)/test1?charset=utf8",30)
	orm.RegisterDataBase("default","mysql",dsn)
	orm.RegisterModel(new(User))
	//orm.RunSyncdb("default",false,true)
}
