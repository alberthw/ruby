package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/alberthw/ruby/models"
	"github.com/astaxie/beego"
)

type RepoSettingController struct {
	beego.Controller
}

func (c RepoSettingController) Get() {
	row := models.GetRepoSetting()
	c.Data["json"] = &row
	c.ServeJSON()
}

func (c RepoSettingController) Post() {
	var setting models.Reposetting
	setting.Id, _ = c.GetInt64("Id")
	setting.Remoteserver = c.GetString("Remoteserver")
	//	setting.Remotefolder = c.GetString("Remotefolder")
	//	setting.Localfolder = c.GetString("Localfolder")
	setting.Isconnected, _ = c.GetBool("Isconnected")
	setting.UpdateRemoteConnectionStatus()
	c.Ctx.WriteString(strconv.FormatInt(setting.Id, 10))
}

func (c RepoSettingController) Test() {
	setting := models.GetRepoSetting()
	url := "http://" + setting.Remoteserver + setting.Remotefolder
	resp, err := http.Get(url)
	var result string
	if err != nil {
		log.Println("RepoSettingController::Test(): ", err.Error())
		result = "failed"
	} else {
		result = strconv.Itoa(resp.StatusCode)
	}
	c.Data["json"] = &result
	c.ServeJSON()
}
