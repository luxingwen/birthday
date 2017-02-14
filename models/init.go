package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/go-xorm/xorm"
)

var (
	AppConfig config.Configer
)

func init() {
	AppConfig = beego.AppConfig
}

func GetEngine() (engine *xorm.Engine, err error) {
	server := AppConfig.DefaultString("mysql::server", "")
	return xorm.NewEngine("mysql", server)
}
