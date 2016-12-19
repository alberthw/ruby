package ebdprotocol

type FlashUpgradeHostAppRequest struct {
	Frame
}

func (d FlashUpgradeHostAppRequest) Message() []byte {
	d.Frame.Init()
	d.MessageID = []byte{0x41, 0x30}
	d.MessageData = nil
	return d.ByteArray()
}
