package models

import (
	"log"
	"time"

	"github.com/astaxie/beego/orm"
)

type Rubyconfig struct {
	ID          int64     `orm:"pk;auto;column(id)"`
	SerialName  string    `orm:"unique;column(serialname)"`
	SerialBaud  int64     `orm:"default(115200);column(serialbaud)"`
	IsConnected bool      `orm:"column(isconnected)"`
	DeviceName  string    `orm:"column(devicename)"`
	Created     time.Time `orm:"auto_now_add;type(datetime)"`
	Updated     time.Time `orm:"auto_now;type(datetime)"`
}

func (c *Rubyconfig) init() {
	c.SerialBaud = 115200
}

func GetRubyconfig() Rubyconfig {
	o := orm.NewOrm()
	var result Rubyconfig

	err := o.QueryTable("rubyconfig").One(&result)
	if err == orm.ErrNoRows {
		result.init()
		o.Insert(&result)
	}
	return result
}

func (c *Rubyconfig) UpdateSerialName() error {
	c.Updated = time.Now()
	o := orm.NewOrm()
	o.Begin()
	_, err := o.Update(c, "SerialName", "Updated")
	if err != nil {
		log.Println(err.Error())
		o.Rollback()
		return err
	}
	o.Commit()
	return err
}

func (c *Rubyconfig) UpdateSerialConnectionStatus() error {
	c.Updated = time.Now()
	o := orm.NewOrm()
	o.Begin()
	_, err := o.Update(c, "IsConnected", "Updated")
	if err != nil {
		log.Println(err.Error())
		o.Rollback()
		return err
	}
	o.Commit()
	return err
}
