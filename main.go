package main

import (
	_ "Merchants_test/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func main() {
	beego.BConfig.WebConfig.Session.SessionOn=true
	beego.InsertFilter("*", beego.BeforeRouter,cors.Allow(&cors.Options{
		AllowOrigins:[]string{"http://localhost:8181/login"},
		AllowMethods:[]string{"PUT","PATCH","POST","GET"},
		AllowHeaders:[]string{"origin"},
		ExposeHeaders:[]string{"Content-Length"},
		AllowAllOrigins:true,
	}))
	//添加日志
	loggerConfig := `{
		"filename":"logs/test.log",
		"maxlines" : 100000,
		"maxsize"  : 102400,
		"rotate": true,
		"daily":true,
		"maxdays":5
	}`
	beego.SetLogger("file",loggerConfig)
	beego.BeeLogger.DelLogger("console")
	beego.Run()
}



