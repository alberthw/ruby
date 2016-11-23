package controllers

import (
	"github.com/alberthw/ruby/models"
	"github.com/astaxie/beego"
)

type ResponseController struct {
	beego.Controller
}

func (c ResponseController) Get() {
	var msg models.Message
	msg.GetOneResponse()
	c.Data["json"] = &msg
	c.ServeJSON()
}

/*
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
*/
