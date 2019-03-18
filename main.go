package main

import (
	_ "Merchants_test/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.WebConfig.Session.SessionOn=true
	beego.Run()
}



