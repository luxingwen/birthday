package main

import (
	"birthday/daemon"
	_ "birthday/models"
	_ "birthday/routers"
	"github.com/astaxie/beego"
	//	"time"
)

func main() {
	go func() {
		for {
			//if time.Now().Hour() == 8 {
			daemon.Notify()
			//}

		}
	}()
	beego.Run()
}
