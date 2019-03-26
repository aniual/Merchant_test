package controllers

import (
	_ "fmt"
	"github.com/astaxie/beego"
	_"strings"
	"github.com/astaxie/beego/orm"
	"Merchants_test/models"
	"encoding/json"
	"net/http"
	"strings"
	"io/ioutil"
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
	number := c.GetString("number")
	beego.Trace("number",number)
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
		c.Data["err"]="用户名不存在"
		c.TplName = "login.html"
		return
	}
	if user.Password != password{
		c.Data["err"] = "密码错误"
		c.TplName = "login.html"
		return
	}
	//设置用户名session
	c.SetSession("loginuser",username)
	//设置商户号cookie
	c.Ctx.SetCookie("number",number)
	res :=&CreatePlay{
		MerchantId:"YLTEST99",
		CoUserName:username} //请求api中的Data的提取
	Pubilc_("createplayer",res)
	c.Redirect("/gamelist",302)
}


//用于请求的公共代码，直接调用此方法
func Pubilc_(key  string, res *CreatePlay) string{
	s :=&Server{}
	key_body, _ := json.Marshal(res)
	resp, err := http.Post("http://35.234.53.32:8443/" +key, "application/json", strings.NewReader(string(key_body)))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal([]byte(body),&s);err != nil{
		panic(err)
	}
	beego.Trace("Create_body:",string(body))
	return s.Data
}


func Access(key  string, res *GetAccessToken) string{
	s :=&Server{}
	key_body, _ := json.Marshal(res)
	resp, err := http.Post("http://35.234.53.32:8443/" +key, "application/json", strings.NewReader(string(key_body)))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal([]byte(body),&s);err != nil{
		panic(err)

	}
	beego.Trace("Access_Body",string(body))
	return s.Data
}


