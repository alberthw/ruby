package routers

import (
	"github.com/alberthw/ruby/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/config", &controllers.RubyConfigController{})

	beego.Router("/reposetting", &controllers.RepoSettingController{})
	beego.Router("/testremoteserver", &controllers.RepoSettingController{}, "GET:Test")

	beego.Router("/getfilerepo", &controllers.FileRepoController{}, "GET:GetFiles")
	beego.Router("/downloadfile", &controllers.FileRepoController{}, "POST:DownloadFile")
	beego.Router("/burnhostimage", &controllers.FileRepoController{}, "POST:BurnHostImage")

	beego.Router("/openserial", &controllers.SerialController{}, "POST:Open")
	beego.Router("/closeserial", &controllers.SerialController{}, "GET:Close")
	beego.Router("/command", &controllers.SerialController{}, "POST:Post")

	beego.Router("/getreceivecommands", &controllers.CommandController{}, "POST:GetReceiveCommands")

}
