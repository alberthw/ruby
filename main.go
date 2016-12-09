package main

import (
	"encoding/hex"
	"io"
	"log"
	"os"
	"runtime"
	"strings"
	"time"

	"path/filepath"

	"github.com/alberthw/ruby/ebdprotocol"
	"github.com/alberthw/ruby/models"
	_ "github.com/alberthw/ruby/routers"
	"github.com/alberthw/ruby/serial"
	"github.com/astaxie/beego"
)

var bSoftDelete = true

func IncreaseSeq(seq byte) byte {
	if seq == 0x39 {
		return 0x30
	}
	return seq + 1
}

func IncreaseOneSequence() {
	config := models.GConfig.Get()

	//	log.Printf(" seq  :%x", byte(s.Sequence[0]))
	config.Sequence = string(IncreaseSeq(byte(config.Sequence[0])))
	//	log.Printf("Sequence : %s", s.Sequence)
	config.UpdateSequence()
}

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

func main() {

	beego.SetStaticPath("/release", "release")

	//	go open()
	//	go generate(200)
	//	go writer(100)
	//	go reader(100)

	pwd, _ := os.Getwd()
	go syncReleaseFolder(remoteFileRepoFolder, pwd+"/static/release")
	beego.Run()

}

func generate(t time.Duration) {
	for {
		config := models.GConfig.Get()
		if config.Isconnected {
			var k ebdprotocol.KeepAlive
			k.SessionKey = []byte(config.Sessionkey)

			k.Sequence = byte(config.Sequence[0])

			//			log.Printf("session key : %X, sequence : %X\n", k.SessionKey, k.Sequence)

			var m models.Message
			m.Messagetype = models.REQUEST
			m.Info = hex.EncodeToString(k.Message())
			m.Status = models.NONE

			m.InsertMessage()

		}
		time.Sleep(time.Millisecond * time.Duration(t))
	}
}

func writer(t time.Duration) {
	// get one request
	for {
		config := models.GConfig.Get()
		//		time.Sleep(time.Millisecond * time.Duration(s.Writeinterval))
		time.Sleep(time.Millisecond * time.Duration(t))
		if config.Isconnected {
			var m models.Message
			err := m.GetOneRequest()
			if err != nil {
				//				log.Println("writer:", err.Error())
				continue
			}

			//			log.Printf("get one request :%v", m)

			b, _ := hex.DecodeString(m.Info)
			err = serial.Writer(b)
			if err != nil {
				log.Println(err.Error())
				continue
			}
			//			log.Printf("Send:%X", b)
			IncreaseOneSequence()

			if bSoftDelete {
				m.Status = models.DELETED
				m.UpdateStatus()
			} else {
				m.DeleteMessage()
				//		log.Println("the request is deleted.")
			}
		}

	}
}

type TransferStatus int

const (
	IDLE = iota
	START
	TRANSFERING
	END
)

type Input struct {
	Content []byte
	Status  TransferStatus
}

func (input *Input) Receive(b []byte) {
	if len(b) == 0 {
		return
	}
	if b[0] == ebdprotocol.ACK {
		input.Status = START
	}
	return
}

var buf chan []byte

func reader(t time.Duration) {
	var f ebdprotocol.Frame
	var buf []byte
	for {
		b, _ := serial.Reader()
		if len(b) == 0 {
			continue
		}
		log.Printf("received : 0x%X\n", b)
		_, err := f.Parse(buf)
		if err != nil {
			buf = append(buf, b...)
			log.Println(err.Error())
			log.Printf("buffer : 0x%X\n", buf)
			//		continue
		}
		buf = []byte{}
		/*
			var m models.Message
			b, _ := serial.Reader()
			if len(b) > 0 {
				m.Messagetype = models.RESPONSE
				m.Info = hex.EncodeToString(b)
				m.Status = models.NONE
				//			log.Printf("Received : %X", b)
				//			log.Printf("Received message : %s", m.Info)
				m.InsertMessage()
			}
		*/
		time.Sleep(time.Millisecond * t)
	}
}

func open() {
	for {
		models.GConfig = models.GConfig.Get()
		//		log.Printf("serial name : %s, serial baud : %d, connection status : %v", models.GConfig.Serialname, models.GConfig.Serialbaud, models.GConfig.Isconnected)

		connected := false
		err := serial.Open(models.GConfig.Serialname, int(models.GConfig.Serialbaud))

		if err == nil {
			connected = true
		} else {
			log.Println(err.Error())
		}
		models.GConfig.Isconnected = connected
		models.GConfig.UpdateStatus()
		time.Sleep(time.Millisecond * 1000)
	}
}
