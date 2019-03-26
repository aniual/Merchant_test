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
	loggerConfig := `{
		"filename":"logs/test.log",
		"maxlines" : 1000,
		"maxsize"  : 10240,
		"rotate": true,
		"daily":true,
		"maxdays":10
	}`
	beego.SetLogger("file",loggerConfig)
	beego.Run()
}



