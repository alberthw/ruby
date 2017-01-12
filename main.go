package main

import (
	"log"
	"time"

	"github.com/alberthw/ruby/models"
	_ "github.com/alberthw/ruby/routers"
	"github.com/alberthw/ruby/serial"
	"github.com/astaxie/beego"
)

func main() {

	beego.BConfig.Listen.HTTPPort = 8089

	//	beego.SetStaticPath("/release", "release")

	go open()
	go syncReleaseFileRepository(6000)
	go reader(100)

	beego.Run()

	//	pwd, _ := os.Getwd()
	//	go syncReleaseFolder(remoteFileRepoFolder, pwd+"/static/release")

}

func reader(t time.Duration) {
	for {
		b, _ := serial.Reader()
		if len(b) == 0 {
			continue
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

		models.GConfig = models.GConfig.Get()
		//		log.Printf("serial name : %s, serial baud : %d, connection status : %v", models.GConfig.Serialname, models.GConfig.Serialbaud, models.GConfig.Isconnected)

		connected := false
		err := serial.Open(models.GConfig.Serialname, int(models.GConfig.Serialbaud))

		if err != nil {
			log.Println(err.Error())
			index++
		} else {
			connected = true
			index = 0
			waitTime = 1000
		}

		models.GConfig.Isconnected = connected
		models.GConfig.UpdateStatus()
		if index > 100 {
			waitTime = 100000
		}
		time.Sleep(time.Millisecond * waitTime)
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
