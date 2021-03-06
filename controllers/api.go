package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"net/http"
	"strings"
	_"fmt"
	"io/ioutil"

	_"fmt"
)

//获取gamelist的游戏列表
type ApiController struct {
	beego.Controller
}

type Server struct {
	result string
	reason string
	Data string
}

//定义gamelist结构类型
type Body struct {
	MerchantId string
	Sign string
}

//定义create结构类型
type CreatePlay struct {
	MerchantId string
	CoUserName string
	Sign string
}

type GetAccessToken struct {
	MerchantId string
	CoUserName string
	Sign string
}


//api接口请求
//参数key为接口地址
//res为Post参数
func Pubilc(key  string, res *Body) string{
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
	beego.Trace("游戏列表body:",string(body))
	return s.Data
}

//获取游戏列表
func GameList()string{
	res := &Body{
		MerchantId:"YLTEST99"}
	//请求api中的Data的提取
	Data := Pubilc("gamelist",res)
	return Data
}