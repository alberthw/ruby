package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/alberthw/ruby/models"
	"github.com/astaxie/beego"
)

type RemoteServerController struct {
	beego.Controller
}

func (c RemoteServerController) Get() {
	var config models.Remoteserver
	row := config.Get()
	c.Data["json"] = &row
	c.ServeJSON()
}

func (c RemoteServerController) Post() {
	var config models.Remoteserver
	config.Id, _ = c.GetInt64("Id")
	config.Remoteserver = c.GetString("Remoteserver")
	config.Isconnected, _ = c.GetBool("Isconnected")
	config.Update()
	c.Ctx.WriteString(strconv.FormatInt(config.Id, 10))
}

func (c RemoteServerController) Test() {
	var config models.Remoteserver
	config = config.Get()
	url := "http://" + config.Remoteserver + "/userContent/Release/"
	resp, err := http.Get(url)
	var result string
	if err != nil {
		log.Println("RemoteServerController::Test(): ", err.Error())
		result = "failed"
	} else {
		result = strconv.Itoa(resp.StatusCode)
	}
	c.Data["json"] = &result
	c.ServeJSON()
}
