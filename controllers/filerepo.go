package controllers

import (
	"strconv"

	"github.com/alberthw/ruby/models"
	"github.com/astaxie/beego"
)

type FileRepoController struct {
	beego.Controller
}

func (c FileRepoController) GetFiles() {
	models.SyncReleaseFilesInfo()

	rows := models.GetALLReleaseFiles()
	c.Data["json"] = &rows
	c.ServeJSON()
}

func (c FileRepoController) Post() {
	c.Ctx.WriteString(strconv.FormatInt(0, 10))
}

func (c FileRepoController) DownloadFile() {
	filepath := c.GetString("filepath")
	c.Data["json"] = &filepath
	c.ServeJSON()

}
