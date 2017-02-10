package routers

import (
	"github.com/alberthw/ruby/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/device", &controllers.DeviceController{})
	beego.Router("/log", &controllers.LogController{})

	beego.Router("/config", &controllers.RubyConfigController{})

	beego.Router("/reposetting", &controllers.RepoSettingController{})
	beego.Router("/testremoteserver", &controllers.RepoSettingController{}, "GET:Test")

	beego.Router("/getfilerepo", &controllers.FileRepoController{}, "GET:GetFiles")
	beego.Router("/downloadfile", &controllers.FileRepoController{}, "POST:DownloadFile")
	beego.Router("/burnhostimage", &controllers.FileRepoController{}, "POST:BurnHostImage")

	beego.Router("/openserial", &controllers.SerialController{}, "POST:Open")
	beego.Router("/closeserial", &controllers.SerialController{}, "GET:Close")
	beego.Router("/command", &controllers.SerialController{}, "POST:Post")

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

}
