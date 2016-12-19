package ebdprotocol

type FlashUpgradeHostAppDataUpload struct {
	Frame
	RecordData string
	LineFeed   string
}

func (f FlashUpgradeHostAppDataUpload) ToByte() []byte {
	var result []byte
	result = append(result, []byte(f.RecordData)...)
	result = append(result, []byte(f.LineFeed)...)
	return result
}

func (f FlashUpgradeHostAppDataUpload) Message() []byte {
	f.Frame.Init()
	f.MessageID = []byte{0x41, 0x32}
	f.MessageData = f.ToByte()
	return f.ByteArray()
}
