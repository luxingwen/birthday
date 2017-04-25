package main

import (
	"birthday/daemon"
	_ "birthday/models"
	_ "birthday/routers"
	"time"

	"github.com/astaxie/beego"
)

func main() {

	go func() {
		for {
			if time.Now().Hour() == 8 {
				daemon.Notify()
				time.Sleep(time.Second * 60 * 40)
			}

		}

	}()

	go func() {
		for {
			daemon.SendWchatNotify()
			//一分钟执行一次
			time.Sleep(time.Second * 60 * 60)

		}
	}()

	beego.Run()
}
