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
	ID         int64  `orm:"pk;auto;column(id)"`
	Name       string `orm:"size(20)"`
	Type       SoftwareType
	PartNumber string    `orm:"size(20);column(partnumber)"`
	Version    string    `orm:"size(20)"`
	ImageCRC   string    `orm:"size(20);column(imagecrc)"`
	Updated    time.Time `orm:"type(datetime);null"`
	Created    time.Time `orm:"auto_now_add;type(datetime)"`
}

func (c Devicesoftwareconfig) ToByte() []byte {

	result := make([]byte, ConfigRecordSize)
	for i, _ := range result {
		result[i] = 0xFF
	}

	copy(result[:20], StringToByteArray(c.Name, 20))
	copy(result[20:40], StringToByteArray(c.PartNumber, 20))
	copy(result[40:60], StringToByteArray(c.Version, 20))
	copy(result[60:80], StringToByteArray(c.ImageCRC, 20))

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
	c.PartNumber = "NA"
	c.Version = "NA"
	c.ImageCRC = "NA"
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
		c.ID = id
	}
	o.Commit()
}

func (c *Devicesoftwareconfig) Update() error {
	c.Updated = time.Now()
	o := orm.NewOrm()
	o.Begin()
	_, err := o.Update(c, "name", "type", "partnumber", "version", "imagecrc", "updated")
	if err != nil {
		log.Println(err.Error())
		o.Rollback()
		return err
	}
	o.Commit()
	return err
}
