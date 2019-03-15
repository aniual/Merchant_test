package controllers

import (
	_"Merchants_test/models"
	"github.com/astaxie/beego"
	"encoding/json"
	"Merchants_test/models"
	"fmt"
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
	var num models.AutoDeposit
	data := c.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &num)
	fmt.Println("Serialnum:",num.Serialnum)
	auto := &Withdraw{TransCationid:"001"}
	json_auto,_ :=json.Marshal(auto)
	//对交易id进行aes加密
	json_encode,_ := Encrypt(json_auto)
	if err == nil {
		c.Data["json"] = map[string]interface{}{"result":0,"reason":"","Serialnum":num.Serialnum,"TransCationid":json_encode}
	}else {
		c.Data["json"] = map[string]interface{}{"result":1,"reason":"","Data":err.Error()}
	}
	c.ServeJSON()
}