package routers

import (
	"github.com/alberthw/ruby/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/device", &controllers.DeviceController{})
	beego.Router("/log", &controllers.LogController{})
	beego.Router("/calibration", &controllers.CalibrationController{})
	beego.Router("/command", &controllers.CommandController{})

	beego.Router("/config", &controllers.RubyConfigController{})

	beego.Router("/reposetting", &controllers.RepoSettingController{})
	beego.Router("/testremoteserver", &controllers.RepoSettingController{}, "GET:Test")

	beego.Router("/getfilerepo", &controllers.FileRepoController{}, "POST:GetFiles")
	beego.Router("/downloadfile", &controllers.FileRepoController{}, "POST:DownloadFile")
	beego.Router("/burnhostimage", &controllers.FileRepoController{}, "POST:BurnHostImage")

	beego.Router("/openserial", &controllers.SerialController{}, "POST:Open")
	beego.Router("/closeserial", &controllers.SerialController{}, "GET:Close")
	beego.Router("/sendcommand", &controllers.SerialController{}, "POST:Post")

	beego.Router("/validateconfig", &controllers.DeviceController{}, "GET:ValidateConfig")

	beego.Router("/setsysconfig", &controllers.DeviceSystemConfigController{}, "POST:SetSysConfig")
	beego.Router("/getsysconfig", &controllers.DeviceSystemConfigController{}, "POST:GetSysConfig")

	beego.Router("/sethwconfig", &controllers.DeviceHardwareConfigController{}, "POST:SetHwConfig")
	beego.Router("/gethwconfig", &controllers.DeviceHardwareConfigController{}, "POST:GetHwConfig")

	//	beego.Router("/setswconfig", &controllers.DeviceSoftwareConfigController{}, "POST:SetSwConfig")
	beego.Router("/getswconfig", &controllers.DeviceSoftwareConfigController{}, "POST:GetSwConfig")

	beego.Router("/getreceivecommands", &controllers.CommandController{}, "POST:GetReceiveCommands")
	beego.Router("/getversion", &controllers.CommandController{}, "GET:GetVersions")
	beego.Router("/getlastknownversion", &controllers.CommandController{}, "GET:GetLastKnownVersions")
	beego.Router("/enterservicemode", &controllers.CommandController{}, "GET:EnterServiceMode")
	beego.Router("/exitservicemode", &controllers.CommandController{}, "GET:ExitServiceMode")

	beego.Router("/getdevicelog", &controllers.LogController{}, "POST:GetDeviceLog")

	beego.Router("/startcalibration", &controllers.CalibrationController{}, "GET:StartCalibration")
	beego.Router("/printcalibration", &controllers.CalibrationController{}, "GET:PrintCalibrationData")
	beego.Router("/setcalibration", &controllers.CalibrationController{}, "POST:SetCalibrationRMS")

}
