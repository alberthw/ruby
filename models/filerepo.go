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
	Id           int64     `orm:"pk;auto;column(id)"`
	FileName     string    `orm:"column(filename)"`
	CRC          string    `orm:"column(crc)"`
	BuildNumber  uint64    `orm:"column(buildnumber)"`
	CheckSum     string    `orm:"column(checksum)"`
	LocalPath    string    `orm:"column(localpath)"`
	RemotePath   string    `orm:"column(remotepath)"`
	FileType     FileType  `orm:"column(filetype)"`
	FileSize     string    `orm:"column(filesize)"`
	IsDownloaded bool      `orm:"column(isdownloaded)"`
	Created      time.Time `orm:"auto_now_add;type(datetime)"`
	Updated      time.Time `orm:"auto_now;type(datetime)"`
}

func getRemoteFileRepositoryURL() (string, string, error) {
	rs := GetRepoSetting()
	if len(rs.Remoteserver) == 0 {
		return "", "", errors.New("invalid remote server URL")
	}
	if !rs.Isconnected {
		err := fmt.Sprintf("conncect to repository failed , %s", rs.Remoteserver)
		return "", "", errors.New(err)
	}
	repoURL := "http://" + rs.Remoteserver + rs.Remotefolder

	return repoURL, rs.Remotefolder, nil

}

func getFileInfoFromRemoteRepo() ([]Filerepo, error) {

	repoURL, remoteFolder, err := getRemoteFileRepositoryURL()
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
		f.FileName = filename
		f.getBuildNumber()
		if f.BuildNumber == 0 {
			return
		}
		filesize := s.Find("td.fileSize").Text()
		if len(filesize) > 4 {
			f.FileSize = filesize[:len(filesize)-3]
		}
		f.RemotePath = remoteFolder + "/" + f.FileName
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
	if len(f.FileName) == 0 {
		return
	}
	f.CheckSum = caculateChecksum(f.LocalPath)
}

func (f *Filerepo) checkFileType() {
	filename := strings.ToLower(f.FileName)
	index := strings.Index(filename, "dsp")
	if index > -1 {
		f.FileType = FILETYPE_DSP
		return
	}
	index = strings.Index(filename, "boot")
	if index > -1 {
		f.FileType = FILETYPE_BOOT
		return
	}
	index = strings.Index(filename, "host")
	if index > -1 {
		f.FileType = FILETYPE_APP
		return
	}
	f.FileType = FILETYPE_UNKNOWN
}

func findCRCLine(filepath string, address string) string {
	f, err := os.Open(filepath)
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
}

func (f *Filerepo) getCRC() {
	switch f.FileType {
	case FILETYPE_APP:
		f.CRC = findCRCLine(f.LocalPath, "000BFFF0")
	case FILETYPE_BOOT:
		f.CRC = findCRCLine(f.LocalPath, "0003FFF0")
	case FILETYPE_DSP:
		f.CRC = findCRCLine(f.LocalPath, "S003")
	case FILETYPE_UNKNOWN:
		// do nothing
	default:
		// do nothing
	}
	f.CRC = findCRC(f.CRC, f.FileType)
	//	log.Println("CRC:", f.Crc)
}

func (f *Filerepo) getBuildNumber() {
	if len(f.FileName) < 12 {
		err := errors.New("invalid file name")
		log.Println("(f *Filerepo)getBuildNumber()", err.Error())
		return
	}
	bnStr := f.FileName[len(f.FileName)-12 : len(f.FileName)-4]
	f.BuildNumber, _ = strconv.ParseUint(bnStr, 10, 64)
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

func (c *Filerepo) getFileInfo() error {
	c.LocalPath = strings.Replace(c.LocalPath, "\\", "/", -1)
	filename := filepath.Base(c.LocalPath)
	c.FileName = filename
	c.getBuildNumber()
	//	c.Created = fi.ModTime()
	c.checkFileType()
	c.getCRC()
	c.getCheckSum()

	c.checkDownloadStatus()

	//	c.UpdateCRC()
	//	c.UpdateChecksum()
	//	c.UpdateFileSize()
	return nil
}

func getLocalReleaseFilesInfo(folder string) []Filerepo {
	var result []Filerepo
	filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if info == nil {
			return nil
		}
		//		log.Println("file infor", info, " , path:", path)
		if !info.IsDir() {
			var fr Filerepo
			fr.LocalPath = path
			fr.getFileInfo()

			//		log.Println(fr)
			result = append(result, fr)
		}
		return nil
	})
	return result
}

/*
func clearFileRepo() error {
	o := orm.NewOrm()
	_, err := o.Raw("delete from filerepo").Exec()
	if err != nil {
		log.Println("clearFileRepo():", err.Error())
	}
	return err
}
*/

func (c *Filerepo) checkDownloadStatus() {
	fi, err := os.Stat(c.LocalPath)
	//	fmt.Printf("file path : %s, error : %v\n", c.Filepath, err)
	if err == nil {
		c.FileSize = fmt.Sprintf("%.2f", float64(fi.Size())/1024)
		c.IsDownloaded = true
		return
	}
	/*
			checkSum := caculateChecksum(localFullPath)
			if checkSum == c.Checksum {
				c.Isdownloaded = true
				return
			}

		localFileSize := fmt.Sprintf("%.2f", float64(fi.Size())/1024)
		log.Printf("remote size :%f, local size : %f", c.Filesize, localFileSize)
		if localFileSize == c.Filesize {
			c.Isdownloaded = true
			return
		}
	*/
	c.IsDownloaded = false
}

/*
func (c *Filerepo) checkLocalFileStatus() {
	setting := GetRepoSetting()
	pwd, _ := os.Getwd()
	c.Filepath = pwd + setting.Localfolder + "/" + c.Filename
	c.getFileInfo()
	c.checkDownloadStatus()
}

*/

func SyncReleaseFilesInfo() {
	//	clearFileRepo()
	//	files := getReleaseFilesInfo("./static/release")
	// check release files in remote file repository
	var releaseFiles []Filerepo
	remoteFiles, _ := getFileInfoFromRemoteRepo()
	releaseFiles = append(releaseFiles, remoteFiles...)

	//check release files in local file repository
	pwd, _ := os.Getwd()
	repoSetting := GetRepoSetting()
	localRepoFolder := pwd + repoSetting.Localfolder

	localFiles := getLocalReleaseFilesInfo(localRepoFolder)
	releaseFiles = append(releaseFiles, localFiles...)

	removeObsoleteRecords(releaseFiles)

	for _, file := range releaseFiles {
		file.CreateOrUpdate()
	}
}

func (c Filerepo) CheckFileIsExistsInArray(files []Filerepo) bool {
	for _, file := range files {
		if strings.Compare(strings.ToLower(c.FileName), strings.ToLower(file.FileName)) == 0 {
			return true
		}
	}
	return false
}

func removeObsoleteRecords(files []Filerepo) {
	recordsInDB := GetALLReleaseFiles()
	for _, file := range recordsInDB {
		if !file.CheckFileIsExistsInArray(files) {
			file.delete()
		}
	}
}

func (c Filerepo) Insert() error {
	o := orm.NewOrm()
	_, err := o.Insert(&c)
	return err
}
func GetALLReleaseFiles() []Filerepo {
	var lists []Filerepo
	o := orm.NewOrm()
	o.QueryTable("Filerepo").GroupBy("Filetype", "Id").OrderBy("FileType", "-id").All(&lists, "ID", "FileName", "CRC", "BuildNumber", "LocalPath", "FileType", "IsDownloaded", "RemotePath", "FileSize")
	/*
		host := getReleaseFilesByType(FILETYPE_APP, 5)
		lists = append(lists, host...)

		boot := getReleaseFilesByType(FILETYPE_BOOT, 5)
		lists = append(lists, boot...)

		dsp := getReleaseFilesByType(FILETYPE_DSP, 5)
		lists = append(lists, dsp...)
	*/
	return lists
}

func GetReleaseFilesWithFilter(date string) []Filerepo {
	var lists []Filerepo

	//	fmt.Println("filter : ", date)
	if len(date) != 10 { //   dd/mm/yyyy
		return lists
	}
	f := date[:2] + date[3:5]
	//	fmt.Println("filter : ", f)
	o := orm.NewOrm()
	o.QueryTable("Filerepo").Filter("Filename__icontains", f).GroupBy("FileType", "ID").OrderBy("FileType", "-id").All(&lists, "ID", "FileName", "CRC", "BuildNumber", "LocalPath", "FileType", "IsDownloaded", "RemotePath", "FileSize")

	return lists
}

func getReleaseFilesByType(t FileType, limit int64) []Filerepo {
	o := orm.NewOrm()
	var result []Filerepo
	o.QueryTable("Filerepo").Filter("FileType", t).OrderBy("-ID").Limit(limit).All(&result, "ID", "FileName", "CRC", "BuildNumber", "LocalPath", "FileType", "IsDownloaded", "RemotePath", "FileSize")
	return result
}

func (c *Filerepo) CreateOrUpdate() error {
	o := orm.NewOrm()
	var tmp Filerepo
	err := o.QueryTable("Filerepo").Filter("FileName", c.FileName).One(&tmp)
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
	_, err := o.Update(c, "IsDownloaded")
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
	_, err := o.Update(c, "CRC")
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
	_, err := o.Update(c, "CheckSum")
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
	_, err := o.Update(c, "FileSize")
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
	_, err := o.Update(c, "FileName", "CRC", "CheckSum", "BuildNumber", "FileSize", "LocalPath", "FileType", "IsDownloaded", "Updated")
	if err != nil {
		log.Println(err.Error())
		o.Rollback()
		return err
	}
	o.Commit()
	return err
}

func (c Filerepo) delete() error {
	o := orm.NewOrm()
	o.Begin()
	_, err := o.Delete(&c)
	if err != nil {
		log.Println(err.Error())
		o.Rollback()
		return err
	}
	o.Commit()
	return err

}
