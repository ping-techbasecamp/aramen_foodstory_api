package repositories

import (
	authRepo "aramen-api/srv/repositories/auth"
	branchRepo "aramen-api/srv/repositories/branch"
	newsRepo "aramen-api/srv/repositories/news"
	qrcodeRepo "aramen-api/srv/repositories/qrcode"
)

type Repositories struct {
	AuthRepo   authRepo.AuthRepo
	BranchRepo branchRepo.BranchRepo
	NewsRepo   newsRepo.NewsRepo
	QrcodeRepo qrcodeRepo.QrcodeRepo
}
