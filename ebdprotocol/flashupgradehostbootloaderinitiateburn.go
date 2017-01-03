package ebdprotocol

type FlashUpgradeHostBootLoaderInitiateBurn struct {
	Frame
}

func (d FlashUpgradeHostBootLoaderInitiateBurn) Message() []byte {
	d.Frame.Init()
	d.MessageID = []byte{0x43, 0x32}
	d.MessageData = nil
	return d.ByteArray()
}
