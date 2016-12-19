package ebdprotocol

type DeviceNameRequest struct {
	Frame
}

func (d DeviceNameRequest) Message() []byte {
	d.Frame.Init()
	d.MessageID = []byte{0x31, 0x44}
	d.MessageData = nil
	return d.ByteArray()
}
