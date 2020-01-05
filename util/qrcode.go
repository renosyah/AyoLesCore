package util

import (
	qrcode "github.com/skip2/go-qrcode"
)

type Qrcode struct {
	Value string
	Size  int
}

func (q *Qrcode) MakeByte() ([]byte, error) {
	var png []byte
	png, err := qrcode.Encode(q.Value, qrcode.Medium, q.Size)
	if err != nil {
		return png, err
	}
	return png, nil
}
