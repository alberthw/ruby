package mongoose

import (
	"fmt"
	"time"

	"github.com/alberthw/ruby/models"
	"github.com/alberthw/ruby/serial"
)

func sendSetSystemConfig() {
	sendMongooseCommand("config.set.sys")
}

func sendSetHardwareConfig() {
	sendMongooseCommand("config.set.hw")
}

func SendGetVersions() {
	sendMongooseCommand("ver.get")
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


*/
