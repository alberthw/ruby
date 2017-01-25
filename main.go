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
	go reader(300)

	go parseMessage()

	beego.Run()

	//	pwd, _ := os.Getwd()
	//	go syncReleaseFolder(remoteFileRepoFolder, pwd+"/static/release")

}

func reader(t time.Duration) {
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
		log.Printf("received : %s\n", b)
		var c models.Command
		c.Commandtype = models.RECEIVE
		c.Info = string(b)
		c.InsertCommand()
		time.Sleep(time.Millisecond * t)
	}
}

func syncReleaseFileRepository(t time.Duration) {
	for {
		models.SyncReleaseFilesInfo()
		time.Sleep(time.Millisecond * t)
	}
}

func open() {
	index := 0
	var waitTime time.Duration = 1000
	for {

		cfg := models.GetRubyconfig()
		//		log.Printf("serial name : %s, serial baud : %d, connection status : %v", models.GConfig.Serialname, models.GConfig.Serialbaud, models.GConfig.Isconnected)

		connected := false
		err := serial.Open(cfg.Serialname, int(cfg.Serialbaud))

		if err != nil {
			log.Println(err.Error())
			index++
		} else {
			connected = true
			index = 0
			waitTime = 1000
		}

		cfg.Isconnected = connected
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

func parseMessage() {
	for {
		tmp := <-buffer

		if strings.Contains(string(tmp), "{\"sysconfig\"") {
			var f SysConfig
			err := json.Unmarshal(tmp, &f)
			if err != nil {
				log.Println(err.Error())
			} else {
				cfg := models.GetDeviceSystemConfig()
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
				cfg := models.GetDeviceHardwareConfig()
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
				cfg := models.GetDeviceSoftwareConfig(t)
				f.Swconfig.ID = cfg.ID
				f.Swconfig.Type = t
				f.Swconfig.Update()
			}
			log.Println(f)
		}
		log.Println("\n-----------END------------\n")
	}

}

/*
func copyFile(src, dst string) error {
	log.Printf("copy from %s to %s\n", src, dst)
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return err
	}
	err = out.Sync()
	return err
}

const remoteFileRepoFolder = "C:/tools/Jenkins/userContent/Release"

func getFilesInFolder(folder string) []string {
	var result []string
	//	log.Println(folder)
	filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		//	log.Println("path:", path, ",info:", info)
		if info == nil {
			return nil
		}
		if !info.IsDir() {
			result = append(result, path)
		}
		return nil
	})
	return result
}

func searchStrFromArray(s string, arr []string) bool {
	c := strings.Join(arr, ",")
	return strings.Contains(c, s)
}

func syncReleaseFolder(source, dst string) {
	if _, err := os.Stat(source); os.IsNotExist(err) {
		return
	}
	for {
		srcfiles := getFilesInFolder(source)
		dstfiles := getFilesInFolder(dst)
		for _, f := range dstfiles {
			filename := filepath.Base(f)
			if !searchStrFromArray(filename, srcfiles) {
				os.Remove(f)
				var fr models.Filerepo
				fr.DeleteByFilename(filename)
			}
		}
		for _, f := range srcfiles {
			filename := filepath.Base(f)
			localfile := dst + "/" + filename
			if _, err := os.Stat(localfile); os.IsNotExist(err) {
				copyFile(f, localfile)
				var fr models.Filerepo
				fr.Filepath = localfile
				fr.GetFileInfo()
				fr.CreateOrUpdate()
			}
		}
		runtime.Gosched()
	}
}
*/
