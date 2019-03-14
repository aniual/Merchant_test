package controllers

import (
	_"Merchants_test/models"
	"github.com/astaxie/beego"
	"encoding/json"
)


type Withdraw struct {
	result int
	reason string
	Serialnum string
	TransCationid string
}


type AutoWithdraw struct {
	beego.Controller
}


func (c *AutoWithdraw) Post() {
	c.Data["json"] = map[string]string{"result":"","reason":"","Serialnum":"","Transcationid":auto_draw()}
	c.ServeJSON()
}


func auto_draw() string{
	trans := &Withdraw{TransCationid:"0001"}
	json_auto,_ :=json.Marshal(trans)
	json_encode,_ := Encrypt(json_auto)
	return json_encode
}