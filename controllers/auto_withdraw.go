package controllers

import (
	_"Merchants_test/models"
	"github.com/astaxie/beego"
	"encoding/json"
	"Merchants_test/models"
	"fmt"
	"github.com/astaxie/beego/orm"
	"strconv"
	"reflect"
)


type Withdraw struct {
	result int
	reason string
	Serialnum string
	TransCationid string
}


type WithController struct {
	beego.Controller
}


func (c *WithController) Post() {
	var num models.AutoWithdraw
	data := c.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &num)
	//fmt.Println("Serialnum:",num.Serialnum)
	auto := &Withdraw{TransCationid:"001"}
	json_auto,_ :=json.Marshal(auto)
	//对交易id进行aes加密
	json_encode,_ := Encrypt(json_auto)
	amount := strconv.FormatFloat(num.Amount,'E',-1,64)
	fmt.Println(reflect.TypeOf(amount))
	amount1,err := strconv.ParseFloat(amount,64)
	//对Amount进行解密
	if err == nil {
		o := orm.NewOrm()
		var u models.User
		err = o.Raw("SELECT money FROM user WHERE username = ?", num.CoUserName).QueryRow(&u)
		err := o.Read(&u)
		if err == nil{
			fmt.Println("money:",u.Money)
		}
		//调用money的值进行赋值
		mon := u.Money
		u = models.User{Id:1}
		if o.Read(&u) == nil{
			u.Money = mon + amount1
			if s,err := o.Update(&u,"money");err == nil{
				fmt.Println("s:",s)
			}
		}
		c.Data["json"] = map[string]interface{}{"result":0,"reason":"","Serialnum":num.Serialnum,"TransCationid":json_encode}
	}else {
		c.Data["json"] = map[string]interface{}{"result":1,"reason":"请求错误","Data":err.Error()}
	}
	c.ServeJSON()
}