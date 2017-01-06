package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"path/filepath"

	"github.com/alberthw/ruby/models"
	"github.com/alberthw/ruby/mongoose"
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

func downloadFromUrl(url string, local string) {

	fmt.Println("Downloading", url, "to", local)

	// TODO: check file existence first with io.IsExist
	output, err := os.Create(local)
	if err != nil {
		fmt.Println("Error while creating", local, "-", err)
		return
	}
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}
	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}

	fmt.Println(n, "bytes downloaded.")
}

func (c FileRepoController) DownloadFile() {
	fp := c.GetString("filepath")
	//	id, _ := c.GetInt64("id")

	var remote models.Remoteserver
	remote = remote.Get()
	fullURL := "http://" + remote.Remoteserver + "/" + fp

	var setting models.Rubyconfig
	setting = setting.Get()

	if _, err := os.Stat(setting.Localrepo); os.IsNotExist(err) {
		os.Mkdir(setting.Localrepo, os.ModePerm)
	}

	fullpath := setting.Localrepo + "/" + filepath.Base(fp)

	downloadFromUrl(fullURL, fullpath)
	result := filepath.Base(fp) + " is downloaded."
	/*
		var f models.Filerepo
		f.Id = id
		f.Filepath = fullpath
		f.GetFileInfo()
	*/
	c.Data["json"] = &result
	c.ServeJSON()
}

func (c FileRepoController) BurnHostImage() {
	fp := c.GetString("filepath")
	t, _ := c.GetInt64("filetype")
	err := mongoose.BurnHostImage(fp, models.FileType(t))
	result := "ok"
	if err != nil {
		result = err.Error()
	}
	c.Data["json"] = &result
	c.ServeJSON()

}
