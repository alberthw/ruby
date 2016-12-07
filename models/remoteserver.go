package models

import (
	"log"
	"time"

	"github.com/astaxie/beego/orm"
)

type Remoteserver struct {
	Id            int64 `orm:"pk;auto"`
	Remoteserver  string
	Contentfolder string
	Isconnected   bool
	Created       time.Time `orm:"auto_now_add;type(datetime)"`
	Updated       time.Time `orm:"auto_now;type(datetime)"`
}

func (c Remoteserver) Get() Remoteserver {
	o := orm.NewOrm()
	var result Remoteserver

	err := o.QueryTable("remoteserver").One(&result)
	if err == orm.ErrNoRows {
		o.Insert(&result)
	}
	return result
}

func (c *Remoteserver) Update() error {
	c.Updated = time.Now()
	o := orm.NewOrm()
	o.Begin()
	_, err := o.Update(c, "Remoteserver", "Contentfolder", "Isconnected", "Updated")
	if err != nil {
		log.Println(err.Error())
		o.Rollback()
		return err
	}
	o.Commit()
	return err
}
