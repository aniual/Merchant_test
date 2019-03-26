package controllers

import (
	"Merchants_test/models"
	_ "Merchants_test/models"
	"encoding/json"
	_ "strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"fmt"
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
		//调用money的值进行赋值
		//err := o.Read(&u)
		if err != nil{
			beego.Trace("money",u.Money)
		}
		//调用money的值进行赋值
		mon := u.Money
		//进行比较接口数据给与返回body
		if num.ReqAmount <= mon {
			u = models.User{Id:1}
			if o.Read(&u) == nil{
				u.Money = mon - num.ReqAmount
				fmt.Println("ReqAmout",num.ReqAmount)
				if s,err := o.Update(&u,"money");err == nil{
					beego.Trace("扣款金额成功",s)
				}
			}
			auto := &AutoDepositDataunit{num.ReqAmount, "001", num.Serialnum}
			//auto进行解析
			json_auto, _ := json.Marshal(auto)
			//对json_auto进行加密
			json_encode, _ := Encrypt(json_auto)
			c.Data["json"] = map[string]interface{}{"result": 0, "reason": "", "Data": json_encode}
		}else if num.ReqAmount >= mon && num.MiniAmount <= mon{
			u = models.User{Id:1}
			if o.Read(&u) == nil{
				u.Money = mon - num.MiniAmount
					beego.Trace("ReqAmout",num.MiniAmount)
				beego.Trace("u.Moeny:",u.Money)
				if s,err := o.Update(&u,"money");err == nil{
					beego.Trace("扣款金额成功",s)
				}
			}
			auto := &AutoDepositDataunit{num.MiniAmount, "001", num.Serialnum}
			//对auto进行加密
			json_auto, _ := json.Marshal(auto)
			json_encode, _ := Encrypt(json_auto)
			c.Data["json"] = map[string]interface{}{"result": 0, "reason": "", "Data": json_encode}
		}else if num.MiniAmount > mon{
			c.Data["json"] = map[string]interface{}{"result": 1, "reason": "余额不足带入失败", "Data": "err"}
		}
	} else {
		c.Data["json"] = map[string]interface{}{"result": 1, "reason": "", "Data": err.Error()}
	}
	/*//手动插入测试
	auto := &AutoDepositDataunit{100, "001", num.Serialnum}
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