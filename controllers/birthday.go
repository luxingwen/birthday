package controllers

import (
	"birthday/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"log"
)

type BirthDayController struct {
	BaseController
}

func (this *BirthDayController) Add() {
	var birthday models.BirthDay
	body := this.Ctx.Input.CopyBody(beego.BConfig.MaxMemory)
	err := json.Unmarshal(body, &birthday)
	if err != nil {
		log.Println("json unmarshal err: ", err)
		return
	}
	err = models.InsertObj(&birthday)
	if err != nil {
		log.Println("insert obj err: ", err)
		return
	}
	this.Data["json"] = map[string]interface{}{"msg": "ok", "code": 200}
	this.ServeJSON()
}
