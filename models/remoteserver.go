package models

import (
	"log"
	"time"

	"github.com/astaxie/beego/orm"
)

type RemoteServer struct {
	Id           int64 `orm:"pk;auto"`
	Remoteserver string
	Isconnected  bool
	Created      time.Time `orm:"auto_now_add;type(datetime)"`
	Updated      time.Time `orm:"auto_now;type(datetime)"`
}

func (c RemoteServer) Get() RemoteServer {
	o := orm.NewOrm()
	var result RemoteServer

	err := o.QueryTable("remoteServer").One(&result)
	if err == orm.ErrNoRows {
		o.Insert(&result)
	}
	return result
}

func (c *RemoteServer) Update() error {
	c.Updated = time.Now()
	o := orm.NewOrm()
	o.Begin()
	_, err := o.Update(c, "Remoteserver", "Isconnected", "Updated")
	if err != nil {
		log.Println(err.Error())
		o.Rollback()
		return err
	}
	o.Commit()
	return err
}
