package main

import (
	"sync"

	"time"

	_ "github.com/alberthw/ruby/routers"
	"github.com/astaxie/beego"
)

func reader() {
	for {
		beego.BeeLogger.Info("reader")
		time.Sleep(time.Millisecond * 1000)
	}

}

func writer() {
	for {
		beego.BeeLogger.Info("writer")
		time.Sleep(time.Millisecond * 1000)
	}

}

func main() {

	var wg sync.WaitGroup
	wg.Add(1)
	go beego.Run()
	//	go serial.Run()
	//	go reader()
	//	go writer()

	wg.Wait()

}
