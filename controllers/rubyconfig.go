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

	row := config.Get()

	c.Data["json"] = &row
	c.ServeJSON()
}

func (c RubyConfigController) Post() {
	var config models.Rubyconfig
	config.Id, _ = c.GetInt64("id")
	config.Serialline = c.GetString("line")
	config.Serialspeed, _ = c.GetInt64("speed")
	config.Update()
	c.Ctx.WriteString(strconv.FormatInt(config.Id, 10))
}
