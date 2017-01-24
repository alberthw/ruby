package mongoose

import (
	"bufio"
	"log"
	"os"
	"time"

	"fmt"

	"github.com/alberthw/ruby/models"
	"github.com/alberthw/ruby/serial"
)

func sendMongooseCommand(input string) error {

	command := []byte(input + "\r\n")
	return serial.Writer(command)
}

func sendImageUpload() {
	sendMongooseCommand("image.upload")
}

func sendImageUpdate() {
	sendMongooseCommand("image.update")
}

func sendEnterSeviceMode() {
	sendMongooseCommand("service.mode")
}

func sendSetSystemConfig() {
	sendMongooseCommand("config.set.sys")
}

func sendSetHardwareConfig() {
	sendMongooseCommand("config.set.hw")
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
		sendMongooseCommand(v)
		pervPos = curPos
		time.Sleep(time.Millisecond * 1)

	}
	return nil
}

func BurnHostImage(filepath string, t models.FileType) error {
	sendEnterSeviceMode()
	time.Sleep(time.Millisecond * 5000)
	sendImageUpload()
	time.Sleep(time.Millisecond * 1000)
	sendSelectHostImage(t)
	time.Sleep(time.Millisecond * 1000)
	err := sendUploadImage(filepath)
	if err != nil {
		return err
	}
	time.Sleep(time.Millisecond * 5000)
	sendImageUpdate()
	time.Sleep(time.Millisecond * 1000)
	sendSelectHostImage(t)
	return nil
}

func WriteDeviceSystemConfig(config models.Devicesystemconfig) error {
	sendSetSystemConfig()
	time.Sleep(time.Millisecond * 1000)
	command := config.ToByte()
	command = append(command, []byte{'\r', '\n'}...)
	fmt.Printf("Device system config : %x\n", command)
	return serial.Writer(command)
}

func WriteDeviceHardwareConfig(config models.Devicehardwareconfig) error {
	sendSetHardwareConfig()
	time.Sleep(time.Millisecond * 1000)
	command := config.ToByte()
	command = append(command, []byte{'\r', '\n'}...)
	fmt.Printf("Device hardware config : %x\n", command)
	return serial.Writer(command)
}

/*
func WriteDeviceSoftwareConfig(config models.Devicesoftwareconfig) error {
	switch config.Type {
	case models.HOSTBOOT:
		sendSetHostBootConfig()
		break
	case models.HOSTAPP:
		sendSetHostAppConfig()
		break
	case models.DSPAPP:
		sendSetDspAppConfig()
		break
	default:
		return errors.New("unknown software type")
	}
	time.Sleep(time.Millisecond * 1000)
	command := config.ToByte()
	command = append(command, []byte{'\r', '\n'}...)
	fmt.Printf("Device software config : %x\n", command)
	return serial.Writer(command)
}


func BurnHostBootImage(filepath string) error {
	return BurnHostImage(filepath, models.FILETYPE_BOOT)
}

func BurnHostAppImage(filepath string) error {
	return BurnHostImage(filepath, models.FILETYPE_APP)
}


*/
