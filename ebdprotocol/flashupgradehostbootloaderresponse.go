package ebdprotocol

import "errors"

import "encoding/hex"

type FlashUpgradeHostBootLoaderResponse struct {
	Frame
	bAllowed bool
}

func (r *FlashUpgradeHostBootLoaderResponse) Parse(input []byte) {
	r.Frame.Parse(input)
	data := r.Frame.MessageData
	r.ParseMessageData(data)
}

func (r *FlashUpgradeHostBootLoaderResponse) ParseMessageData(data []byte) error {
	if len(data) != 4 {
		return errors.New("FlashUpgradeHostBootLoaderResponse.ParseMessageData:invalid length.")
	}
	//	log.Printf("data : |%s|\n", hex.EncodeToString(data))
	if hex.EncodeToString(data) == "00000001" {
		r.bAllowed = true
		return nil
	}
	r.bAllowed = false
	return nil
}
