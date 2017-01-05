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

	"github.com/PuerkitoBio/goquery"
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
	Remotepath   string
	Filetype     FileType
	Filesize     string
	Isdownloaded bool
	Created      time.Time `orm:"auto_now_add;type(datetime)"`
	Updated      time.Time `orm:"auto_now;type(datetime)"`
}

func getRemoteFileRepositoryURL() (string, error) {
	var remoteServer Remoteserver
	remoteServer = remoteServer.Get()
	if len(remoteServer.Remoteserver) == 0 {
		return "", errors.New("invalid remote server URL")
	}
	if !remoteServer.Isconnected {
		err := fmt.Sprintf("conncect to repository failed , %s", remoteServer.Remoteserver)
		return "", errors.New(err)
	}
	repoURL := "http://" + remoteServer.Remoteserver + "/userContent/Release/"

	return repoURL, nil

}

func getFileInfoFromRemoteRepo() ([]Filerepo, error) {

	repoURL, err := getRemoteFileRepositoryURL()
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocument(repoURL)
	if err != nil {
		e := fmt.Sprintf("parse the file repository failed, %s", err.Error())
		return nil, errors.New(e)
	}
	var result []Filerepo
	doc.Find("tr").Each(func(i int, s *goquery.Selection) {
		node := s.Find("a")
		if node.Length() != 2 {
			return
		}
		filename := node.First().Text()
		var f Filerepo
		f.Filename = filename
		f.getBuildNumber()
		if f.Buildnumber == 0 {
			return
		}
		filesize := s.Find("td.fileSize").Text()
		if len(filesize) > 4 {
			f.Filesize = filesize[:len(filesize)-3]
		}
		f.Remotepath = "/userContent/Release/" + f.Filename
		f.checkFileType()
		result = append(result, f)
	})
	return result, nil
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

func (f *Filerepo) getCheckSum() {
	if len(f.Filename) == 0 {
		return
	}
	f.Checksum = caculateChecksum(f.Filepath)
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
		//		log.Println("checkCRC() : ", err.Error())
		return ""
	}
	defer f.Close()
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

func (f *Filerepo) getBuildNumber() {
	if len(f.Filename) < 12 {
		err := errors.New("invalid file name")
		log.Println("(f *Filerepo)getBuildNumber()", err.Error())
		return
	}
	bnStr := f.Filename[len(f.Filename)-12 : len(f.Filename)-4]
	f.Buildnumber, _ = strconv.ParseUint(bnStr, 10, 64)
}

/*
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
*/

func (c *Filerepo) GetFileInfo() error {
	c.Filepath = strings.Replace(c.Filepath, "\\", "/", -1)
	filename := filepath.Base(c.Filepath)
	c.Filename = filename
	//	c.getBuildNumber()
	//	c.Created = fi.ModTime()
	c.checkFileType()
	c.getCRC()
	//	c.getCheckSum()

	c.UpdateCRC()
	//	c.UpdateChecksum()
	//	c.UpdateFileSize()
	return nil
}

func getReleaseFilesInfo(folder string) []Filerepo {
	var result []Filerepo
	filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if info == nil {
			return nil
		}
		//		log.Println("file infor", info, " , path:", path)
		if !info.IsDir() {
			var fr Filerepo
			fr.Filepath = path
			fr.GetFileInfo()

			//		log.Println(fr)
			result = append(result, fr)
		}
		return nil
	})
	return result
}

/*func clearFileRepo() error {
	o := orm.NewOrm()
	_, err := o.Raw("delete from filerepo").Exec()
	if err != nil {
		log.Println("clearFileRepo():", err.Error())
	}
	return err
}
*/
func (c *Filerepo) checkDownloadStatus() {
	fi, err := os.Stat(c.Filepath)
	if os.IsNotExist(err) {
		c.Isdownloaded = false
		return
	}
	/*
		checkSum := caculateChecksum(localFullPath)
		if checkSum == c.Checksum {
			c.Isdownloaded = true
			return
		}
	*/
	localFileSize := fmt.Sprintf("%.2f", float64(fi.Size())/1024)
	if localFileSize == c.Filesize {
		c.Isdownloaded = true
		return
	}
	c.Isdownloaded = false
}

func (c *Filerepo) checkLocalFileStatus() {
	var setting Rubyconfig
	setting = setting.Get()
	c.Filepath = setting.Localrepo + "/" + c.Filename
	c.GetFileInfo()
	c.checkDownloadStatus()
}

func SyncReleaseFilesInfo() {
	//	clearFileRepo()
	//	files := getReleaseFilesInfo("./static/release")
	files, _ := getFileInfoFromRemoteRepo()
	for _, file := range files {
		file.checkLocalFileStatus()
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

	var lists []Filerepo
	// o := orm.NewOrm()
	//	o.QueryTable("Filerepo").GroupBy("Filetype", "Id").OrderBy("Filetype", "Buildnumber").All(&lists, "Id", "Filename", "Crc", "Buildnumber", "Filepath", "Filetype", "Isdownloaded", "Remotepath", "Filesize")
	host := getReleaseFilesByType(FILETYPE_APP, 5)
	lists = append(lists, host...)

	boot := getReleaseFilesByType(FILETYPE_BOOT, 5)
	lists = append(lists, boot...)

	dsp := getReleaseFilesByType(FILETYPE_DSP, 5)
	lists = append(lists, dsp...)

	return lists
}

func getReleaseFilesByType(t FileType, limit int64) []Filerepo {
	o := orm.NewOrm()
	var result []Filerepo
	o.QueryTable("Filerepo").Filter("Filetype", t).OrderBy("Buildnumber").Limit(limit).All(&result, "Id", "Filename", "Crc", "Buildnumber", "Filepath", "Filetype", "Isdownloaded", "Remotepath", "Filesize")
	return result
}

func (c *Filerepo) CreateOrUpdate() error {
	o := orm.NewOrm()
	var tmp Filerepo
	err := o.QueryTable("Filerepo").Filter("Filename", c.Filename).One(&tmp)
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

func (c *Filerepo) UpdateCRC() error {
	c.Updated = time.Now()
	o := orm.NewOrm()
	o.Begin()
	_, err := o.Update(c, "Crc")
	if err != nil {
		log.Println(err.Error())
		o.Rollback()
		return err
	}
	o.Commit()
	return err
}

func (c *Filerepo) UpdateChecksum() error {
	c.Updated = time.Now()
	o := orm.NewOrm()
	o.Begin()
	_, err := o.Update(c, "Checksum")
	if err != nil {
		log.Println(err.Error())
		o.Rollback()
		return err
	}
	o.Commit()
	return err
}

func (c *Filerepo) UpdateFileSize() error {
	c.Updated = time.Now()
	o := orm.NewOrm()
	o.Begin()
	_, err := o.Update(c, "Filesize")
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
	_, err := o.Update(c, "Filename", "Crc", "Buildnumber", "Filepath", "Updated", "Filetype", "Isdownloaded")
	if err != nil {
		log.Println(err.Error())
		o.Rollback()
		return err
	}
	o.Commit()
	return err
}

func (c *Filerepo) DeleteByFilename(filename string) error {
	o := orm.NewOrm()
	err := o.QueryTable("Filerepo").Filter("Filename", filename).One(c)
	if err == orm.ErrNoRows {
		return err
	}
	_, err = o.Delete(c)
	return err
}
