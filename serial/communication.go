package serial

import (
	"log"
	"time"

	"github.com/tarm/serial"
)

var (
	s   *serial.Port
	err error
)

func Open(name string, baud int) error {
	log.Println("before open:", s)

	//	Close()

	if s != nil {
		return nil
	}

	var c serial.Config
	c.Name = name
	c.Baud = baud
	c.ReadTimeout = time.Millisecond * 10

	log.Println("before open 1 :", c, s)
	s, err = serial.OpenPort(&c)

	log.Println("after open:", c, s, err)
	return err
}

func Close() error {
	log.Println("before close", s)
	if s == nil {
		return nil
	}
	err = s.Close()
	log.Println("after close", s, err)

	s = nil

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
	time.Sleep(time.Millisecond * 10)
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

	if s == nil {
		return []byte(""), nil
	}

	n, err := s.Read(result)
	//	log.Println("serial read:", n, err)
	if err != nil {
		log.Println(err.Error())
		return []byte(""), err
	}
	return result[:n], nil
}
