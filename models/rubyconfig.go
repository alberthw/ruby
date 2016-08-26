package models

import (
	"log"
	"time"

	"github.com/astaxie/beego/orm"
)

var (
	GConfig Rubyconfig
)

type Rubyconfig struct {
	Id          int64  `orm:"pk;auto"`
	Serialname  string `orm:"unique"`
	Serialbaud  int64
	Isconnected bool
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
	_, err := o.Update(c, "Serialname", "Serialbaud", "Updated")
	if err != nil {
		log.Println(err.Error())
		o.Rollback()
		return err
	}
	o.Commit()
	return err
}

func (c *Rubyconfig) UpdateStatus() error {
	c.Updated = time.Now()
	o := orm.NewOrm()
	o.Begin()
	_, err := o.Update(c, "Isconnected", "Updated")
	if err != nil {
		log.Println(err.Error())
		o.Rollback()
		return err
	}
	o.Commit()
	return err

}
