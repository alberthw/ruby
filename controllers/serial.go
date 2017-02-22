package controllers

import (
	"log"
	"time"

	"github.com/alberthw/ruby/models"
	"github.com/alberthw/ruby/serial"
	"github.com/astaxie/beego"
)

type SerialController struct {
	beego.Controller
}

func (c *SerialController) Open() {
	//	models.GConfig.Id, _ = c.GetInt64("Id")
	//	models.GConfig.Serialname = c.GetString("Serialname")
	//	models.GConfig.Serialbaud, _ = c.GetInt64("Serialbaud")

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

func (c *SerialController) Post() {
	command := c.GetString("command")
	err := serial.Writer([]byte(command + "\n"))
	var result string
	if err != nil {
		log.Println(err.Error())
		result = err.Error()
	} else {
		result = "ok"
	}
	var com models.Command
	com.CommandType = models.SEND
	com.Info = command
	com.InsertCommand()

	time.Sleep(time.Second * 2)
	result = string(serial.GBuffer)
	//	serial.GBuffer = nil
	c.Data["json"] = &result
	c.ServeJSON()
}

func (c *SerialController) Get() {
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
	result := serial.Sender([]byte(command + "\r\n"))
	s := string(result)
	log.Println(s)

	c.Data["json"] = &s
	c.ServeJSON()
}
