package controllers

import (
	"strconv"

	"github.com/alberthw/ruby/models"
	"github.com/astaxie/beego"
)

type RemoteServerController struct {
	beego.Controller
}

func (c RemoteServerController) Get() {
	var config models.RemoteServer
	row := config.Get()
	c.Data["json"] = &row
	c.ServeJSON()
}

func (c RemoteServerController) Post() {
	var config models.RemoteServer
	config.Id, _ = c.GetInt64("Id")
	config.Remoteserver = c.GetString("Remoteserver")
	config.Isconnected, _ = c.GetBool("Isconnected")
	config.Update()
	c.Ctx.WriteString(strconv.FormatInt(config.Id, 10))
}
