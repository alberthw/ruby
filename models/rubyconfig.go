package models

import (
	"log"
	"time"

	"github.com/astaxie/beego/orm"
)

var (
	GConfig Rubyconfig
)

type DeviceID int64

const (
	ForceTraid        DeviceID = 0x00
	ValleylabExchange DeviceID = 0x01
	PatriotGenerator  DeviceID = 0x02
	IntegratedOR      DeviceID = 0x3D
	ServiceApps       DeviceID = 0xD8
)

type Rubyconfig struct {
	Id             int64  `orm:"pk;auto"`
	Serialname     string `orm:"unique"`
	Serialbaud     int64  `orm:"default(115200)"`
	Isconnected    bool
	Deviceid       DeviceID
	Protocolver    string
	Sessionkey     string
	Sequence       string `orm:"default(0)"`
	Writeinterval  int
	Sessionstatus  uint32
	Sessiontimeout uint32
	Messagetimeout uint32
	Maxretrycount  uint32
	Devicename     string
	Created        time.Time `orm:"auto_now_add;type(datetime)"`
	Updated        time.Time `orm:"auto_now;type(datetime)"`
}

func (c *Rubyconfig) init() {
	c.Serialbaud = 115200
	c.Sequence = "0"
}

func (c Rubyconfig) Get() Rubyconfig {
	o := orm.NewOrm()
	var result Rubyconfig

	err := o.QueryTable("rubyconfig").One(&result)
	if err == orm.ErrNoRows {
		result.init()
		o.Insert(&result)
	}
	if len(result.Sequence) == 0 {
		result.Sequence = "0"
	}
	result.Serialbaud = 115200
	return result
}

func (c *Rubyconfig) Update() error {
	c.Updated = time.Now()
	o := orm.NewOrm()
	o.Begin()
	_, err := o.Update(c, "Serialname", "Updated")
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

func (c *Rubyconfig) UpdateSequence() error {
	c.Updated = time.Now()
	o := orm.NewOrm()
	o.Begin()
	_, err := o.Update(c, "Sequence", "Updated")
	if err != nil {
		log.Println(err.Error())
		o.Rollback()
		return err
	}
	o.Commit()
	return err
}
