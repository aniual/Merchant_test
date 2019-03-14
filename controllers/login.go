package controllers

import (
	_ "fmt"
	"github.com/astaxie/beego"
	_"strings"
	"github.com/astaxie/beego/orm"
	"Merchants_test/models"
)

type LoginController struct {
	beego.Controller
}

///login登录get请求
func (c *LoginController) Get() {
	c.TplName = "login.html"
}

//login登录post请求
func (c *LoginController) Post(){
	//1.拿到数据
	o := orm.NewOrm()
	user := models.User{}
	username := c.GetString("Username")
	password := c.GetString("Password")
	//2.判断是否合法
	if username=="" || password == ""{
		c.Abort("输入错误")
		c.TplName = "login.html"
		return
	}
	user.Username = username
	user.Password = password
	err := o.Read(&user,"Username")
	if err != nil{
		c.Abort("输入错误")
		c.TplName = "login.html"
		return
	}
	c.Redirect("/gamelist",302)
}

