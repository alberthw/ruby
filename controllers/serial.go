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
	models.GConfig.Id, _ = c.GetInt64("Id")
	models.GConfig.Serialname = c.GetString("Serialname")
	models.GConfig.Serialbaud, _ = c.GetInt64("Serialbaud")

	result := "ok"

	c.Data["json"] = &result
	c.ServeJSON()
}

func (c *SerialController) Close() {
	err := serial.Close()
	var result string
	if err != nil {
		log.Println(err.Error())
		result = err.Error()
	} else {
		result = "ok"
	}
	c.Data["json"] = &result
	c.ServeJSON()
}

func (c *SerialController) Write() {
	command := c.GetString("command")
	err := serial.Writer([]byte(command + "\r"))
	var result string
	if err != nil {
		log.Println(err.Error())
		result = err.Error()
	} else {
		result = "ok"
	}
	c.Data["json"] = &result
	c.ServeJSON()

}

func (c *SerialController) Read() {
	var result string
	b, err := serial.Reader()
	if err != nil {
		log.Println(err.Error())
		result = err.Error()
	} else {
		result = string(b)
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
