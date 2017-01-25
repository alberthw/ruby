package controllers

import (
	"github.com/alberthw/ruby/models"
	"github.com/alberthw/ruby/mongoose"
	"github.com/astaxie/beego"
)

type CommandController struct {
	beego.Controller
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
