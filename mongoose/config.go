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

func SendGetLastKnownVersions() {
	sendMongooseCommand("ver.get.lastknown")
}

func SendValidateConfig() {
	sendMongooseCommand("config.validate")
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
