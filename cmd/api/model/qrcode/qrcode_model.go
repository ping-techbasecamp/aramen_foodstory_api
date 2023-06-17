package qrcodeModel

import "encoding/json"

func UnmarshalQrCode(data []byte) (QrCode, error) {
	var r QrCode
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *QrCode) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type QrCode struct {
	Qrcode string `json:"qrcode"`
}
