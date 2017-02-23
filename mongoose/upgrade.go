package mongoose

import (
	"bufio"
	"log"
	"os"
	"time"

	"github.com/alberthw/ruby/models"
)

func sendImageUpload() {
	sendMongooseCommand("image.upload")
}

func sendImageUpdate() {
	sendMongooseCommand("image.update")
}

func sendSelectHostImage(t models.FileType) {
	command := ""
	switch t {
	case models.FILETYPE_BOOT:
		command = "3"
		break
	case models.FILETYPE_APP:
		command = "2"
		break
	default:
	}
	sendMongooseCommand(command)
}

func sendUploadImage(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		msg := scanner.Text()
		lines = append(lines, msg)

	}
	if err := scanner.Err(); err != nil {
		log.Println(err.Error())
		return err
	}
	curPos := 0
	pervPos := 0
	for i, v := range lines {
		curPos = 100 * i / len(lines)
		if i == 0 {
			log.Printf(" |%d| ", curPos)
		}
		if curPos != pervPos {
			log.Printf(" |%d| ", curPos)
		}
		err := sendMongooseCommand(v)
		if err != nil {
			log.Printf("failed to upload the image line : %d, %s, error : %s\n", i, v, err.Error())
		}
		pervPos = curPos
		time.Sleep(time.Millisecond * 1)

	}
	return nil
}

func BurnHostImage(filepath string, t models.FileType) error {
	//	SendEnterSeviceMode()
	//	time.Sleep(time.Millisecond * 5000)
	sendImageUpload()
	time.Sleep(time.Millisecond * 1000)
	sendSelectHostImage(t)
	time.Sleep(time.Millisecond * 1000)
	sendUploadImage(filepath)
	time.Sleep(time.Millisecond * 5000)
	sendImageUpdate()
	time.Sleep(time.Millisecond * 1000)
	sendSelectHostImage(t)
	return nil
}

/*

func BurnHostBootImage(filepath string) error {
	return BurnHostImage(filepath, models.FILETYPE_BOOT)
}

func BurnHostAppImage(filepath string) error {
	return BurnHostImage(filepath, models.FILETYPE_APP)
}


*/
