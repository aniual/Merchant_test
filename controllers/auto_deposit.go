package controllers

import (
	"Merchants_test/models"
	_ "Merchants_test/models"
	"encoding/json"
	"fmt"
	_ "strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Response struct {
	result int
	reason string
	Data   string
}

type AutoDepositDataunit struct {
	Amount        float64
	Transcationid string
	Serialnum     string
}

type AutoController struct {
	beego.Controller
}

func (c *AutoController) Post() {
	var num models.AutoDeposit
	data := c.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &num)
	if err == nil {
		//创建一个orm对象
		o := orm.NewOrm()
		var u models.User
		err = o.Raw("SELECT money FROM user WHERE username = ?", num.CoUserName).QueryRow(&u)
		mon := u.Money
		if num.ReqAmount <= mon {
			auto := &AutoDepositDataunit{num.ReqAmount, "001", num.Serialnum}
			//Datajson进行加密
			json_auto, _ := json.Marshal(auto)
			fmt.Println("json_auto:", string(json_auto))
			json_encode, _ := Encrypt(json_auto)
			fmt.Println("json_encode:", string(json_encode))
			c.Data["json"] = map[string]interface{}{"result": 0, "reason": "", "Data": json_encode}
			fmt.Println("c.Data", c.Data["json"])
		}else if num.ReqAmount >= mon && num.MiniAmount <= mon{
			auto := &AutoDepositDataunit{num.MiniAmount, "001", num.Serialnum}
			//Datajson进行加密
			json_auto, _ := json.Marshal(auto)
			fmt.Println("json_auto:", string(json_auto))
			json_encode, _ := Encrypt(json_auto)
			c.Data["json"] = map[string]interface{}{"result": 0, "reason": "", "Data": json_encode}
		}else if num.ReqAmount >= mon && num.MiniAmount >= mon{
			c.Data["json"] = map[string]interface{}{"result": 1, "reason": "", "Data": err.Error()}
		}
	} else {
		c.Data["json"] = map[string]interface{}{"result": 1, "reason": "", "Data": err.Error()}
	}
	/*auto := &AutoDepositDataunit{100, "001", num.Serialnum}
	//Datajson进行加密
	json_auto, _ := json.Marshal(auto)
	fmt.Println("json_auto:", string(json_auto))
	json_encode, _ := Encrypt(json_auto)
	fmt.Println("json_encode:", string(json_encode))
	if err == nil {
		c.Data["json"] = map[string]interface{}{"result": 0, "reason": "", "Data": json_encode}
	} else {
		c.Data["json"] = map[string]interface{}{"result": 1, "reason": "", "Data": err.Error()}
	}*/
	c.ServeJSON()
}
