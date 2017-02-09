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

	hwconfig.ID, _ = c.GetInt64("id")

	hwconfig.Name = c.GetString("name")
	hwconfig.PartNumber = c.GetString("partNumber")
	hwconfig.Revision = c.GetString("revision")
	hwconfig.SerialNumber = c.GetString("serialNumber")

	mongoose.WriteDeviceHardwareConfig(hwconfig)

	hwconfig.Update()

	result := "ok"
	c.Data["json"] = &result
	c.ServeJSON()
}

func (c DeviceHardwareConfigController) GetHwConfig() {
	tmp, _ := c.GetInt64("block")
	block := models.ConfigBlock(tmp)
	row := models.GetDeviceHardwareConfig(block)
	c.Data["json"] = &row
	c.ServeJSON()
}
