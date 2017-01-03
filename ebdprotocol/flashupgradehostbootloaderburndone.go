package ebdprotocol

import (
	"bytes"
	"encoding/binary"
)

type FlashUpgradeHostBootLoaderBurnDone struct {
	Frame
	IsCompleted bool
}

func (f FlashUpgradeHostBootLoaderBurnDone) ToByte() []byte {
	buf := new(bytes.Buffer)
	if f.IsCompleted {
		binary.Write(buf, binary.BigEndian, uint32(1))
	} else {
		binary.Write(buf, binary.BigEndian, uint32(0))
	}

	return buf.Bytes()
}

func (f FlashUpgradeHostBootLoaderBurnDone) Message() []byte {
	f.Frame.Init()
	f.MessageID = []byte{0x43, 0x33}
	f.MessageData = f.ToByte()
	return f.ByteArray()
}
