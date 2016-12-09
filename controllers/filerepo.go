package controllers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"path/filepath"

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

func downloadFromUrl(url string) {
	tokens := strings.Split(url, "/")
	fileName := tokens[len(tokens)-1]

	var setting models.Rubyconfig
	setting = setting.Get()

	fullpath := setting.Localrepo + "/" + fileName
	fmt.Println("Downloading", url, "to", fullpath)

	// TODO: check file existence first with io.IsExist
	output, err := os.Create(fullpath)
	if err != nil {
		fmt.Println("Error while creating", fullpath, "-", err)
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

	var remote models.Remoteserver
	remote = remote.Get()
	log.Println(remote)
	fullURL := "http://" + remote.Remoteserver + "/" + fp

	downloadFromUrl(fullURL)

	result := filepath.Base(fp) + " is downloaded."
	c.Data["json"] = &result
	c.ServeJSON()
}
