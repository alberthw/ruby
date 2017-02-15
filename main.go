package main

import (
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/alberthw/ruby/models"
	_ "github.com/alberthw/ruby/routers"
	"github.com/alberthw/ruby/serial"
	"github.com/astaxie/beego"
)

var buffer chan []byte

func main() {
	beego.BConfig.Listen.HTTPPort = 8089

	//	beego.SetStaticPath("/repository", "repository")

	buffer = make(chan []byte)

	go open()
	go syncReleaseFileRepository(6000)
	go reader(5)

	go parseMessage()

	beego.Run()

	//	pwd, _ := os.Getwd()
	//	go syncReleaseFolder(remoteFileRepoFolder, pwd+"/static/release")

}

func reader(t time.Duration) {
	go func(t time.Duration) {
		for {
			var tmp []byte
			bGetMsg := false
			b, _ := serial.Reader()
			if len(b) == 0 {
				continue
			}
			for _, v := range b {
				if bGetMsg {
					if v != 0x02 && v != 0x03 {
						tmp = append(tmp, v)
					}
				}
				if v == 0x02 {
					bGetMsg = true
				}
				if v == 0x03 {
					bGetMsg = false
					buffer <- tmp
					tmp = nil
				}
			}

			/*
				go func(b []byte) {
					log.Printf("received : %s\n", b)
					var c models.Command
					c.CommandType = models.RECEIVE
					c.Info = string(b)
					c.InsertCommand()
				}(b)
			*/
			time.Sleep(time.Millisecond * t)
		}

	}(t)

}

func syncReleaseFileRepository(t time.Duration) {
	go func() {
		for {
			models.SyncReleaseFilesInfo()
			time.Sleep(time.Millisecond * t)
		}

	}()

}

func open() {
	index := 0
	var waitTime time.Duration = 1000
	for {

		cfg := models.GetRubyconfig()
		//		log.Printf("serial name : %s, serial baud : %d, connection status : %v", models.GConfig.Serialname, models.GConfig.Serialbaud, models.GConfig.Isconnected)

		connected := false
		err := serial.Open(cfg.SerialName, int(cfg.SerialBaud))

		if err != nil {
			log.Println(err.Error())
			index++
		} else {
			connected = true
			index = 0
			waitTime = 1000
		}

		cfg.IsConnected = connected
		cfg.UpdateSerialConnectionStatus()
		if index > 100 {
			waitTime = 100000
		}
		time.Sleep(time.Millisecond * waitTime)
	}
}

type SysConfig struct {
	Sysconfig models.Devicesystemconfig
}

type HwConfig struct {
	Hwconfig models.Devicehardwareconfig
}

type SwConfig struct {
	Swconfig models.Devicesoftwareconfig
}

type ConfigValidate struct {
	ConfigValidate models.Rubyconfig
}

type DeviceLog struct {
	DeviceLog models.DeviceLog
}

func parseMessage() {
	go func() {
		for {

			tmp := <-buffer
			log.Println("\n-----------START------------")
			if strings.Contains(string(tmp), "{\"sysconfig\"") {
				var f SysConfig
				err := json.Unmarshal(tmp, &f)
				if err != nil {
					log.Println(err.Error())
				} else {
					cfg := models.GetDeviceSystemConfig(f.Sysconfig.Block)
					f.Sysconfig.ID = cfg.ID
					f.Sysconfig.Update()
				}
				log.Println(f)
			}
			if strings.Contains(string(tmp), "{\"hwconfig\"") {
				var f HwConfig
				err := json.Unmarshal(tmp, &f)
				if err != nil {
					log.Println(err.Error())
				} else {
					cfg := models.GetDeviceHardwareConfig(f.Hwconfig.Block)
					f.Hwconfig.ID = cfg.ID
					f.Hwconfig.Update()
				}
				log.Println(f)
			}

			if strings.Contains(string(tmp), "{\"swconfig\"") {
				var f SwConfig
				err := json.Unmarshal(tmp, &f)
				if err != nil {
					log.Println(err.Error())
				} else {
					var t models.SoftwareType
					if strings.Contains(f.Swconfig.Name, "Host Boot") {
						t = models.HOSTBOOT
					}
					if strings.Contains(f.Swconfig.Name, "Host Application") {
						t = models.HOSTAPP
					}
					if strings.Contains(f.Swconfig.Name, "DSP Application") {
						t = models.DSPAPP
					}
					cfg := models.GetDeviceSoftwareConfig(t, f.Swconfig.Block)
					f.Swconfig.ID = cfg.ID
					f.Swconfig.Type = t
					f.Swconfig.Update()
				}
				log.Println(f)
			}

			if strings.Contains(string(tmp), "{\"configvalidate\"") {
				var f ConfigValidate
				err := json.Unmarshal(tmp, &f)
				if err != nil {
					log.Println(err.Error())
				} else {
					log.Println("result:", f.ConfigValidate)
					setting := models.GetRubyconfig()
					setting.IsConfigValidated = f.ConfigValidate.IsConfigValidated
					setting.UpdateConfigValidateStatus()
				}
				log.Println(f)
			}

			if strings.Contains(string(tmp), "{\"devicelog\"") {
				go func() {
					var f DeviceLog
					err := json.Unmarshal(tmp, &f)
					if err != nil {
						log.Println(err.Error())
					} else {
						//			log.Println("result:", f.DeviceLog)
						err := f.DeviceLog.ParseContent()
						if err == nil {
							f.DeviceLog.Insert()
						}
					}
					//			log.Println(f)
				}()

			}

			log.Println("\n-----------END------------")
		}

	}()

}
