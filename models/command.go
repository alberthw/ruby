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
	_, err := o.Raw("select id,info from (select id,info from command where commandtype = ? order by id DESC limit ?) t order by id ASC", SEND, limit).QueryRows(&result)
	return result, err
}

func GetReceiveCommands(limit int64) ([]Command, error) {
	o := orm.NewOrm()
	var result []Command
	_, err := o.Raw("select id,info from (select id,info from command where commandtype = ? order by id DESC limit ?) t order by id ASC", RECEIVE, limit).QueryRows(&result)
	return result, err
}
