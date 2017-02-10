package controllers

import "github.com/astaxie/beego"

type LogController struct {
	beego.Controller
}

func (c *LogController) Get() {
	c.TplName = "log.html"
	c.Layout = "layout.html"
}
