package mongoose

import "github.com/alberthw/ruby/serial"

func sendMongooseCommand(input string) error {

	command := []byte(input + "\r\n")
	return serial.Writer(command)
}

func SendEnterSeviceMode() {
	sendMongooseCommand("service.mode")
}

func SendExitServiceMode() {
	sendMongooseCommand("submode.exit")
}
