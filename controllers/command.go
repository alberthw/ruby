package controllers

import (
	"time"

	"github.com/alberthw/ruby/models"
	"github.com/alberthw/ruby/mongoose"
	"github.com/alberthw/ruby/serial"
	"github.com/astaxie/beego"
)

type CommandController struct {
	beego.Controller
}

func (c *CommandController) Get() {
	c.TplName = "command.html"
	c.Layout = "layout.html"
}

func (c CommandController) GetSendCommands() {
	limit, _ := c.GetInt64("limit")
	result, _ := models.GetSendCommands(limit)
	c.Data["json"] = &result
	c.ServeJSON()
}

func (c CommandController) GetReceiveCommands() {
	limit, _ := c.GetInt64("limit")
	result, _ := models.GetReceiveCommands(limit)
	c.Data["json"] = &result
	c.ServeJSON()
}

func (c CommandController) GetVersions() {
	mongoose.SendGetVersions()
	result := "ok"
	c.Data["json"] = &result
	c.ServeJSON()
}

func (c CommandController) GetLastKnownVersions() {
	mongoose.SendGetLastKnownVersions()
	result := "ok"
	c.Data["json"] = &result
	c.ServeJSON()
}

func (c CommandController) EnterServiceMode() {
	mongoose.SendEnterSeviceMode()
	time.Sleep(time.Second)
	result := string(serial.GBuffer)
	serial.GBuffer = nil
	c.Data["json"] = &result
	c.ServeJSON()
}

func (c CommandController) ExitServiceMode() {
	mongoose.SendExitServiceMode()
	time.Sleep(time.Second)
	result := string(serial.GBuffer)
	serial.GBuffer = nil
	c.Data["json"] = &result
	c.ServeJSON()
}
