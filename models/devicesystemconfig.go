package models

import (
	"encoding/binary"
	"log"
	"time"

	"fmt"

	"github.com/alberthw/ruby/util"
	"github.com/astaxie/beego/orm"
)

const (
	ConfigRecordSize = 132
)

type Devicesystemconfig struct {
	Id              int64
	Devicename      string `orm:"size(20)"`
	Systemversion   string `orm:"size(20)"`
	Devicesku       string `orm:"size(20)"`
	Serialnumber    string `orm:"size(20)"`
	Softwarebuild   string `orm:"size(10)"`
	Partnumber      string `orm:"size(10)"`
	Hardwareversion string `orm:"size(20)"`
	Country         uint8
	Region          uint8
	Crc             uint16
	Updatetime      time.Time `orm:"type(datetime);null"`
	Createtime      time.Time `orm:"auto_now_add;type(datetime)"`
}

func StringToByteArray(s string, l int) []byte {
	result := make([]byte, l)
	copy(result, s)
	return result
}

func (c Devicesystemconfig) ToByte() []byte {

	result := make([]byte, ConfigRecordSize)
	for i, _ := range result {
		result[i] = 0xFF
	}

	copy(result[:20], StringToByteArray(c.Devicename, 20))
	copy(result[20:40], StringToByteArray(c.Systemversion, 20))
	copy(result[40:60], StringToByteArray(c.Devicesku, 20))
	copy(result[60:80], StringToByteArray(c.Serialnumber, 20))
	copy(result[80:90], StringToByteArray(c.Softwarebuild, 10))
	copy(result[90:100], StringToByteArray(c.Partnumber, 10))
	copy(result[100:120], StringToByteArray(c.Hardwareversion, 20))

	fmt.Printf("%X\n", result)

	result[120] = byte(c.Country)
	result[124] = byte(c.Region)
	crc := util.Crc16(result[:128])
	fmt.Printf("%X\n", crc)
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, crc)

	fmt.Printf("%X\n", buf)

	copy(result[128:], buf)
	return result[:]
}

func (c *Devicesystemconfig) Insert() {
	o := orm.NewOrm()
	o.Begin()

	id, err := o.Insert(c)
	if err != nil {
		log.Println(err.Error())
		o.Rollback()
	} else {
		//		log.Println(id)
		c.Id = id
	}
	o.Commit()
}
