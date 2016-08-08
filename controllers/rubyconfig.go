package controllers

import (
	"github.com/alberthw/ruby/models"
	"github.com/astaxie/beego"
)

type RubyConfig struct {
	beego.Controller
}

type RubyConfigController struct {
	beego.Controller
}

func (c *RubyConfigController) GetRubyConfig() {
	var config models.Rubyconfig

	config.Get()

	c.Data["json"] = &config
	c.ServeJSON()
}
