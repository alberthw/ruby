package controllers

import (
	"encoding/hex"
	"fmt"

	"strconv"

	"github.com/alberthw/ruby/ebdprotocol"
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

func (c RequestController) Generate() {
	var req models.Message
	t, _ := c.GetInt64("type")
	req.Messagetype = models.MessageType(t)

	result := fmt.Sprintf("ok, get the request : %d", t)

	c.Data["json"] = &result
	c.ServeJSON()
}

func (c RequestController) OpenSession() {
	var req ebdprotocol.RequestSession
	var setting models.Rubyconfig
	setting = setting.Get()
	req.NoAck = true
	pVer, _ := strconv.ParseUint(setting.Protocolver, 10, 32)

	req.SessionKey = []byte(setting.Sessionkey)
	req.Sequence = byte(setting.Sequence[0])

	req.DeviceID = uint32(setting.Deviceid)
	req.ProtocolVersion = uint32(pVer)

	var m models.Message
	m.Messagetype = models.REQUEST
	m.Info = hex.EncodeToString(req.Message())
	m.Status = models.NONE

	m.InsertMessage()

	result := m.Info
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
