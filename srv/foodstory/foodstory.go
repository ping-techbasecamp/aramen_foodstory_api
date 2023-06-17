package foodstory

import (
	qrcodeModel "aramen-api/cmd/api/model/qrcode"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func ScanQrcodeFoodstory(qrcode qrcodeModel.FoodstoryQrcode) (*qrcodeModel.FoodStoryOrder, error) {
	order, err := getOrderDetail(qrcode)
	if err != nil {
		return nil, err
	}
	if order.Data.UsedQrcode {
		return nil, errors.New("QRCODE USED")
	}
	// err = updateOrder(qrcode)
	// if err != nil {
	// 	return nil, err
	// }
	return order, nil
}

func getOrderDetail(qrcode qrcodeModel.FoodstoryQrcode) (*qrcodeModel.FoodStoryOrder, error) {
	client := &http.Client{}
	jsonValue, err := json.Marshal(qrcode)
	if err != nil {
		return nil, err
	}
	// Get Food Story Order Detail
	req, err := http.NewRequest("POST", "https://fs-ent-crm-api.fs-sandbox.com/v1/crm-aramen/send-order-data", bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, err
	}
	req.Header.Add("x-api-key", os.Getenv("FOODSTORY_API_KEY"))
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))
	fmt.Println("real body")
	order, err := qrcodeModel.UnmarshalFoodStoryOrder(body)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

// func updateOrder(qrcode qrcodeModel.FoodstoryQrcode) error {
// 	client := &http.Client{}
// 	jsonValue, err := json.Marshal(qrcode)
// 	if err != nil {
// 		return err
// 	}
// 	req, err := http.NewRequest("POST", "https://fs-ent-crm-api.fs-sandbox.com/v1/crm-aramen/update-qrcode-status", bytes.NewBuffer(jsonValue))
// 	if err != nil {
// 		return err
// 	}
// 	req.Header.Add("x-api-key", os.Getenv("FOODSTORY_API_KEY"))
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return err
// 	}
// 	if resp.StatusCode == 200 {
// 		return nil
// 	}
// 	return errors.New("Invalid Data")
// }
