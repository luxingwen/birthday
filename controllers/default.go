package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (this *BaseController) Fail(code int, msg string) {
	this.Data["json"] = map[string]interface{}{"code": code, "msg": msg}
	this.ServeJSON()
	this.StopRun()
}

func (this *BaseController) Succuess(data interface{}) {
	this.Data["json"] = map[string]interface{}{"code": 0, "msg": "success", "data": data}

	b, err := json.Marshal(data)
	if err != nil {
		fmt.Println("marshal err. ", err)
	}
	this.ServeJSON()

}
