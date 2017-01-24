package controllers

import (
	"github.com/alberthw/ruby/models"
	"github.com/alberthw/ruby/mongoose"
	"github.com/astaxie/beego"
)

type DeviceSoftwareConfigController struct {
	beego.Controller
}

func (c DeviceSoftwareConfigController) SetSwConfig() {
	var swconfig models.Devicesoftwareconfig

	swconfig.Id, _ = c.GetInt64("id")
	tmp, _ := c.GetInt64("type")
	swconfig.Type = models.SoftwareType(tmp)

	swconfig.Name = c.GetString("name")
	swconfig.Partnumber = c.GetString("partNumber")
	swconfig.Version = c.GetString("version")
	swconfig.Imagecrc = c.GetString("imageCRC")

	mongoose.WriteDeviceSoftwareConfig(swconfig)

	swconfig.Update()

	result := "ok"
	c.Data["json"] = &result
	c.ServeJSON()
}

func (c DeviceSoftwareConfigController) GetSwConfig() {
	t, _ := c.GetInt64("type")
	row := models.GetDeviceSoftwareConfig(models.SoftwareType(t))
	c.Data["json"] = &row
	c.ServeJSON()
}
