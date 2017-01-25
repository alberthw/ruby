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
	ID              int64     `orm:"column(id)"`
	DeviceName      string    `orm:"size(20);column(devicename)"`
	SystemVersion   string    `orm:"size(20);column(systemversion)"`
	DeviceSKU       string    `orm:"size(20);column(devicesku)"`
	SerialNumber    string    `orm:"size(20);column(serialnumber)"`
	SoftwareBuild   string    `orm:"size(10);column(softwarebuild)"`
	PartNumber      string    `orm:"size(10);column(partnumber)"`
	HardwareVersion string    `orm:"size(20);column(hardwareversion)"`
	CRC             uint16    `orm:"column(crc)"`
	Updated         time.Time `orm:"type(datetime);null"`
	Created         time.Time `orm:"auto_now_add;type(datetime)"`
}

func (c *Devicesystemconfig) init() {
	c.DeviceName = "NA"
	c.SystemVersion = "NA"
	c.DeviceSKU = "NA"
	c.SerialNumber = "NA"
	c.SoftwareBuild = "NA"
	c.PartNumber = "NA"
	c.HardwareVersion = "NA"
}

func GetDeviceSystemConfig() Devicesystemconfig {
	o := orm.NewOrm()
	var result Devicesystemconfig

	err := o.QueryTable("Devicesystemconfig").One(&result)
	if err == orm.ErrNoRows {
		result.init()
		o.Insert(&result)
	}
	return result
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

	copy(result[:20], StringToByteArray(c.DeviceName, 20))
	copy(result[20:40], StringToByteArray(c.SystemVersion, 20))
	copy(result[40:60], StringToByteArray(c.DeviceSKU, 20))
	copy(result[60:80], StringToByteArray(c.SerialNumber, 20))
	copy(result[80:90], StringToByteArray(c.SoftwareBuild, 10))
	copy(result[90:100], StringToByteArray(c.PartNumber, 10))
	copy(result[100:120], StringToByteArray(c.HardwareVersion, 20))

	fmt.Printf("%X\n", result)
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
		c.ID = id
	}
	o.Commit()
}

func (c *Devicesystemconfig) Update() error {
	c.Updated = time.Now()
	o := orm.NewOrm()
	o.Begin()
	_, err := o.Update(c, "devicename", "systemversion", "devicesku", "serialnumber", "softwarebuild", "partnumber", "hardwareversion", "updated")
	if err != nil {
		log.Println(err.Error())
		o.Rollback()
		return err
	}
	o.Commit()
	return err
}
