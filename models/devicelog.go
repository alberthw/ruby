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

type DeviceLogResult struct {
	Total int64
	Data  []DeviceLog
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

func GetDeviceLog(filter []int64, limit int64, offset int64, order string) (DeviceLogResult, error) {
	o := orm.NewOrm()
	var result DeviceLogResult
	qs := o.QueryTable("Devicelog")

	if len(filter) == 1 {
		qs = qs.Filter("Logtype", DeviceLogType(filter[0]))
	}
	result.Total, _ = qs.Count()
	if result.Total == 0 {
		return result, nil
	}

	qs = qs.Limit(limit)
	qs = qs.Offset(offset)

	orderStr := "Created"
	if order != "descend" {
		orderStr = "-Created"
	}
	qs = qs.OrderBy(orderStr)

	//	o.QueryTable("Devicelog").Filter("Logtype", t).Limit(limit).Offset(offset).OrderBy("-Created").All(&result)

	qs.All(&result.Data)
	return result, nil
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
