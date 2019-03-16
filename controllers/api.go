package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"net/http"
	"strings"
	_"fmt"
	"io/ioutil"

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
	return s.Data
}

//获取游戏列表
func GameList()string{
	res := &Body{
		MerchantId:"XBW001"}
	//请求api中的Data的提取
	Data := Pubilc("gamelist",res)
	//fmt.Println(string(Data))
	return Data
}

/*//用于请求的公共代码，直接调用此方法
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

//创建游戏账户
func CreatePlayer() string{
	res :=&CreatePlay{
		MerchantId:"XBW001",
		CoUserName:"a"}
		//请求api中的Data的提取
	Data := Pubilc_("createplayer",res)
	return Data
}*/


//用户访问游戏平台获取AccessToken
func GetAccessToken() string{
	res :=&CreatePlay{
		MerchantId:"XBW001",
		CoUserName:"a"}
	//请求api中的Data的提取
	Data := Pubilc_("getaccesstoken",res)
	return Data
}
