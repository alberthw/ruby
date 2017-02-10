package models

import "time"

type LogType int64

const (
	EVENT LogType = 0
	ERROR LogType = 1
)

type DeviceLog struct {
	ID       int64   `orm:"pk;auto;column(id)"`
	LogType  LogType `orm:"column(logtype)"`
	DataType int64   `orm:"column(datatype)"`
	Content  string
	Created  time.Time `orm:"type(datetime)"`
}

func (log DeviceLog) TableName() string {
	return "devicelog"
}

/*
func (c *Command) InsertCommand() {
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
