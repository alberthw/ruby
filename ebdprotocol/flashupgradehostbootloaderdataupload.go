package ebdprotocol

type FlashUpgradeHostBootLoaderDataUpload struct {
	Frame
	RecordData string
	LineFeed   string
}

func (f FlashUpgradeHostBootLoaderDataUpload) ToByte() []byte {
	var result []byte
	result = append(result, []byte(f.RecordData)...)
	result = append(result, []byte(f.LineFeed)...)
	return result
}

func (f FlashUpgradeHostBootLoaderDataUpload) Message() []byte {
	f.Frame.Init()
	f.MessageID = []byte{0x43, 0x30}
	f.MessageData = f.ToByte()
	return f.ByteArray()
}
