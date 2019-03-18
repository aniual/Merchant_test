package controllers

import (
	_"Merchants_test/models"
	"github.com/astaxie/beego"
	"encoding/json"
	"Merchants_test/models"
)

type Response struct {
	result int
	reason string
	Data string
}

type AutoDepositDataunit struct {
	Amount int64
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
	auto := &AutoDepositDataunit{1000,"001",num.Serialnum}
	//Datajson进行加密
	json_auto,_ :=json.Marshal(auto)
	json_encode,_ := Encrypt(json_auto)
	if err == nil {
		c.Data["json"] = map[string]interface{}{"result":0,"reason":"","Data":json_encode}
	}else {
		c.Data["json"] = map[string]interface{}{"result":1,"reason":"","Data":err.Error()}
	}
	c.ServeJSON()
}



