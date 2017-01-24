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

	hwconfig.Id, _ = c.GetInt64("id")

	hwconfig.Name = c.GetString("name")
	hwconfig.Partnumber = c.GetString("partNumber")
	hwconfig.Revision = c.GetString("revision")
	hwconfig.Serialnumber = c.GetString("serialNumber")

	mongoose.WriteDeviceHardwareConfig(hwconfig)

	hwconfig.Update()

	result := "ok"
	c.Data["json"] = &result
	c.ServeJSON()
}

func (c DeviceHardwareConfigController) GetSysConfig() {
	row := models.GetDeviceHardwareConfig()
	c.Data["json"] = &row
	c.ServeJSON()
}
