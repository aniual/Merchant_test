package controllers

import (
	_"Merchants_test/models"
	"github.com/astaxie/beego"
	"encoding/json"
)


type AutoWithdraw struct {
	result int
	reason string
	Serialnum string
	TransCationid string
}


type DrawController struct {
	beego.Controller
}


func (c *DrawController) Post() {
	c.Data["json"] = map[string]string{"result":"","reason":"","Serialnum":"","Transcationid":auto_draw()}
	c.ServeJSON()
}


func auto_draw() string{
	trans := &AutoWithdraw{TransCationid:"0001"}
	json_auto,_ :=json.Marshal(trans)
	json_encode,_ := Encrypt(json_auto)
	return json_encode
}