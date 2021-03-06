package routers

import (
	"birthday/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.BaseController{})
	beego.Router("/trans", &controllers.ShopCakeController{}, "*:Trans")
}
