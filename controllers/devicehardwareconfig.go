package controllers

import (
	"github.com/alberthw/ruby/models"
	"github.com/alberthw/ruby/mongoose"
	"github.com/astaxie/beego"
)

type DeviceHardwareConfigController struct {
	beego.Controller
}

func (c DeviceHardwareConfigController) SetHwConfig() {
	var hwconfig models.Devicehardwareconfig
	hwconfig.Name = "Hardware"
	hwconfig.Partnumber = "1.0"
	hwconfig.Revision = "1.0"
	hwconfig.Serialnumber = "3234"

	mongoose.WriteDeviceHardwareConfig(hwconfig)

	result := "ok"
	c.Data["json"] = &result
	c.ServeJSON()
}
