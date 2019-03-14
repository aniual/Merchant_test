package controllers

import (
	_"Merchants_test/models"
	"github.com/astaxie/beego"
	"encoding/json"
	"fmt"
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
	/*var num models.AutoWithdraw
	data := c.Ctx.Input.RequestBody
	fmt.Println(data)
	json.Unmarshal(data, &num)
	c.ServeJSON()*/
	var data map[string]interface{}
	fmt.Println(c.Ctx.Input.RequestBody)
	json.Unmarshal(c.Ctx.Input.RequestBody, &data)
	fmt.Println(data["MiniAmount"])
	c.ServeJSON()
}



func auto_response() string{
	auto := &AutoDepositDataunit{100,"00a1",""}
	json_auto,_ :=json.Marshal(auto)
	json_encode,_ := Encrypt(json_auto)
	return json_encode
}




/*func (c *AutoController) Delete(){
	idStr := c.Ctx.Input.Param(":id")
	fmt.Println(idStr)
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteApp(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}*/