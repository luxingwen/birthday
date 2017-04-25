package controllers

import (
	"birthday/models/bmob"
)

type ShopCakeController struct {
	BaseController
}

func (this *ShopCakeController) Trans() {
	list, err := bmob.Transh()
	if err != nil {
		this.Fail(400, err.Error())
	}

	for _, item := range list {
		err := item.Update()
		if err != nil {
			this.Fail(400, err.Error())
		}
	}
	this.Succuess(list)
}
