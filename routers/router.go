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
	beego.Router("/gamelist",&controllers.GameListController{},"get:Get")
}
