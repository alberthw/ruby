package controllers

import (
	"fmt"

	"github.com/alberthw/ruby/models"
	"github.com/alberthw/ruby/mongoose"
	"github.com/astaxie/beego"
)

type DeviceSystemConfigController struct {
	beego.Controller
}

func (c DeviceSystemConfigController) SetSysConfig() {
	var sysconfig models.Devicesystemconfig

	sysconfig.ID, _ = c.GetInt64("id")

	sysconfig.DeviceName = c.GetString("deviceName")
	sysconfig.SystemVersion = c.GetString("sysVersion")
	sysconfig.DeviceSKU = c.GetString("deviceSKU")
	sysconfig.SerialNumber = c.GetString("serialNumber")
	sysconfig.SoftwareBuild = c.GetString("softwareBuild")
	sysconfig.PartNumber = c.GetString("partNumber")
	sysconfig.HardwareVersion = c.GetString("hardwareVersion")

	fmt.Println("system config : ", sysconfig)

	mongoose.WriteDeviceSystemConfig(sysconfig)

	sysconfig.Update()

	result := "ok"
	c.Data["json"] = &result
	c.ServeJSON()
}

func (c DeviceSystemConfigController) GetSysConfig() {
	row := models.GetDeviceSystemConfig()
	c.Data["json"] = &row
	c.ServeJSON()
}
