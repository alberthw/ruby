package controllers

import (
	"github.com/alberthw/ruby/models"
	"github.com/astaxie/beego"
	"strconv"
)

type RubyConfig struct {
	beego.Controller
}

type RubyConfigController struct {
	beego.Controller
}

func (c RubyConfigController) Get() {
	var config models.Rubyconfig

	lines := config.Get()

	c.Data["json"] = &lines
	c.ServeJSON()
}

func (c *RubyConfigController) Post() {
	var config models.Rubyconfig
	config.Serialline = c.GetString("Serialline")
	config.Serialspeed, _ = c.GetInt64("Serialspeed")
	config.Insert()
	c.Ctx.WriteString(strconv.FormatInt(config.Id, 10))
}
