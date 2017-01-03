package ebdprotocol

type GetMongooseCommandLockRequest struct {
	Frame
}

func (d GetMongooseCommandLockRequest) Message() []byte {
	d.Frame.Init()
	d.MessageID = []byte{0x44, 0x34}
	d.MessageData = nil
	return d.ByteArray()
}
