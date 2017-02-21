package controllers

import (
	"time"

	"log"

	"github.com/alberthw/ruby/mongoose"
	"github.com/astaxie/beego"
)

type CalibrationController struct {
	beego.Controller
}

func (c *CalibrationController) Get() {
	c.TplName = "calibration.html"
	c.Layout = "layout.html"
}

func (c *CalibrationController) PrintCalibrationData() {
	mongoose.SendEnterSeviceMode()
	time.Sleep(time.Second)
	mongoose.SendPrintCalibrationData()
	result := "ok"
	c.Data["json"] = &result
	c.ServeJSON()
}

func (c *CalibrationController) StartCalibration() {
	mongoose.SendEnterSeviceMode()
	time.Sleep(time.Second)
	mongoose.SendStartCalibration()
	result := "ok"
	c.Data["json"] = &result
	c.ServeJSON()
}

func (c *CalibrationController) SetCalibrationRMS() {
	mongoose.SendEnterSeviceMode()
	time.Sleep(time.Second)
	rmsValue := c.GetString("rms")
	log.Printf("set calibration data : %s\n", rmsValue)
	mongoose.SendCalibratedRMS(rmsValue)
	time.Sleep(time.Second)
	mongoose.SendCalibrateSetIrms()
	result := "ok"
	c.Data["json"] = &result
	c.ServeJSON()
}
