package ebdprotocol

import (
	"bytes"
	"encoding/binary"
)

type FlashUpgradeHostAppBurnDone struct {
	Frame
	IsCompleted bool
}

func (f FlashUpgradeHostAppBurnDone) ToByte() []byte {
	buf := new(bytes.Buffer)
	if f.IsCompleted {
		binary.Write(buf, binary.BigEndian, uint32(1))
	} else {
		binary.Write(buf, binary.BigEndian, uint32(0))
	}

	return buf.Bytes()
}

func (f FlashUpgradeHostAppBurnDone) Message() []byte {
	f.Frame.Init()
	f.MessageID = []byte{0x42, 0x44}
	f.MessageData = f.ToByte()
	return f.ByteArray()
}
