package main

import (
	"log"
	"time"

	"encoding/hex"

	"github.com/alberthw/ruby/ebdprotocol"
	"github.com/alberthw/ruby/models"
	_ "github.com/alberthw/ruby/routers"
	"github.com/alberthw/ruby/serial"
	"github.com/astaxie/beego"
)

var bSoftDelete = true

func IncreaseSeq(seq byte) byte {
	if seq == 0x39 {
		return 0x30
	}
	return seq + 1
}

func IncreaseOneSequence() {
	config := models.GConfig.Get()

	//	log.Printf(" seq  :%x", byte(s.Sequence[0]))
	config.Sequence = string(IncreaseSeq(byte(config.Sequence[0])))
	//	log.Printf("Sequence : %s", s.Sequence)
	config.UpdateSequence()
}

func main() {
	go open()
	go generate(200)
	go writer(100)
	go reader(100)

	//	go read()

	beego.Run()

}

func generate(t time.Duration) {
	for {
		config := models.GConfig.Get()
		if config.Isconnected {
			var k ebdprotocol.KeepAlive
			k.SessionKey = []byte(config.Sessionkey)

			k.Sequence = byte(config.Sequence[0])

			log.Printf("session key : %X, sequence : %X\n", k.SessionKey, k.Sequence)

			var m models.Message
			m.Messagetype = models.REQUEST
			m.Info = hex.EncodeToString(k.Message())
			m.Status = models.NONE

			m.InsertMessage()

		}
		time.Sleep(time.Millisecond * time.Duration(t))
	}
}

func writer(t time.Duration) {
	// get one request
	for {
		config := models.GConfig.Get()
		//		time.Sleep(time.Millisecond * time.Duration(s.Writeinterval))
		time.Sleep(time.Millisecond * time.Duration(t))
		if config.Isconnected {
			var m models.Message
			err := m.GetOneRequest()
			if err != nil {
				//				log.Println("writer:", err.Error())
				continue
			}

			//			log.Printf("get one request :%v", m)

			b, _ := hex.DecodeString(m.Info)
			err = serial.Writer(b)
			if err != nil {
				log.Println(err.Error())
				continue
			}
			//			log.Printf("Send:%X", b)
			IncreaseOneSequence()

			if bSoftDelete {
				m.Status = models.DELETED
				m.UpdateStatus()
			} else {
				m.DeleteMessage()
				//		log.Println("the request is deleted.")
			}
		}

	}
}

func reader(t time.Duration) {
	for {
		var m models.Message
		b, _ := serial.Reader()
		if len(b) > 0 {
			m.Messagetype = models.RESPONSE
			m.Info = hex.EncodeToString(b)
			m.Status = models.NONE
			log.Printf("Received : %X", b)
			log.Printf("Received message : %s", m.Info)
			m.InsertMessage()
		}
		time.Sleep(time.Millisecond * t)
	}
}

func open() {
	for {
		models.GConfig = models.GConfig.Get()
		//		log.Printf("serial name : %s, serial baud : %d, connection status : %v", models.GConfig.Serialname, models.GConfig.Serialbaud, models.GConfig.Isconnected)

		connected := false
		err := serial.Open(models.GConfig.Serialname, int(models.GConfig.Serialbaud))

		if err == nil {
			connected = true
		} else {
			log.Println(err.Error())
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
