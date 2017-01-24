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

	sysconfig.Id, _ = c.GetInt64("id")

	sysconfig.Devicename = c.GetString("deviceName")
	sysconfig.Systemversion = c.GetString("sysVersion")
	sysconfig.Devicesku = c.GetString("deviceSKU")
	sysconfig.Serialnumber = c.GetString("serialNumber")
	sysconfig.Softwarebuild = c.GetString("softwareBuild")
	sysconfig.Partnumber = c.GetString("partNumber")
	sysconfig.Hardwareversion = c.GetString("hardwareVersion")

	sysconfig.Country, _ = c.GetUint8("country")
	sysconfig.Region, _ = c.GetUint8("region")

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
