package main

import (
	//"birthday/daemon"
	//"birthday/models/bmob"
	_ "birthday/routers"
	"github.com/astaxie/beego"
	//"time"
	//"fmt"
)

func main() {
	// go func() {
	// 	for {
	// 		if time.Now().Hour() == 8 {
	// 			daemon.Notify()
	// 			time.Sleep(time.Second * 60 * 40)
	// 		}

	// 	}
	// }()
	//bmob.ShopList()

	beego.Run()
}
