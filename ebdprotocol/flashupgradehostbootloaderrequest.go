package ebdprotocol

type FlashUpgradeHostBootLoaderRequest struct {
	Frame
}

func (d FlashUpgradeHostBootLoaderRequest) Message() []byte {
	d.Frame.Init()
	d.MessageID = []byte{0x42, 0x45}
	d.MessageData = nil
	return d.ByteArray()
}
