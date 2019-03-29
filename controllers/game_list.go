package controllers

import (
	_ "fmt"
	"github.com/astaxie/beego"
	"encoding/json"
	"net/url"
	"sort"
	"Merchants_test/models"
	"github.com/astaxie/beego/orm"
	"strconv"
	"fmt"
	"strings"
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
	//var user string
	//通过session获取当下客户的用户名
	user := c.GetSession("loginuser")
	//接口传递过来数据类型需要断言 即user.(string)
	res :=&GetAccessToken{
		MerchantId:"YLTEST99",
		CoUserName:user.(string)} //请求api中的Data的提取
	//调用Access函数，返回Data
	Data := Access("getaccesstoken", res)
	//调用get函数
	s := Get()
	//参数传递解析
	//num:= c.Ctx.GetCookie("number")
	//params := url.Values{}
	//	遍历s对url传递
	for _,v := range s{
		//调用GetAccessToken()方法返回值
		//调用Decrypt方法解密Data
		data,_ := Decrypt(Data)
		var list coCreaterPlay
		//解码返回data json字符串
		if err := json.Unmarshal([]byte(data),&list);err != nil{
			panic(err)
		}
		gameid := strconv.Itoa(list.GameUserID)
		//192.168.2.102服务环境地址
		keysUrl := "CoUserName=" + "YLTEST99_"+user.(string) + "&&nickname=" + user.(string) + "&&AccessToken=" +list.AccessToken + "&&terminaltype=" + "MacOS"+ "&&GameUserID=" + gameid + "&&merchantid=" + "YLTEST99" + "&&model=" + "2" + "&&music=" + "true" + "&&SoundEffect=" + "true" + "&&BackUrl=" + encodeURIComponent("http://192.168.4.216:8181/gamelist/?")
		//192.168.4.216本地测试地址
		//keysUrl := "CoUserName=" + name + "&&nickname=" + user.(string) + "&&AccessToken=" +list.AccessToken + "&&terminaltype=" + "MacOS"+ "&&GameUserID=" + gameid + "&&merchantid=" + "YLTEST99" + "&&model=" + "2" + "&&music=" + "true" + "&&SoundEffect=" + "true" + "&&BackUrl=" + encodeURIComponent("http://192.168.4.216:8181/gamelist/?")
		keysURL := encodeURIComponent(keysUrl)
		kUrl := v.Url + "/?"+ keysURL
		//fmt.Println(keysUrl)
		//keysURL := encodeURIComponent(keysUrl)
		//fmt.Println(keysURL)a
		/*Url, err := url.Parse(v.Url)
		if err != nil {
			panic(err.Error())
		}
		//调用GetAccessToken()方法返回值
		//调用Decrypt方法解密Data
		data,_ := Decrypt(Data)
		var list coCreaterPlayg
		//解码返回data json字符串
		if err := json.Unmarshal([]byte(data),&list);err != nil{
			panic(err)
		}
		//进行整型转字符串传入
		str := "http://192.168.2.102:8181/gamelist"
		s_url := url.QueryEscape(str)
		gameid := strconv.Itoa(list.GameUserID)
		params.Add("CoUserName", user.(string))
		params.Add("nickname", user.(string))
		params.Add("AccessToken", list.AccessToken)
		params.Add("terminaltype", "MacOS")
		params.Add("GameUserID", gameid)
		params.Add("merchantid", "XBW001")
		params.Add("model", "2")
		params.Add("music", "true")
		params.Add("SoundEffect", "true")
		params.Add("BackUrl", s_url)
		//如果参数中有中文参数,这个方法会进行URLEncode
		Url.RawQuery = params.Encode()
		urlPath := Url.String()
		new_url := strings.Replace(urlPath,"&","&&",10)*/
		//fmt.Println(new_url)
		//total_url := url.QueryEscape(new_url)
		//fmt.Println(total_url)
		a := GameListDataUnit{v.GameName,kUrl}
		game_name = append(game_name, a)
	}

	var u models.User
	o := orm.NewOrm()
	o.Raw("SELECT money FROM user WHERE username = ?", user).QueryRow(&u)
	//调用money的值进行赋值
	mon := u.Money
	c.Data["money"] = mon
	c.Data["username"] = user
	c.Data["gamename"] = game_name
	c.Data["merchantsid"] = "YLTEST99"
	c.TplName = "game.html"

}


func (c *GameListController) post() {
	c.DelSession("username")
	c.Redirect("/",302)
}


//此函数为
func (c *GameListController) Post(){
	//取到充值数据
	money,err := c.GetFloat("money")
	//fmt.Println("money:",money)
	if err == nil{
		/*if money <= 0 {
			beego.Info("输入正确金额")
			c.Redirect("/gamelist",302)
			return
		}else{*/
		//对用户输入金额对数据库进行更改
			o := orm.NewOrm()
			var u models.User
			u = models.User{Id:1}
			err := o.Raw("SELECT money FROM user WHERE id = ?", u.Id).QueryRow(&u)
			if err == nil{
				fmt.Println("money:",u.Money)
			}
			//调用money的值进行赋值
			mon := u.Money
			if o.Read(&u) == nil{
				u.Money = mon + money
				if num,err := o.Update(&u,"money");err == nil{
					fmt.Println(num)
				}
			}
		//}
	}
	c.Redirect("/gamelist",302)
}


//传递参数游戏的链接game_url
func Get() []GameListDataUnit{
	//初始化gamename数组
	var gamename []GameListDataUnit
	// 返回解码后的gamelist的data值
	s, _ := Decrypt(GameList())
	list := &GameListData{}
	//对gamelist的datajson进行解码
	if err := json.Unmarshal([]byte(s), &list); err != nil {
		panic(err)
	}
	//用map的key键来进行排序
	var names []int
	//将值添加到names切片中
	for name := range list.GameListDataArray {
		names = append(names,name)
	}
	//调用sort方法有序化
	sort.Ints(names)
	//遍历key进行排序
	//使用key有序
	for _,name := range names{
		gamename = append(gamename, list.GameListDataArray[name])
	}
	return gamename
}


func encodeURIComponent(str string) string{
	r := url.QueryEscape(str)
	r = strings.Replace(r,"+","%20",-1)
	return r
}