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
	var filters []int64
	c.Ctx.Input.Bind(&filters, "FileType")

	pageNumber, _ := c.GetInt64("page")
	pageSize, _ := c.GetInt64("pageSize")
	//	sortField := c.GetString("sortField")
	sortOrder := c.GetString("sortOrder")
	searchDate := c.GetString("searchDate")

	fmt.Println("filter :", filters)
	fmt.Println("searchDate :", searchDate)
	fmt.Println("page number : ", pageNumber)
	fmt.Println("page size : ", pageSize)
	fmt.Println("sort order : ", sortOrder)

	result, _ := models.GetReleaseFiles(searchDate, filters, pageSize, pageNumber, sortOrder)

	c.Data["json"] = &result
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

	rs := models.GetRepoSetting()
	fullURL := "http://" + rs.Remoteserver + "/" + fp

	pwd, _ := os.Getwd()
	localFolderFullPath := pwd + rs.Localfolder

	if _, err := os.Stat(localFolderFullPath); os.IsNotExist(err) {
		os.Mkdir(localFolderFullPath, os.ModePerm)
	}

	fullpath := localFolderFullPath + "/" + filepath.Base(fp)

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
	ft, _ := c.GetInt64("filetype")
	go mongoose.BurnHostImage(fp, models.FileType(ft))
	result := "ok"
	c.Data["json"] = &result
	c.ServeJSON()

}
