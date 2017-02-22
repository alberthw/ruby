package controllers

import (
	"time"

	"log"

	"fmt"

	"github.com/alberthw/ruby/mongoose"
	"github.com/alberthw/ruby/serial"
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
	//	mongoose.SendEnterSeviceMode()
	//	time.Sleep(time.Second)
	mongoose.SendPrintCalibrationData()
	time.Sleep(time.Second)
	result := string(serial.GBuffer)
	serial.GBuffer = nil
	fmt.Println("---------------------------")
	log.Println(string(result))
	fmt.Println("---------------------------")
	c.Data["json"] = &result
	c.ServeJSON()
}

func (c *CalibrationController) StartCalibration() {
	//	mongoose.SendEnterSeviceMode()
	//	time.Sleep(time.Second)
	mongoose.SendStartCalibration()
	time.Sleep(time.Second)
	result := string(serial.GBuffer)
	serial.GBuffer = nil
	c.Data["json"] = &result
	c.ServeJSON()
}

func (c *CalibrationController) SetCalibrationRMS() {
	//	mongoose.SendEnterSeviceMode()
	//	time.Sleep(time.Second)
	rmsValue := c.GetString("rms")
	log.Printf("set calibration data : %s\n", rmsValue)
	mongoose.SendCalibratedRMS(rmsValue)
	time.Sleep(time.Second)
	mongoose.SendCalibrateSetIrms()
	time.Sleep(time.Second)
	mongoose.SendPrintCalibrationData()
	time.Sleep(time.Second * 3)
	result := string(serial.GBuffer)
	serial.GBuffer = nil
	c.Data["json"] = &result
	c.ServeJSON()
}
