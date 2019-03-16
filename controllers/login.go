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
	"fmt"
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
	fmt.Println("username:",username)
	num := c.GetString("number")
	fmt.Println("num:",num)
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
		c.Data["err"]="用户名不存在"
		c.TplName = "login.html"
		return
	}
	if user.Password != password{
		c.Data["err"] = "密码错误"
		c.TplName = "login.html"
		return
	}
	c.SetSession("loginuser",username)
	res :=&CreatePlay{
		MerchantId:"XBW001",
		CoUserName:username} //请求api中的Data的提取
	Pubilc_("createplayer",res)
	c.Redirect("/gamelist",302)
}



//用于请求的公共代码，直接调用此方法
func Pubilc_(key  string, res *CreatePlay) string{
	s :=&Server{}
	key_body, _ := json.Marshal(res)
	resp, err := http.Post("http://192.168.2.102:8443/" +key, "application/json", strings.NewReader(string(key_body)))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal([]byte(body),&s);err != nil{
		panic(err)
	}
	//fmt.Println(string(body))
	return s.Data
}


func Access(key  string, res *CreatePlay) string{
	s :=&Server{}
	key_body, _ := json.Marshal(res)
	resp, err := http.Post("http://192.168.2.102:8443/" +key, "application/json", strings.NewReader(string(key_body)))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal([]byte(body),&s);err != nil{
		panic(err)
	}
	return s.Data
}


