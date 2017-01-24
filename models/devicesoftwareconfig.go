package models

import (
	"encoding/binary"
	"fmt"
	"log"
	"time"

	"github.com/alberthw/ruby/util"
	"github.com/astaxie/beego/orm"
)

type SoftwareType int64

const (
	HOSTBOOT SoftwareType = 0
	HOSTAPP  SoftwareType = 1
	DSPAPP   SoftwareType = 2
)

type Devicesoftwareconfig struct {
	Id         int64
	Name       string `orm:"size(20)"`
	Type       SoftwareType
	Partnumber string    `orm:"size(20)"`
	Version    string    `orm:"size(20)"`
	Imagecrc   string    `orm:"size(20)"`
	Updatetime time.Time `orm:"type(datetime);null"`
	Createtime time.Time `orm:"auto_now_add;type(datetime)"`
}

func (c Devicesoftwareconfig) ToByte() []byte {

	result := make([]byte, ConfigRecordSize)
	for i, _ := range result {
		result[i] = 0xFF
	}

	copy(result[:20], StringToByteArray(c.Name, 20))
	copy(result[20:40], StringToByteArray(c.Partnumber, 20))
	copy(result[40:60], StringToByteArray(c.Version, 20))
	copy(result[60:80], StringToByteArray(c.Imagecrc, 20))

	fmt.Printf("%X\n", result)

	crc := util.Crc16Byte2(result[:128])
	fmt.Printf("%X\n", crc)
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, crc)

	fmt.Printf("%X\n", buf)

	copy(result[128:], buf)
	return result[:]
}

func (c *Devicesoftwareconfig) init() {
	c.Name = "NA"
	c.Partnumber = "NA"
	c.Version = "NA"
	c.Imagecrc = "NA"
}

func GetDeviceSoftwareConfig(t SoftwareType) Devicesoftwareconfig {
	o := orm.NewOrm()
	var result Devicesoftwareconfig

	err := o.QueryTable("Devicesoftwareconfig").Filter("Type", t).One(&result)
	if err == orm.ErrNoRows {
		result.init()
		o.Insert(&result)
	}
	return result
}

func (c *Devicesoftwareconfig) Insert() {
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

func (c *Devicesoftwareconfig) Update() error {
	c.Updatetime = time.Now()
	o := orm.NewOrm()
	o.Begin()
	_, err := o.Update(c, "Name", "Type", "Partnumber", "Version", "Imagecrc", "Updatetime")
	if err != nil {
		log.Println(err.Error())
		o.Rollback()
		return err
	}
	o.Commit()
	return err
}
