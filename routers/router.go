package routers

import (
	"Merchants_test/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/login",&controllers.LoginController{},"get:Get;post:Post")
	beego.Router("/register",&controllers.RegisterController{},"get:Get;post:Post")
	beego.Router("/gamelist",&controllers.GameListController{})
	/*ns := beego.NewNamespace("/",
			beego.NSNamespace("/autodeposit",
				beego.NSRouter("/",&controllers.AutoController{}),
			),
			beego.NSNamespace("/autowithdraw",
				beego.NSRouter("/",&controllers.AutoWithdraw{}),
			),
		)
	beego.AddNamespace(ns)*/
	beego.Router("/autodeposit",&controllers.AutoController{})
	beego.Router("/autowithdraw",&controllers.WithController{})
}
