package models

import (
	"log"
	"time"

	"github.com/astaxie/beego/orm"
)

type Reposetting struct {
	Id           int64 `orm:"pk;auto"`
	Remoteserver string
	Remotefolder string
	Localfolder  string
	Isconnected  bool
	Created      time.Time `orm:"auto_now_add;type(datetime)"`
	Updated      time.Time `orm:"auto_now;type(datetime)"`
}

func (c *Reposetting) init() {
	c.Remotefolder = "/userContent/Release"
	c.Localfolder = "/repository"
}

func GetRepoSetting() Reposetting {
	o := orm.NewOrm()
	var result Reposetting

	err := o.QueryTable("Reposetting").One(&result)
	if err == orm.ErrNoRows {
		result.init()
		o.Insert(&result)
	}
	return result
}

func (c *Reposetting) Update() error {
	c.Updated = time.Now()
	o := orm.NewOrm()
	o.Begin()
	_, err := o.Update(c, "Remoteserver", "Remotefolder", "Localfolder", "Isconnected", "Updated")
	if err != nil {
		log.Println(err.Error())
		o.Rollback()
		return err
	}
	o.Commit()
	return err
}

func (c *Reposetting) UpdateRemoteConnectionStatus() error {
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
