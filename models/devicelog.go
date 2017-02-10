package models

import (
	"errors"
	"log"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type DeviceLogType int64

const (
	ERROR DeviceLogType = 0
	EVENT DeviceLogType = 1
)

type DeviceLog struct {
	ID       int64         `orm:"pk;auto;column(id)"`
	LogType  DeviceLogType `orm:"column(logtype)"`
	DataType int64         `orm:"column(datatype)"`
	Content  string
	Created  time.Time `orm:"unique;type(datetime)"`
}

func (c DeviceLog) TableName() string {
	return "devicelog"
}

func (c *DeviceLog) Insert() {
	o := orm.NewOrm()
	o.Begin()

	id, err := o.Insert(c)
	if err != nil {
		log.Println(err.Error())
		o.Rollback()
	} else {
		//		log.Println(id)
		c.ID = id
	}

	o.Commit()
}

func GetDeviceLog(t DeviceLogType, limit int64) ([]DeviceLog, error) {
	o := orm.NewOrm()
	var result []DeviceLog
	_, err := o.Raw("select id,content,created from (select id,content,created from devicelog where logtype = ? order by id DESC limit ?) t order by id ASC", t, limit).QueryRows(&result)
	return result, err
}

func (c *DeviceLog) ParseContent() error {
	lst := strings.Split(c.Content, ",")
	if len(lst) < 3 {
		e := "invalid log content : " + c.Content
		return errors.New(e)
	}
	const form = "02/01/06 15:04:05.000"
	c.Created, _ = time.Parse(form, strings.Join(lst[:2], " "))
	c.Content = strings.Join(lst[2:], " ")
	return nil
}

/*

func GetSendCommands(limit int64) ([]Command, error) {
	o := orm.NewOrm()
	var result []Command
	_, err := o.Raw("select id,info from (select id,info from command where commandtype = ? order by id DESC limit ?) t order by id ASC", SEND, limit).QueryRows(&result)
	return result, err
}

func GetReceiveCommands(limit int64) ([]Command, error) {
	o := orm.NewOrm()
	var result []Command
	_, err := o.Raw("select id,info from (select id,info from command where commandtype = ? order by id DESC limit ?) t order by id ASC", RECEIVE, limit).QueryRows(&result)
	return result, err
}
*/
