package ebdprotocol

import (
	"encoding/hex"
	"errors"
)

type GetMongooseCommandLockResponse struct {
	Frame
	bLocked bool
}

func (r *GetMongooseCommandLockResponse) Parse(input []byte) {
	r.Frame.Parse(input)
	data := r.Frame.MessageData
	r.ParseMessageData(data)
}

func (r *GetMongooseCommandLockResponse) ParseMessageData(data []byte) error {
	if len(data) != 4 {
		return errors.New("GetMongooseCommandLockResponse.ParseMessageData:invalid length.")
	}
	//	log.Printf("data : |%s|\n", hex.EncodeToString(data))
	if hex.EncodeToString(data) == "00000001" {
		r.bLocked = true
		return nil
	}
	r.bLocked = false
	return nil
}
