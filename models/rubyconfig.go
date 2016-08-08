package models

import (
	"github.com/astaxie/beego/orm"
	"log"
	"time"
)

type Rubyconfig struct {
	Id          int64  `orm:"pk;auto"`
	Serialline  string `orm:"unique"`
	Serialspeed int64
	Created     time.Time `orm:"auto_now_add;type(datetime)"`
	Updated     time.Time `orm:"auto_now;type(datetime)"`
}

func (c *Rubyconfig) Get() error {
	o := orm.NewOrm()
	err := o.Read(c)
	return err
}

func (c *Rubyconfig) Insert() error {
	o := orm.NewOrm()
	o.Begin()

	id, err := o.Insert(c)
	if err != nil {
		log.Println(err.Error())
		o.Rollback()
		return err
	}
	c.Id = id

	o.Commit()
	return nil
}
