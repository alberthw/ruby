package models

import (
	"log"
	"time"

	"github.com/astaxie/beego/orm"
)

type CommandType int64

const (
	SEND    CommandType = 0
	RECEIVE CommandType = 1
)

type Command struct {
	Id          int64
	Commandtype CommandType
	Info        string
	Updatetime  time.Time `orm:"type(datetime);null"`
	Createtime  time.Time `orm:"auto_now_add;type(datetime)"`
}

func (c *Command) InsertCommand() {
	o := orm.NewOrm()
	o.Begin()

	id, err := o.Insert(c)
	if err != nil {
		log.Println(err.Error())
		o.Rollback()
	} else {
		//		log.Println(id)
		c.Id = id
	}

	o.Commit()
}

func GetSendCommands(limit int64) ([]Command, error) {
	o := orm.NewOrm()
	var result []Command
	_, err := o.QueryTable("command").Filter("Commandtype", SEND).OrderBy("id").Limit(limit).All(&result, "Info")
	return result, err
}

func GetReceiveCommands(limit int64) ([]Command, error) {
	o := orm.NewOrm()
	var result []Command
	_, err := o.QueryTable("command").Filter("Commandtype", RECEIVE).OrderBy("id").Limit(limit).All(&result, "Info")
	return result, err
}
