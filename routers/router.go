package routers

import (
	"Merchants_test/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)



func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))


    beego.Router("/", &controllers.MainController{})
    beego.Router("/login",&controllers.LoginController{},"get:Get;post:Post")
	beego.Router("/register",&controllers.RegisterController{},"get:Get;post:Post")
	beego.Router("/gamelist",&controllers.GameListController{})
	ns := beego.NewNamespace("/",
			beego.NSNamespace("/autodeposit",
				beego.NSInclude(
					&controllers.AutoController{},
					),
			),
			beego.NSNamespace("/autowithdraw",
				beego.NSInclude(
					&controllers.WithController{},
				),
			),
			beego.NSNamespace("/getusertotalbalance",
				beego.NSInclude(
					&controllers.GetUserController{},
				),
			),
		)
	beego.AddNamespace(ns)
}
