package main

import (
	"encoding/hex"
	"log"
	"time"

	"sync"

	"github.com/alberthw/ruby/ebdprotocol"
	"github.com/alberthw/ruby/models"
	_ "github.com/alberthw/ruby/routers"
	"github.com/alberthw/ruby/serial"
	"github.com/astaxie/beego"
)

var bSoftDelete = false

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

func main() {

	beego.BConfig.Listen.HTTPPort = 8089

	//	beego.SetStaticPath("/release", "release")

	go open()
	go keepAlive(200)
	go writer(100)
	go reader(100)

	go checkBuffer(100)
	go parseBuffer(100)

	beego.Run()

	//	pwd, _ := os.Getwd()
	//	go syncReleaseFolder(remoteFileRepoFolder, pwd+"/static/release")

}

func keepAlive(t time.Duration) {
	for {
		config := models.GConfig.Get()
		if config.Isconnected {
			var k ebdprotocol.KeepAlive
			k.NoAck = true
			k.SessionKey = []byte(config.Sessionkey)

			k.Sequence = byte(config.Sequence[0])

			//			log.Printf("session key : %X, sequence : %X\n", k.SessionKey, k.Sequence)
			/*
				var m models.Message
				m.Messagetype = models.REQUEST
				m.Info = hex.EncodeToString(k.Message())
				m.Status = models.NONE
				m.InsertMessage()
			*/
			if serial.GSerial != nil {
				err := serial.Writer(k.Message())
				if err != nil {
					log.Println(err.Error())
					continue
				}
			}
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
			log.Printf("Send:%X", b)
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

type Buffer struct {
	content []byte
	locker  sync.Mutex
}

func (b Buffer) ReadAll() []byte {
	b.locker.Lock()
	defer b.locker.Unlock()
	return b.content
}
func (b *Buffer) Add(s []byte) {
	b.locker.Lock()
	defer b.locker.Unlock()
	b.content = append(b.content, s...)
}

func (b *Buffer) Parse() []byte {
	b.locker.Lock()
	defer b.locker.Unlock()
	if len(b.content) < 2 {
		return nil
	}
	if b.content[0] == ebdprotocol.ACK {

		result := b.content[:2]
		b.content = b.content[2:]
		return result
	}
	return nil
}

func checkBuffer(t time.Duration) {
	for {
		log.Printf("received : 0x%X\n", gBuffer.ReadAll())
		time.Sleep(time.Millisecond * t)
	}

}
func parseBuffer(t time.Duration) {
	for {
		result := gBuffer.Parse()
		if len(result) == 0 {
			continue
		}
		log.Printf("parsed : 0x%X\n", result)
		time.Sleep(time.Millisecond * t)
	}

}

var gBuffer Buffer

func reader(t time.Duration) {
	for {
		b, _ := serial.Reader()
		if len(b) == 0 {
			continue
		}
		log.Printf("received : 0x%X\n", b)
		gBuffer.Add(b)

		time.Sleep(time.Millisecond * t)
	}
}

func open() {
	for {
		models.GConfig = models.GConfig.Get()
		//		log.Printf("serial name : %s, serial baud : %d, connection status : %v", models.GConfig.Serialname, models.GConfig.Serialbaud, models.GConfig.Isconnected)

		connected := false
		err := serial.Open(models.GConfig.Serialname, int(models.GConfig.Serialbaud))

		if err != nil {
			log.Println(err.Error())
		} else {
			connected = true
		}
		models.GConfig.Isconnected = connected
		models.GConfig.UpdateStatus()
		time.Sleep(time.Millisecond * 1000)
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
