package controllers

import (
	"github.com/alberthw/ruby/models"
	"github.com/astaxie/beego"
)

type ResponseController struct {
	beego.Controller
}

func (c ResponseController) Get() {
	var res models.Response
	rows := res.GetAllResponse()
	c.Data["json"] = &rows
	c.ServeJSON()
}

func (c ResponseController) Post() {
	var res models.Response
	res.Requestid, _ = c.GetInt64("requestid")
	res.Content = c.GetString("Content")
	err := res.Insert()
	result := "ok"
	if err != nil {
		result = err.Error()
	}
	c.Data["json"] = &result
	c.ServeJSON()
}
