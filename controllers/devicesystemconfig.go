package controllers

import (
	"github.com/alberthw/ruby/models"
	"github.com/alberthw/ruby/mongoose"
	"github.com/astaxie/beego"
)

type DeviceSystemConfigController struct {
	beego.Controller
}

func (c DeviceSystemConfigController) SetSysConfig() {
	var sysconfig models.Devicesystemconfig
	sysconfig.Devicename = "Ruby"
	sysconfig.Systemversion = "NA"
	sysconfig.Devicesku = "NA"
	sysconfig.Serialnumber = "NA"
	sysconfig.Softwarebuild = "NA"
	sysconfig.Partnumber = "NA"
	sysconfig.Hardwareversion = "NA"

	sysconfig.Country = 1
	sysconfig.Region = 1

	mongoose.WriteDeviceSystemConfig(sysconfig)

	result := "ok"
	c.Data["json"] = &result
	c.ServeJSON()
}
