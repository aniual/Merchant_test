package routers

import (
	"Merchants_test/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/login",&controllers.LoginController{},"get:Get;post:Post")
	beego.Router("/register",&controllers.RegisterController{},"get:Get;post:Post")
	beego.Router("/gamelist",&controllers.GameListController{},"get:Get")
	ns :=
		beego.NewNamespace("/",
			beego.NSNamespace("/autodeposit",
				beego.NSInclude(
					&controllers.AutoController{},
				),
			),
			beego.NSNamespace("/autowithdraw",
				beego.NSInclude(
					&controllers.AutoWithdraw{},
				),
			),
		)
	beego.AddNamespace(ns)
}
