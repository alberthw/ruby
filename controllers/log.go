package controllers

import (
	"fmt"

	"github.com/alberthw/ruby/models"
	"github.com/astaxie/beego"
)

type LogController struct {
	beego.Controller
}

func (c *LogController) Get() {
	c.TplName = "log.html"
	c.Layout = "layout.html"
}

func (c LogController) GetDeviceLog() {
	var filters []int64
	c.Ctx.Input.Bind(&filters, "LogType")

	pageNumber, _ := c.GetInt64("page")
	pageSize, _ := c.GetInt64("pageSize")
	//	sortField := c.GetString("sortField")
	sortOrder := c.GetString("sortOrder")
	/*
		fmt.Println("filter :", filters)
		fmt.Println("filter1 :", filters1)
		fmt.Println("page number : ", pageNumber)
		fmt.Println("page size : ", pageSize)
		fmt.Println("sort order : ", sortOrder)
	*/
	result, _ := models.GetDeviceLog(filters, pageSize, pageNumber, sortOrder)

	fmt.Println(result)
	c.Data["json"] = &result
	c.ServeJSON()
}
