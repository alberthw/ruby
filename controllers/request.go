package controllers

import (
	"github.com/alberthw/ruby/models"
	"github.com/astaxie/beego"
)

type RequestController struct {
	beego.Controller
}

func (c RequestController) Get() {
	var req models.Request
	row := req.Get()
	c.Data["json"] = &row
	c.ServeJSON()
}

func (c RequestController) Post() {
	var req models.Request
	req.Id, _ = c.GetInt64("Id")
	req.Content = c.GetString("Content")
	req.Isprocessed, _ = c.GetBool("Isprocessed")
	err := req.Insert()
	result := "ok"
	if err != nil {
		result = err.Error()
	}
	c.Data["json"] = &result
	c.ServeJSON()
}

func (c RequestController) UpdateStatus() {
	var req models.Request
	req.Id, _ = c.GetInt64("Id")
	req.Isprocessed, _ = c.GetBool("Isprocessed")
	err := req.UpdateStatus()
	result := "ok"
	if err != nil {
		result = err.Error()
	}
	c.Data["json"] = &result
	c.ServeJSON()

}
