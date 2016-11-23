package models

import (
	"log"
	"time"

	"github.com/astaxie/beego/orm"
)

type MessageType int64

const (
	REQUEST  MessageType = 0
	RESPONSE MessageType = 1
)

type MessageStatus int64

const (
	NONE      MessageStatus = 0
	PROCESSED MessageStatus = 1
	INVALID   MessageStatus = 2
	DELETED   MessageStatus = 3
)

type Message struct {
	Id          int64
	Messagetype MessageType
	Info        string
	Status      MessageStatus
	Updatetime  time.Time `orm:"type(datetime);null"`
	Createtime  time.Time `orm:"auto_now_add;type(datetime)"`
}

func (m *Message) InsertMessage() {
	o := orm.NewOrm()
	o.Begin()

	id, err := o.Insert(m)
	if err != nil {
		log.Println(err.Error())
		o.Rollback()
	} else {
		//		log.Println(id)
		m.Id = id
	}

	o.Commit()
}

func (m *Message) Get() error {
	o := orm.NewOrm()
	err := o.Read(m)
	return err

}

func (m *Message) GetOneRequest() error {
	o := orm.NewOrm()

	err := o.QueryTable("message").Filter("messagetype", REQUEST).Filter("status", NONE).OrderBy("id").Limit(1).One(m)

	return err
}

func (m *Message) GetOneResponse() error {
	o := orm.NewOrm()
	err := o.QueryTable("message").Filter("messagetype", RESPONSE).Filter("status", NONE).OrderBy("-id").Limit(1).One(m)

	return err
}

func (m *Message) UpdateStatus() error {
	o := orm.NewOrm()
	_, err := o.Update(m, "Status")
	return err
}

func (m *Message) UpdateInfo() error {
	o := orm.NewOrm()
	_, err := o.Update(m, "Info")
	return err
}

func (m *Message) DeleteMessage() {

	o := orm.NewOrm()

	_, err := o.Delete(m)
	if err != nil {
		log.Println("DeleteMessage:", err.Error())
	}
}
