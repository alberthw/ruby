package serial

import (
	"log"
	"time"

	"github.com/tarm/serial"
)

var (
	c   *serial.Config
	s   *serial.Port
	err error
)

func Open(name string, baud int) error {
	if s != nil {
		return nil
	}
	if c == nil {
		c = new(serial.Config)
		c.Baud = baud
		c.Name = name
		c.ReadTimeout = time.Millisecond * 10
	}
	s = new(serial.Port)
	s, err = serial.OpenPort(c)
	return err
}

func Close() error {
	if s == nil {
		return nil
	}
	err = s.Close()
	if err == nil {
		s = nil
	}
	return err
}

func Sender(msg []byte) []byte {

	//	s, _ := serial.OpenPort(c)

	log.Printf("SEND: %X", msg)

	//	defer s.Close()
	var n int
	n, err := s.Write(msg)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Millisecond * 100)
	result := make([]byte, 2048)
	n, err = s.Read(result)
	if err != nil {
		log.Fatal(err)
	}
	return result[:n]

}

func Writer(msg []byte) error {
	_, err := s.Write(msg)
	//	log.Println("serial write:", n, err)
	if err != nil {
		log.Println(err.Error())
	}
	return err
}

func Reader() ([]byte, error) {
	result := make([]byte, 2048)

	n, err := s.Read(result)
	//	log.Println("serial read:", n, err)
	if err != nil {
		log.Println(err.Error())
		return []byte(""), err
	}
	return result[:n], nil
}
