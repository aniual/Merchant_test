package controllers

import (
	_ "fmt"
	"github.com/astaxie/beego"
	"encoding/json"
	"net/url"
	"strings"
	"strconv"
	"fmt"
)


type GameListController struct {
	beego.Controller
}

//定义gamelist的数据结构类型
type GameListData struct {
	GameListDataArray map[int]GameListDataUnit
}

//gamelist输出结构类型
type GameListDataUnit struct {
	GameName string
	Url string
}

//创建用户的结构类型
type coCreaterPlay struct {
	GameUserID  int    `json:"GameUserID"`
	AccessToken string `json:"AccessToken"`

}


//get请求游戏列表数据
func (c *GameListController)  Get() {
	var game_name= make([]GameListDataUnit,0)
	var user interface{}
	//通过session获取当下客户的用户名
	user = c.GetSession("loginuser")
	//接口传递过来数据类型需要断言 即user.(string)
	res :=&CreatePlay{
		MerchantId:"XBW001",
		CoUserName:user.(string)} //请求api中的Data的提取
		//调用Access函数，返回Data
	Data := Access("getaccesstoken",res)
	//调用get函数
	s := Get()
	//切片s便利到前端页面
	//参数传递解析
	params := url.Values{}

	for _,v := range s{
		Url, err := url.Parse(v.Url)
		if err != nil {
			panic(err.Error())
		}
		//调用GetAccessToken()方法返回值
		var ID string
		//调用Decrypt方法解密Data
		data,_ := Decrypt(Data)
		var list coCreaterPlay
		//解码返回data json字符串
		if err := json.Unmarshal([]byte(data),&list);err != nil{
			panic(err)
		}
		//进行整型转字符串传入

		ID = strconv.Itoa(list.GameUserID)
		params.Set("CoUserName", user.(string))
		params.Set("nickname", user.(string))
		params.Set("AccessToken", list.AccessToken)
		params.Set("terminaltype", "MacOS")
		params.Set("GameUserID", ID)
		params.Set("merchantid", num)
		params.Set("model", "2")
		params.Set("music", "true")
		params.Set("SoundEffect", "true")
		params.Set("BackUrl", "close")
		//如果参数中有中文参数,这个方法会进行URLEncode
		Url.RawQuery = params.Encode()
		urlPath := Url.String()
		new_url := strings.Replace(urlPath,"&","&&",10)
		//fmt.Println("new_url:",new_url)
		a := GameListDataUnit{v.GameName,new_url}
		//v.Url = new_url
		//s :=append(s, new_url)
		game_name = append(game_name, a)
		//fmt.Println("game_name:",game_name)

	}
	c.Data["gamename"] = game_name
	//fmt.Println("new_url:",new_url)
	//fmt.Println()
	c.Data["s"] = s
	//c.Data["new_url"]=new_url
	c.TplName = "game.html"
}


//传递参数游戏的链接game_url
func Get() []GameListDataUnit{
	//初始化gamename数组
	var gamename= make([]GameListDataUnit,0)
	// 返回解码后的gamelist的data值
	s, _ := Decrypt(GameList())
	list := &GameListData{}
	//对gamelist的datajson进行解码
	if err := json.Unmarshal([]byte(s), &list); err != nil {
		panic(err)
	}
	for _, v := range list.GameListDataArray {
		gamename = append(gamename,v)
		//fmt.Println(v.Url)
	}
	return gamename
}


/*CoUserName=xxx&&nickname=xxx&&AccessToken=xxx&&terminaltype=xxx&&
GameUserID=xxx&&merchantid=xxx&&model=xx&&music=true&&SoundEffect=true&&
BackUrl=" + encodeURIComponent(http://xxx/?parm1=xxx))

必须参数：CoUserName（商户平台内的用户名），，AccessToken，GameUserID，
merchantid，model（同上）， BackUrl, NickName

nickname用户昵称，这个参数如果没有，请使用用户名填充

terminaltype 终端类型
Windows PC    0;
iOS  1;
Android 2;
Mac OS 3;

参数BackUrl处理方式：
（1）“BackUrl=close”，点击返回按钮关闭当前窗口。
（2）BackUrl=“要返回的地址”，点击返回按钮，返回到指定的url地址。
（3）“BackUrl=”，屏蔽返回按钮。

商户带入参数需要先对BackUrl进行encodeURIComponent()编码，
再对整个参数进行encodeURIComponent()编码*/




/*params := url.Values{}
Url, err := url.Parse(game_url)
if err != nil {
	panic(err.Error())
}
//调用GetAccessToken()方法返回值
var ID string
Data := GetAccessToken()
data,_ := Decrypt(Data)
//list_ := &coCreatePlay{}
var list coCreaterPlay
if err := json.Unmarshal([]byte(data),&list);err != nil{
	panic(err)
}
//进行整型转字符串传入
ID = strconv.Itoa(list.GameUserID)
params.Set("CoUserName", "a")
params.Set("nickname", "a")
params.Set("AccessToken", list.AccessToken)
params.Set("terminaltype", "MacOS")
params.Set("GameUserID", ID)
params.Set("merchantid", "XBW001")
params.Set("model", "2")
params.Set("music", "true")
params.Set("SoundEffect", "true")
params.Set("BackUrl", "close")
//如果参数中有中文参数,这个方法会进行URLEncode
Url.RawQuery = params.Encode()
urlPath := Url.String()
new_url := strings.Replace(urlPath,"&","&&",10)
//fmt.Println(new_url)
s := string(new_url)
return s*/







//var gamename= make([]GameListDataUnit,0)
//初始化gamename数组
/*var gamename = make([]string,0)
// 返回解码后的gamelist的data值
s, _ := Decrypt(GameList())
list := &GameListData{}
//fmt.Println(s)
//对gamelist的datajson进行解码
if err := json.Unmarshal([]byte(s), &list); err != nil {
	panic(err)
}
var Url string
for _, v := range list.GameListDataArray {
	//gamename = append(gamename,v.GameName)
	gamename = append(gamename,Get(v.Url))
	//fmt.Println(v.Url)
}*/
//fmt.Println(gamename)
//c.Data["gamename"] = gamename
//fmt.Println(gamename)
//c.Data["url"] = Url