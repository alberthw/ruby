package models

import (
	"encoding/binary"
	"log"
	"time"

	"fmt"

	"github.com/alberthw/ruby/util"
	"github.com/astaxie/beego/orm"
)

type Devicehardwareconfig struct {
	ID           int64     `orm:"pk;auto;column(id)"`
	Name         string    `orm:"size(20)"`
	PartNumber   string    `orm:"size(20);column(partnumber)"`
	Revision     string    `orm:"size(20)"`
	SerialNumber string    `orm:"size(20);column(serialnumber)"`
	Updated      time.Time `orm:"type(datetime);null"`
	Created      time.Time `orm:"auto_now_add;type(datetime)"`
}

func (c *Devicehardwareconfig) init() {
	c.Name = "NA"
	c.PartNumber = "NA"
	c.Revision = "NA"
	c.SerialNumber = "NA"
}

func GetDeviceHardwareConfig() Devicehardwareconfig {
	o := orm.NewOrm()
	var result Devicehardwareconfig

	err := o.QueryTable("Devicehardwareconfig").One(&result)
	if err == orm.ErrNoRows {
		result.init()
		o.Insert(&result)
	}
	return result
}

func (c Devicehardwareconfig) ToByte() []byte {

	result := make([]byte, ConfigRecordSize)
	for i, _ := range result {
		result[i] = 0xFF
	}

	copy(result[:20], StringToByteArray(c.Name, 20))
	copy(result[20:40], StringToByteArray(c.PartNumber, 20))
	copy(result[40:60], StringToByteArray(c.Revision, 20))
	copy(result[60:80], StringToByteArray(c.SerialNumber, 20))

	fmt.Printf("%X\n", result)

	crc := util.Crc16Byte2(result[:128])
	fmt.Printf("%X\n", crc)
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, crc)

	fmt.Printf("%X\n", buf)

	copy(result[128:], buf)
	return result[:]
}

func (c *Devicehardwareconfig) Insert() {
	o := orm.NewOrm()
	o.Begin()

	id, err := o.Insert(c)
	if err != nil {
		log.Println(err.Error())
		o.Rollback()
	} else {
		//		log.Println(id)
		c.ID = id
	}
	o.Commit()
}

func (c *Devicehardwareconfig) Update() error {
	c.Updated = time.Now()
	o := orm.NewOrm()
	o.Begin()
	_, err := o.Update(c, "name", "PartNumber", "Revision", "SerialNumber", "Updated")
	if err != nil {
		log.Println(err.Error())
		o.Rollback()
		return err
	}
	o.Commit()
	return err
}
