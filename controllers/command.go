package controllers

import (
	"github.com/alberthw/ruby/models"
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
