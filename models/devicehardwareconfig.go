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
	Id           int64
	Name         string    `orm:"size(20)"`
	Partnumber   string    `orm:"size(20)"`
	Revision     string    `orm:"size(20)"`
	Serialnumber string    `orm:"size(20)"`
	Updatetime   time.Time `orm:"type(datetime);null"`
	Createtime   time.Time `orm:"auto_now_add;type(datetime)"`
}

func (c Devicehardwareconfig) ToByte() []byte {

	result := make([]byte, ConfigRecordSize)
	for i, _ := range result {
		result[i] = 0xFF
	}

	copy(result[:20], StringToByteArray(c.Name, 20))
	copy(result[20:40], StringToByteArray(c.Partnumber, 20))
	copy(result[40:60], StringToByteArray(c.Revision, 20))
	copy(result[60:80], StringToByteArray(c.Serialnumber, 20))

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
		c.Id = id
	}
	o.Commit()
}
