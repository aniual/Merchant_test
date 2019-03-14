package controllers

import (
	_ "fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"Merchants_test/models"
)

type RegisterController struct {
	beego.Controller
}

type USERNAME struct {
	username string
}

func (c *RegisterController) Get() {
	c.TplName = "register.html"
}

func (c *RegisterController) Post() {
	//1,拿到数据
	username :=c.GetString("Username")
	password :=c.GetString("Password")
	//2.数据校验
	if username == "" || password == ""{
		beego.Info("数据不能为空")
		c.Redirect("/register",302)
		return
	}
	//3.插入数据库
	o := orm.NewOrm()
	user := models.User{}
	user.Username = username
	user.Password = password
	_,err := o.Insert(&user)
	if err!=nil{
		beego.Info("插入数据失败")
		c.Redirect("/register",302)
	}

	//4.返回登录界面
	c.Redirect("/login",302)
}

