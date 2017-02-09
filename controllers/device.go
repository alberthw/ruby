package controllers

import "github.com/astaxie/beego"
import "github.com/alberthw/ruby/mongoose"

type DeviceController struct {
	beego.Controller
}

func (c *DeviceController) Get() {
	c.TplName = "device.html"
	c.Layout = "layout.html"
}

func (c DeviceController) ValidateConfig() {
	mongoose.SendValidateConfig()
	result := "ok"
	c.Data["json"] = &result
	c.ServeJSON()
}
