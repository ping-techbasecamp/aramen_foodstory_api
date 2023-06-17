package qrcodeRepo

import (
	"aramen-api/cmd/api/ent"
	qrcodeModel "aramen-api/cmd/api/model/qrcode"
	"aramen-api/srv/aes"
	"aramen-api/srv/foodstory"
	"fmt"
	"os"
)

type QrcodeRepoInterface interface {
	ScanQr(userId int, qrcode string) (*int, *int, error)
}

type QrcodeRepo struct {
	db *ent.Client
}

func NewQrcodeRepo(db *ent.Client) QrcodeRepo {
	return QrcodeRepo{db}
}

func (r QrcodeRepo) ScanQr(userId int, qrcode string) (*int, *int, error) {
	decryptedCode, err := aes.DecryptECB(os.Getenv("FOODSTORY_AES_KEY"), qrcode)
	if err != nil {
		return nil, nil, err
	}
	foodstoryQrcode, err := qrcodeModel.UnmarshalQrcode([]byte(decryptedCode))
	if err != nil {
		return nil, nil, err
	}
	fmt.Println(foodstoryQrcode)

	foodstoryOrder, err := foodstory.ScanQrcodeFoodstory(foodstoryQrcode)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}

	fmt.Println(foodstoryOrder)

	// tx, err := r.db.Tx(context.Background())
	// if err != nil {
	// 	return nil, nil, err
	// }
	// qrcodeSavedResult, err := tx.Qrcode.Create().SetBranchName(foodstoryOrder.Data.BranchName).SetOrderChannel(foodstoryOrder.Data.BranchName).SetOrderID(foodstoryOrder.Data.OrderID).SetOrderStatus(foodstoryOrder.Data.OrderStatus).SetQrcodeID(foodstoryOrder.Data.QrcodeID).SetTotalPaid(foodstoryOrder.Data.TotalPaid).SetUserID(userId).Save(context.Background())
	// if err != nil {
	// 	tx.Rollback()
	// }

	point := int(*foodstoryOrder.Data.TotalPaid)
	star := int(point / 200)
	// _, err = tx.User.UpdateOneID(userId).SetStar(star).SetPoint(point).Save(context.Background())
	// if err != nil {
	// 	tx.Rollback()
	// }
	// err = tx.Commit()
	// if err != nil {
	// 	return nil, nil, err
	// }

	return &point, &star, nil
}
