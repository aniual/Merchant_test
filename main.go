package main

import (
	_ "Merchants_test/routers"
	"github.com/astaxie/beego"
	"Merchants_test/controllers"
)

func main() {
	beego.BConfig.WebConfig.Session.SessionOn=true
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}



