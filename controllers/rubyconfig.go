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
	row := models.GetRubyconfig()

	c.Data["json"] = &row
	c.ServeJSON()
}

func (c RubyConfigController) Post() {
	var config models.Rubyconfig
	config.ID, _ = c.GetInt64("id")
	config.SerialName = c.GetString("serialName")
	config.SerialBaud, _ = c.GetInt64("serialBaud")
	//	config.Isconnected, _ = c.GetBool("connect")
	config.UpdateSerialName()
	c.Ctx.WriteString(strconv.FormatInt(config.ID, 10))
}
