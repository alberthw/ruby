package models

import (
	"log"
	"time"

	"github.com/astaxie/beego/orm"
)

type Request struct {
	Id          int64 `orm:"pk;auto"`
	Content     string
	Isprocessed bool
	Created     time.Time `orm:"auto_now_add;type(datetime)"`
	Updated     time.Time `orm:"auto_now;type(datetime)"`
}

func (c Request) Insert() error {
	o := orm.NewOrm()
	_, err := o.Insert(&c)
	return err
}

func (c Request) Get() []Request {
	o := orm.NewOrm()
	var results []Request
	o.QueryTable("request").Filter("Isprocessed", false).All(&results)
	return results
}

func (c *Request) UpdateStatus() error {
	c.Updated = time.Now()
	o := orm.NewOrm()
	o.Begin()
	_, err := o.Update(c, "Isprocessed", "Updated")
	if err != nil {
		log.Println(err.Error())
		o.Rollback()
		return err
	}
	o.Commit()
	return err
}
