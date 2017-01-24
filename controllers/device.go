package controllers

import "github.com/astaxie/beego"

type DeviceController struct {
	beego.Controller
}

func (c *DeviceController) Get() {
	c.TplName = "device.html"
	c.Layout = "layout.html"
}
