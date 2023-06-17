package qrcodeModel

import "encoding/json"

func UnmarshalQrcode(data []byte) (FoodstoryQrcode, error) {
	var r FoodstoryQrcode
	err := json.Unmarshal(data, &r)
	return r, err
}

type FoodstoryQrcode struct {
	OrderId    string `json:"orderId" binding:"required"`
	BranchUUId string `json:"branchUUId" binding:"required"`
	QrcodeKey  string `json:"qrcodeKey" binding:"required"`
}
