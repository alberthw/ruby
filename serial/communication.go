package serial

import (
	"errors"
	"log"
	"time"

	"github.com/tarm/serial"
)

var (
	GSerial *serial.Port
)

func Open(name string, baud int) error {
	//	log.Println("before open:", GSerial)

	//	Close()

	if GSerial != nil {
		return nil
	}

	var c serial.Config
	c.Name = name
	c.Baud = baud
	//	c.ReadTimeout = time.Millisecond * 1000

	//	log.Println("before open 1 :", c, GSerial)
	var err error
	GSerial, err = serial.OpenPort(&c)

	//	log.Println("after open:", c, GSerial, err)
	return err
}

func Close() error {
	//	log.Println("before close", GSerial)
	if GSerial == nil {
		return nil
	}
	err := GSerial.Close()

	// waiting for the port to shutdown
	time.Sleep(time.Millisecond * 1500)
	//	log.Println("after close", GSerial, err)

	GSerial = nil

	return err
}

func Sender(msg []byte) []byte {

	//	s, _ := serial.OpenPort(c)

	//	log.Printf("SEND: %X", msg)

	//	defer s.Close()
	var n int
	n, err := GSerial.Write(msg)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Millisecond * 500)
	result := make([]byte, 2048)
	n, err = GSerial.Read(result)
	if err != nil {
		log.Fatal(err)
	}
	return result[:n]

}

func Writer(msg []byte) error {
	if GSerial == nil {
		err := "Serial port is closed."
		return errors.New(err)
	}
	_, err := GSerial.Write(msg)
	//	log.Println("serial write:", n, err)
	if err != nil {
		log.Println(err.Error())
	}
	return err
}

func Reader() ([]byte, error) {
	result := make([]byte, 2048)

	if GSerial == nil {
		return []byte(""), nil
	}

	n, err := GSerial.Read(result)
	//	log.Println("serial read:", n, err)
	if err != nil {
		log.Println(err.Error())
		return []byte(""), err
	}
	return result[:n], nil
}
