package controllers

import (
	"log"

	"github.com/alberthw/ruby/models"
	"github.com/alberthw/ruby/serial"
	"github.com/astaxie/beego"
)

type SerialController struct {
	beego.Controller
}

func (c *SerialController) Open() {
	var config models.Rubyconfig
	config.Id, _ = c.GetInt64("Id")
	config.Serialname = c.GetString("Serialname")
	config.Serialbaud, _ = c.GetInt64("Serialbaud")
	//	log.Println("name", config.Serialname)
	//	log.Println("baud", config.Serialbaud)
	//	log.Println("status", config.Isconnected)
	err := serial.Open(config.Serialname, int(config.Serialbaud))
	//	log.Println(err.Error())
	var result string
	if err != nil {
		result = err.Error()
		log.Println(result)
	} else {
		result = "ok"
		config.Isconnected = true
		config.UpdateStatus()
	}
	c.Data["json"] = &result
	c.ServeJSON()
}
func (c *SerialController) Close() {
	var config models.Rubyconfig
	config.Id, _ = c.GetInt64("Id")

	err := serial.Close()
	var result string
	if err != nil {
		log.Println(err.Error())
		result = err.Error()
	} else {
		result = "ok"
		config.Isconnected = false
		config.UpdateStatus()
	}
	c.Data["json"] = &result
	c.ServeJSON()
}

func (c *SerialController) Send() {
	command := c.GetString("command")
	result := serial.Sender([]byte(command + "\r"))

	s := string(result)

	c.Data["json"] = &s
	c.ServeJSON()

}
