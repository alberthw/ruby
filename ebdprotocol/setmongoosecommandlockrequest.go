package ebdprotocol

import (
	"bytes"
	"encoding/binary"
)

type SetMongooseCommandLockRequest struct {
	Frame
	IsLocked bool
}

func (f SetMongooseCommandLockRequest) ToByte() []byte {
	buf := new(bytes.Buffer)
	if f.IsLocked {
		binary.Write(buf, binary.BigEndian, uint32(1))
	} else {
		binary.Write(buf, binary.BigEndian, uint32(0))
	}

	return buf.Bytes()
}

func (d SetMongooseCommandLockRequest) Message() []byte {
	d.Frame.Init()
	d.MessageID = []byte{0x44, 0x32}
	d.MessageData = d.ToByte()
	return d.ByteArray()
}
