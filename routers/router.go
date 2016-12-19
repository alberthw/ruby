package routers

import (
	"github.com/alberthw/ruby/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/config", &controllers.RubyConfigController{})

	beego.Router("/remoteserver", &controllers.RemoteServerController{})
	beego.Router("/testremoteserver", &controllers.RemoteServerController{}, "GET:Test")

	beego.Router("/getfilerepo", &controllers.FileRepoController{}, "GET:GetFiles")
	beego.Router("/downloadfile", &controllers.FileRepoController{}, "POST:DownloadFile")

	beego.Router("/openserial", &controllers.SerialController{}, "POST:Open")
	beego.Router("/closeserial", &controllers.SerialController{}, "GET:Close")

	beego.Router("/command", &controllers.SerialController{}, "POST:Send")

	beego.Router("/generate", &controllers.RequestController{}, "POST:Generate")

	beego.Router("/request", &controllers.RequestController{})

	beego.Router("/response", &controllers.ResponseController{})

	beego.Router("/opensession", &controllers.RequestController{}, "GET:OpenSession")

}
