package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Response struct {
	Id        int64 `orm:"pk;auto"`
	Requestid int64
	Content   string
	Created   time.Time `orm:"auto_now_add;type(datetime)"`
	Updated   time.Time `orm:"auto_now;type(datetime)"`
}

func (c Response) Insert() error {
	o := orm.NewOrm()
	_, err := o.Insert(&c)
	return err
}

func InsertMultiResponse(rs []Response) error {
	o := orm.NewOrm()
	_, err := o.InsertMulti(len(rs), rs)
	return err
}

func (c Response) GetAllResponse() []Response {
	o := orm.NewOrm()
	var resList []Response
	o.QueryTable("response").All(&resList)
	return resList
}

func (c Response) GetResponseList(requestid int64) []Response {
	o := orm.NewOrm()
	var results []Response
	o.QueryTable("response").Filter("Requestid", requestid).All(&results)
	return results
}
