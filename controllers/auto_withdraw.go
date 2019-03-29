package controllers

import (
	_"Merchants_test/models"
	"github.com/astaxie/beego"
	"encoding/json"
	"Merchants_test/models"
	"github.com/astaxie/beego/orm"
	"strconv"
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
	if err := json.Unmarshal(data, &num);err != nil{
		panic(err)
	}
	//amount := strconv.FormatFloat(num.Amount,'E',-1,64)
	//对Amount进行解密
	amount := num.Amount
	amount,err :=Decrypt(amount)
	beego.Trace("退出金额Amount:",amount)
	auto := &Withdraw{TransCationid:"001"}
	json_auto,_ :=json.Marshal(auto)
	//对交易id进行aes加密
	json_encode,_ := Encrypt(json_auto)
	amount1,err := strconv.ParseFloat(amount,64)
	//对Amount进行解密
	if err == nil {
		o := orm.NewOrm()
		var u models.User
		err = o.Raw("SELECT id,money FROM user WHERE username = ?", num.CoUserName).QueryRow(&u)
		//err := o.Read(&u)
		if err != nil{
			//beego.Trace("money:",u.Money)
		}
		//调用money的值进行赋值
		mon := u.Money
		beego.Trace("数据库金额mon:",mon)
		u = models.User{Id:u.Id}
		if o.Read(&u) == nil{
			u.Money = mon + amount1
			beego.Trace("数据库总金额u.Money:",u.Money)
			if s,err := o.Update(&u,"money");err != nil{
				beego.Trace("获取收益:",s)
				c.Data["json"] = map[string]interface{}{"result":1,"reason":"入库失败","Serialnum":"","TransCationid":""}
			}else {
				c.Data["json"] = map[string]interface{}{"result":0,"reason":"","Serialnum":num.Serialnum,"TransCationid":json_encode}
			}
		}

	}else {
		c.Data["json"] = map[string]interface{}{"result":1,"reason":"参数错误","Data":err.Error()}
	}
	c.ServeJSON()
}