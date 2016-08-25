package controllers

import (
	"strconv"

	"github.com/alberthw/ruby/models"
	"github.com/astaxie/beego"
)

type RubyConfigController struct {
	beego.Controller
}

func (c RubyConfigController) Get() {
	var config models.Rubyconfig

	row := config.Get()

	c.Data["json"] = &row
	c.ServeJSON()
}

func (c RubyConfigController) Post() {
	var config models.Rubyconfig
	config.Id, _ = c.GetInt64("Id")
	config.Serialname = c.GetString("Serialname")
	config.Serialbaud, _ = c.GetInt64("Serialbaud")
	//	config.Isconnected, _ = c.GetBool("connect")
	config.Update()
	c.Ctx.WriteString(strconv.FormatInt(config.Id, 10))
}
