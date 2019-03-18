package controllers

import (
	_"Merchants_test/models"
	"github.com/astaxie/beego"
	"encoding/json"
	"Merchants_test/models"
	_"strconv"
	"github.com/astaxie/beego/orm"
	"fmt"
)

type Response struct {
	result int
	reason string
	Data string
}

type AutoDepositDataunit struct {
	Amount float64
	Transcationid string
	Serialnum string
}



type AutoController struct {
	beego.Controller
}


func (c *AutoController) Post() {
	var num models.AutoDeposit
	data := c.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &num)
	//fmt.Println(err)
	if err == nil {
		//创建一个orm对象
		o := orm.NewOrm()
		var u models.User
		err = o.Raw("SELECT money FROM user WHERE username = ?", num.CoUserName).QueryRow(&u)
		fmt.Println(u.Money)
		//创建对象
		var mon float64
		mon = u.Money
		auto := &AutoDepositDataunit{num.ReqAmount,"001",num.Serialnum}
		//Datajson进行加密
		json_auto,_ :=json.Marshal(auto)
		json_encode,_ := Encrypt(json_auto)

		if num.ReqAmount <= mon {
			c.Data["json"] = map[string]interface{}{"result":0,"reason":"","Data":json_encode}
		}else {
			c.Data["json"] = map[string]interface{}{"result":1,"reason":"","Data":"err"}
		}
	}else {
		c.Data["json"] = map[string]interface{}{"result":1,"reason":"","Data":err.Error()}
	}

	c.ServeJSON()
}



