package qrcodeService

import (
	qrcodeRepo "aramen-api/srv/repositories/qrcode"
)

type QrcodeServiceInterface interface {
	ScanQr(userId int, qrcode string) (*int, *int, error)
}

type QrcodeService struct {
	qrcodeRepo qrcodeRepo.QrcodeRepo
}

func NewQrcodeService(qrcodeRepo qrcodeRepo.QrcodeRepo) QrcodeService {
	return QrcodeService{qrcodeRepo: qrcodeRepo}
}
func (s QrcodeService) ScanQr(userId int, qrcode string) (*int, *int, error) {
	return s.qrcodeRepo.ScanQr(userId, qrcode)
}
