package ebdprotocol

type FlashUpgradeHostAppInitiateBurn struct {
	Frame
}

func (d FlashUpgradeHostAppInitiateBurn) Message() []byte {
	d.Frame.Init()
	d.MessageID = []byte{0x41, 0x34}
	d.MessageData = nil
	return d.ByteArray()
}
