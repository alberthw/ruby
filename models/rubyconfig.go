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

func (c Rubyconfig) Get() Rubyconfig {
	o := orm.NewOrm()
	var result Rubyconfig
	err := o.QueryTable("rubyconfig").One(&result)
	if err == orm.ErrNoRows {
		o.Insert(&result)
	}
	return result
}

func (c *Rubyconfig) Update() error {
	c.Updated = time.Now()
	o := orm.NewOrm()
	o.Begin()
	_, err := o.Update(c, "Serialline", "Serialspeed", "Updated")
	if err != nil {
		log.Println(err.Error())
		o.Rollback()
		return err
	}
	o.Commit()
	return err
}
