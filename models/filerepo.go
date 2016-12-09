package models

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"strings"

	"bufio"

	"github.com/astaxie/beego/orm"
)

type FileType int64

const (
	FILETYPE_UNKNOWN FileType = 0
	FILETYPE_APP     FileType = 1
	FILETYPE_BOOT    FileType = 2
	FILETYPE_DSP     FileType = 3
)

type Filerepo struct {
	Id           int64 `orm:"pk;auto"`
	Filename     string
	Crc          string
	Buildnumber  uint64
	Checksum     string
	Filepath     string
	Filetype     FileType
	Isdownloaded bool
	Created      time.Time `orm:"auto_now_add;type(datetime)"`
	Updated      time.Time `orm:"auto_now;type(datetime)"`
}

func caculateChecksum(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println("CaculateChecksum : ", err.Error())
		return ""
	}
	result := fmt.Sprintf("%X", md5.Sum(data))
	return result
}

func (f *Filerepo) checkFileType() {
	filename := strings.ToLower(f.Filename)
	index := strings.Index(filename, "dsp")
	if index > -1 {
		f.Filetype = FILETYPE_DSP
		return
	}
	index = strings.Index(filename, "boot")
	if index > -1 {
		f.Filetype = FILETYPE_BOOT
		return
	}
	index = strings.Index(filename, "host")
	if index > -1 {
		f.Filetype = FILETYPE_APP
		return
	}
	f.Filetype = FILETYPE_UNKNOWN
}

func findCRCLine(filename string, address string) string {
	f, err := os.Open(filename)
	if err != nil {
		log.Println("checkCRC() : ", err.Error())
		return ""
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), address) {
			return scanner.Text()
		}
	}
	if err := scanner.Err(); err != nil {
		log.Println("checkCRC(), scan file failed:", err.Error())
	}
	return ""
}

func findCRC(s string, t FileType) string {
	switch t {
	case FILETYPE_APP, FILETYPE_BOOT:
		if len(s) != 46 {
			return ""
		}
		return s[42:44] + s[40:42] + s[38:40] + s[36:38]
	case FILETYPE_DSP:
		if len(s) != 10 {
			return ""
		}
		return s[8:10] + s[6:8] + s[4:6]
	default:
		return ""
	}
	return ""
}

func (f *Filerepo) getCRC() {
	switch f.Filetype {
	case FILETYPE_APP:
		f.Crc = findCRCLine(f.Filepath, "000BFFF0")
	case FILETYPE_BOOT:
		f.Crc = findCRCLine(f.Filepath, "0003FFF0")
	case FILETYPE_DSP:
		f.Crc = findCRCLine(f.Filepath, "S003")
	case FILETYPE_UNKNOWN:
		// do nothing
	default:
		// do nothing
	}
	f.Crc = findCRC(f.Crc, f.Filetype)
	//	log.Println("CRC:", f.Crc)
}

func getBuildNumberFromFileName(filename string) (uint64, error) {
	if len(filename) < 12 {
		err := errors.New("invalid file name")
		return 0, err
	}
	var result uint64
	bnStr := filename[len(filename)-12 : len(filename)-4]
	result, err := strconv.ParseUint(bnStr, 10, 64)
	if err != nil {
		return 0, err
	}
	return result, nil

}

func getReleaseFilesInfo(folder string) []Filerepo {
	var result []Filerepo
	filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if info == nil {
			return nil
		}
		//		log.Println("file infor", info, " , path:", path)
		if !info.IsDir() {
			buildNumber, err := getBuildNumberFromFileName(info.Name())
			if err != nil {
				return nil
			}
			var fr Filerepo
			fr.Filepath = strings.Replace(path, "\\", "/", -1)
			fr.Filename = info.Name()
			fr.Created = info.ModTime()
			fr.Buildnumber = buildNumber
			fr.checkFileType()
			fr.getCRC()
			fr.Checksum = caculateChecksum(fr.Filepath)

			//		log.Println(fr)
			result = append(result, fr)
		}
		return nil
	})
	return result
}

func clearFileRepo() error {
	o := orm.NewOrm()
	_, err := o.Raw("delete from filerepo").Exec()
	if err != nil {
		log.Println("clearFileRepo():", err.Error())
	}
	return err
}

func (c *Filerepo) checkDownloadStatus() {
	var setting Rubyconfig
	setting = setting.Get()
	localFullPath := setting.Localrepo + "/" + c.Filename
	if _, err := os.Stat(localFullPath); os.IsNotExist(err) {
		c.Isdownloaded = false
		return
	}
	checkSum := caculateChecksum(localFullPath)
	if checkSum == c.Checksum {
		c.Isdownloaded = true
		return
	}
	c.Isdownloaded = false
}

func SyncReleaseFilesInfo() {
	clearFileRepo()
	files := getReleaseFilesInfo("./static/release")
	for _, file := range files {
		file.checkDownloadStatus()
		file.CreateOrUpdate()
	}

	//check local file status
}

func (c Filerepo) Insert() error {
	o := orm.NewOrm()
	_, err := o.Insert(&c)
	return err
}
func GetALLReleaseFiles() []Filerepo {
	o := orm.NewOrm()
	var lists []Filerepo

	o.QueryTable("Filerepo").GroupBy("Filetype", "Id").OrderBy("Filetype", "Buildnumber").All(&lists, "Id", "Filename", "Crc", "Buildnumber", "Filepath", "Filetype", "Isdownloaded")
	return lists
}

func (c *Filerepo) CreateOrUpdate() error {
	o := orm.NewOrm()
	var tmp Filerepo
	err := o.QueryTable("Filerepo").Filter("Checksum", c.Checksum).One(&tmp)
	if err == orm.ErrNoRows {
		return c.Insert()
	}
	c.Id = tmp.Id
	return c.Update()

}

func (c *Filerepo) UpdateDownloadStatus() error {
	c.Updated = time.Now()
	o := orm.NewOrm()
	o.Begin()
	_, err := o.Update(c, "Isdownloaded")
	if err != nil {
		log.Println(err.Error())
		o.Rollback()
		return err
	}
	o.Commit()
	return err

}

func (c *Filerepo) Update() error {
	c.Updated = time.Now()
	o := orm.NewOrm()
	o.Begin()
	_, err := o.Update(c, "Filename", "Crc", "Buildnumber", "Checksum", "Filepath", "Updated", "Filetype", "Isdownloaded")
	if err != nil {
		log.Println(err.Error())
		o.Rollback()
		return err
	}
	o.Commit()
	return err
}
