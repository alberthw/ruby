package main

import (
	"log"
	"time"

	"github.com/alberthw/ruby/models"
	_ "github.com/alberthw/ruby/routers"
	"github.com/alberthw/ruby/serial"
	"github.com/astaxie/beego"
)

func main() {
	go open()
	//	go read()

	beego.Run()

}

func open() {
	for {
		//		log.Println("open serial")

		models.GConfig = models.GConfig.Get()
		//		log.Println(models.GConfig.Serialname, models.GConfig.Serialbaud, models.GConfig.Isconnected)
		connected := false
		err := serial.Open(models.GConfig.Serialname, int(models.GConfig.Serialbaud))
		if err == nil {
			connected = true
		}
		models.GConfig.Isconnected = connected
		models.GConfig.UpdateStatus()
		time.Sleep(time.Millisecond * 1000)
	}
}

func read() {
	for {
		//	log.Println("read serial")
		b, err := serial.Reader()
		if err != nil {
			beego.BeeLogger.Error(err.Error())
		}
		if len(b) > 0 {
			var res models.Response
			res.Requestid = 1
			res.Content = string(b)
			res.Insert()
			log.Println("output:", string(b))
		}

		time.Sleep(time.Millisecond * 100)
	}
}
